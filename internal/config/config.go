package config

type Config struct {
	Addr     string `json:"addr"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "123456",
		DBname:   "db",
	}
}
