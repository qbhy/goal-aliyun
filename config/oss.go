package config

import "github.com/denverdino/aliyungo/oss"

type Oss struct {
	*Key
	KeyName  string
	Region   oss.Region
	Bucket   string
	Endpoint string
	Debug    bool
}

type OssConfig struct {
	Default string
	Disks   map[string]*Oss
}
