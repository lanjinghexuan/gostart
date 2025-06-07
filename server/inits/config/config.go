package config

type Config struct {
	Mysql
	Redis
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
