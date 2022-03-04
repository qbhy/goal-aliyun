package aliyun

import (
	"github.com/denverdino/aliyungo/cdn"
	"github.com/denverdino/aliyungo/push"
	"github.com/denverdino/aliyungo/sms"
	"github.com/goal-web/contracts"
	config2 "github.com/qbhy/goal-aliyun/config"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(app contracts.Application) {
	app.Singleton("aliyun", func(config contracts.Config) Factory {
		return NewFactory(config.Get("aliyun").(*config2.Config))
	})
	app.Singleton("aliyun.key", func(config contracts.Config) *config2.Key {
		var aliyunConfig = config.Get("aliyun").(*config2.Config)
		return aliyunConfig.Keys.Keys[aliyunConfig.Default]
	})
	app.Singleton("aliyun.sms", func(factory Factory) *sms.Client {
		return factory.Sms()
	})
	app.Singleton("aliyun.cdn", func(factory Factory) *cdn.CdnClient {
		return factory.Cdn()
	})
	app.Singleton("aliyun.sms", func(factory Factory) *sms.Client {
		return factory.Sms()
	})
	app.Singleton("aliyun.dysms", func(factory Factory) *sms.DYSmsClient {
		return factory.DYSms()
	})
	app.Singleton("aliyun.push", func(factory Factory) *push.Client {
		return factory.Push()
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
