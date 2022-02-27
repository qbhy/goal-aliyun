package aliyun

import (
	"github.com/denverdino/aliyungo/oss"
	"github.com/denverdino/aliyungo/push"
	"github.com/denverdino/aliyungo/sms"
	config2 "github.com/qbhy/goal-aliyun/config"
	"sync"
)

type factory struct {
	config *config2.Config

	clients sync.Map
}

func (f factory) Push(name ...string) *push.Client {
	//TODO implement me
	panic("implement me")
}

func (f factory) Oss(name ...string) *oss.Client {
	//TODO implement me
	panic("implement me")
}

func (f factory) Sms(name ...string) *sms.Client {
	//TODO implement me
	panic("implement me")
}

func (f factory) DYSms(name ...string) *sms.DYSmsClient {
	//TODO implement me
	panic("implement me")
}

func NewFactory(config *config2.Config) Factory {
	return &factory{
		config:  config,
		clients: sync.Map{},
	}
}
