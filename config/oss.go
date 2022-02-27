package config

type Oss struct {
	*Key
	KeyName  string
	Region   string
	Bucket   string
	EndPoint string
	Debug    bool
}

type OssConfig struct {
	Default string
	Disks   map[string]*Oss
}
