package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Clusters []*Cluster
}

type Cluster struct {
	Name string `json:"Name"`
	File string `json:"File"`
}
