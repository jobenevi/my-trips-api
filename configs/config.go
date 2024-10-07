package config

import (
	"log"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseDriver   string `mapstructure:"DATABASE_DRIVER"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn     string `mapstructure:"JWT_EXPIRES_IN"`
	JWTAuth          *jwtauth.JWTAuth
}

var cfg Config

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.JWTAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return &cfg, nil
}

func BuildStringConnection() string {
	return "host=" + cfg.DatabaseHost +
		" port=" + cfg.DatabasePort +
		" user=" + cfg.DatabaseUser +
		" dbname=" + cfg.DatabaseName +
		" password=" + cfg.DatabasePassword
}
