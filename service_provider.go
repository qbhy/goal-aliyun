package aliyun

import (
	"github.com/goal-web/contracts"
	config2 "github.com/qbhy/goal-aliyun/config"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(app contracts.Application) {
	app.Singleton("alipay", func(config contracts.Config) Factory {
		return NewFactory(config.Get("aliyun").(*config2.Config))
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
