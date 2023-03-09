package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strconv"
	"strings"
	"time"
)

var _ EventModel = (*customEventModel)(nil)

var eventRowsExpect = strings.Join(stringx.Remove(eventFieldNames, "`id`"), ",")

type (
	// EventModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEventModel.
	EventModel interface {
		eventModel
		HasEvent(name, kind, namespace, reason, cluster string, time *time.Time) error
		FindAll(cluster string, page, limit int) (int, []Event, error)
		InsertEvent(data *Event) (err error)
	}

	customEventModel struct {
		*defaultEventModel
	}
)

// NewEventModel returns a model for the database table.
func NewEventModel(conn sqlx.SqlConn) EventModel {
	return &customEventModel{
		defaultEventModel: newEventModel(conn),
	}
}

func (c *customEventModel) HasEvent(name, kind, namespace, reason, cluster string, eventTime *time.Time) (err error) {
	query := fmt.Sprintf("select * from %s where name = ? and kind = ? and namespace = ? and reason = ? and cluster = ? and eventTime = ?", c.table)
	var events Event
	err = c.conn.QueryRow(events, query, name, kind, namespace, reason, cluster, eventTime)
	switch err {
	case nil:
		return nil
	case sqlc.ErrNotFound:
		return ErrNotFound
	default:
		return err
	}
}

func (c *customEventModel) FindAll(cluster string, page, size int) (int, []Event, error) {

	// count
	query := fmt.Sprintf("select count(1) from %s", c.table)
	var count int
	err := c.conn.QueryRow(&count, query)
	if err != nil {
		return 0, nil, err
	}

	// 分页获取
	query = fmt.Sprintf("select * from %s where cluster = ? order by eventTime desc limit ?,? ", c.table)
	events := make([]Event, 0)
	err = c.conn.QueryRows(&events, query, cluster, strconv.Itoa((page-1)*size), strconv.Itoa(size))
	switch err {
	case nil:
		return count, events, nil
	case sqlc.ErrNotFound:
		return count, events, ErrNotFound
	default:
		return count, events, err
	}
}
func (m *defaultEventModel) InsertEvent(data *Event) (err error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, eventRowsExpect)
	_, err = m.conn.Exec(query, data.Kind, data.Namespace, data.Rtype, data.Reason, data.Message, data.EventTime, data.Cluster, data.CreateTime, data.Name)
	return err
}
