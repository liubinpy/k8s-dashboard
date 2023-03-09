package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Clusters       []*Cluster
	PodLogTailLine int `json:"PodLogTailLine"`
	// Mysql配置
	Mysql struct {
		Datasource string
	}
}

type Cluster struct {
	Name string `json:"Name"`
	File string `json:"File"`
}
