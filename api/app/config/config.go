package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "hbstudent",
			Password: "hbstudent",
			Name:     "users_db",
			Charset:  "utf8",
		},
	}
}
