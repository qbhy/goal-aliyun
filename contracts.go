package aliyun

import (
	"github.com/denverdino/aliyungo/oss"
	"github.com/denverdino/aliyungo/push"
	"github.com/denverdino/aliyungo/sms"
)

type Factory interface {
	Push(name ...string) *push.Client
	Oss(name ...string) *oss.Client
	Sms(name ...string) *sms.Client
	DYSms(name ...string) *sms.DYSmsClient
}
