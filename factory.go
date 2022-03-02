package aliyun

import (
	"github.com/denverdino/aliyungo/cdn"
	"github.com/denverdino/aliyungo/oss"
	"github.com/denverdino/aliyungo/push"
	"github.com/denverdino/aliyungo/sms"
	"github.com/goal-web/supports/utils"
	config2 "github.com/qbhy/goal-aliyun/config"
	"sync"
)

type factory struct {
	config *config2.Config

	clients sync.Map
}

func (f *factory) key(name string, key *config2.Key) *config2.Key {
	if key != nil {
		return key
	}

	if key = f.config.Keys.Keys[name]; key == nil {
		key = f.config.Keys.Keys[f.config.Default]
	}

	return key
}

func (f *factory) Cdn(name ...string) *cdn.CdnClient {
	var config = f.config.Cdn.Disks[utils.DefaultString(name, f.config.Cdn.Default)]
	if config == nil {
		return nil
	}
	var key = f.key(config.KeyName, config.Key)
	return cdn.NewClient(key.AccessKeyId, key.AccessKeySecret)
}

func (f *factory) Push(name ...string) *push.Client {
	var config = f.config.Push.Channels[utils.DefaultString(name, f.config.Push.Default)]
	if config == nil {
		return nil
	}
	var key = f.key(config.KeyName, config.Key)
	return push.NewClient(key.AccessKeyId, key.AccessKeySecret)
}

func (f *factory) Oss(name ...string) *oss.Client {
	var config = f.config.Oss.Disks[utils.DefaultString(name, f.config.Oss.Default)]
	if config == nil {
		return nil
	}
	var (
		key  = f.key(config.KeyName, config.Key)
		disk = oss.NewOSSClient(config.Region, config.Internal, key.AccessKeyId, key.AccessKeySecret, config.Debug)
	)

	disk.SetDebug(config.Debug)

	return disk
}

func (f *factory) Sms(name ...string) *sms.Client {
	var config = f.config.Sms.Sms[utils.DefaultString(name, f.config.Sms.DefaultSms)]
	if config == nil {
		return nil
	}
	var key = f.key(config.KeyName, config.Key)
	return sms.NewClient(key.AccessKeyId, key.AccessKeySecret)
}

func (f *factory) DYSms(name ...string) *sms.DYSmsClient {
	var config = f.config.Sms.DYSms[utils.DefaultString(name, f.config.Sms.DefaultDYSms)]
	if config == nil {
		return nil
	}
	var key = f.key(config.KeyName, config.Key)
	return sms.NewDYSmsClient(key.AccessKeyId, key.AccessKeySecret)
}

func NewFactory(config *config2.Config) Factory {
	return &factory{
		config:  config,
		clients: sync.Map{},
	}
}
