package config

type Cdn struct {
	*Key
	KeyName string

	EndPoint string
}

type CdnConfig struct {
	Default string
	Disks   map[string]*Cdn
}
