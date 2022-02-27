package aliyun

import (
	"github.com/goal-web/contracts"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(app contracts.Application) {
	app.Singleton("alipay", func(config contracts.Config) Factory {
		return NewFactory(config.Get("aliyun").(*Config))
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
