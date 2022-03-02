package config

import "github.com/denverdino/aliyungo/oss"

type Oss struct {
	*Key
	KeyName  string
	Region   oss.Region
	Bucket   string
	Endpoint string
	Debug    bool
	Internal bool
	Secure   bool
}

type OssConfig struct {
	Default string
	Disks   map[string]*Oss
}
