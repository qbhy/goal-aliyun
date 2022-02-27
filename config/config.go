package config

type Config struct {
	Keys

	Oss  OssConfig
	Sms  SmsConfig
	Push PushConfig
	Cdn  CdnConfig
}
