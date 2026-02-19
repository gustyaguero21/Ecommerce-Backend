package config

type Config struct {
	HTTPConfig HTTPConfig
	DB         DBConfig
}

type HTTPConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}
