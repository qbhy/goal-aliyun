package aliyun

import (
	"github.com/denverdino/aliyungo/cdn"
	"github.com/denverdino/aliyungo/oss"
	"github.com/denverdino/aliyungo/push"
	"github.com/denverdino/aliyungo/sms"
)

type Factory interface {
	Key(name ...string) (id string, secret string, exists bool)
	Push(name ...string) *push.Client
	Oss(name ...string) *oss.Client
	Cdn(name ...string) *cdn.CdnClient
	Sms(name ...string) *sms.Client
	DYSms(name ...string) *sms.DYSmsClient
}
