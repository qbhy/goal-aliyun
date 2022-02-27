package aliyun

import (
	"sync"
)

type factory struct {
	config *Config

	clients sync.Map
}

func NewFactory(config *Config) Factory {
	return &factory{
		config:  config,
		clients: sync.Map{},
	}
}
