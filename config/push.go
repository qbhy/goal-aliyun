package config

type Push struct {
	*Key
	KeyName string
}

type PushConfig struct {
	Default  string
	Channels map[string]*Push
}
