package config

type Cdn struct {
	*Key
	KeyName string

	Endpoint string
}

type CdnConfig struct {
	Default string
	Disks   map[string]*Cdn
}
