package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"time"
)

var _ EventModel = (*customEventModel)(nil)

type (
	// EventModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEventModel.
	EventModel interface {
		eventModel
		HasEvent(ctx context.Context, name, kind, namespace, reason, cluster string, time *time.Time) (bool, error)
		FindAll(ctx context.Context, cluster string, page, limit int) ([]Event, error)
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

func (c *customEventModel) HasEvent(ctx context.Context, name, kind, namespace, reason, cluster string, eventTime *time.Time) (bool, error) {
	query := fmt.Sprintf("select * from %s where name = ? and kind = ? and namespace = ? and reason = ? and cluster = ? and eventTime = ?", c.table)
	events := make([]Event, 0)
	err := c.conn.QueryRowsCtx(ctx, events, query, name, kind, namespace, reason, cluster, eventTime)
	switch err {
	case nil:
		return true, nil
	case sqlc.ErrNotFound:
		return false, ErrNotFound
	default:
		return true, err
	}
}

func (c *customEventModel) FindAll(ctx context.Context, cluster string, page, size int) ([]Event, error) {
	query := fmt.Sprintf("select * from %s where cluster = %s limit ?,?", c.table)
	events := make([]Event, 0)

	err := c.conn.QueryRowsCtx(ctx, events, query, strconv.Itoa((page-1)*size), strconv.Itoa(size))
	switch err {
	case nil:
		return events, nil
	case sqlc.ErrNotFound:
		return events, ErrNotFound
	default:
		return events, err
	}
}
