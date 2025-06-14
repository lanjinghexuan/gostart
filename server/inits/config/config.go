package config

type Config struct {
	Mysql
	Redis
	MinIO
}

type Mysql struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
	User string `yaml:"User"`
	Pass string `yaml:"Pass"`
	DB   string `yaml:"DB"`
}

type Redis struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type MinIO struct {
	Host      string `yaml:"Host"`
	Port      int    `yaml:"Port"`
	AccessKey string `yaml:"Ak"`
	SecretKey string `yaml:"Sk"`
	UseSSL    bool   `yaml:"UseSSL"`
}
