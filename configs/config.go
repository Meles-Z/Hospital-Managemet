package configs

import "github.com/spf13/viper"

type Config struct {
	Auth     AuthenticationConfigration
	Database DatabaseConfigration
	Server   ServerConfigration
}
type DatabaseConfigration struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

type ServerConfigration struct {
	Host string
	Port string
}

type AuthenticationConfigration struct {
	WebSecret string
	Secret    string
}

func LoadConfig() (*Config, error) {
	var cfg Config
	var err error

	viper.AddConfigPath("./")
	viper.SetConfigName(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.BindEnv("server.host", "SERVER_HOST"); err != nil {
		return nil, err
	}

	if err = viper.BindEnv("server.port", "SERVER_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.host", "DB_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.port", "DB_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.name", "DB_NAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.username", "DB_USER"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.password", "DB_PASSWORD"); err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
