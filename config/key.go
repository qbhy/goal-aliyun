package config

type Key struct {
	AccessKeyId     string
	AccessKeySecret string
}

type Keys struct {
	Default string
	Keys    map[string]Key
}
