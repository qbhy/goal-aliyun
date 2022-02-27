package config

type Sms struct {
	*Key
	KeyName string
}

// DYSms 阿里大于
type DYSms struct {
	*Key
	KeyName string
}

type SmsConfig struct {
	DefaultSms string
	Sms        map[string]*Sms

	DefaultDYSms string
	DYSms        map[string]*DYSms
}
