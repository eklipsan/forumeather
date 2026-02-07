package config

import (
	"os"

	"github.com/joho/godotenv"
)


type AppConfig struct {
	MeteoblueConfig meteoblueConfig
	DatabaseConfig databaseConfig
	NetworkConfig networkConfig
}

type databaseConfig struct {
	PathDB string
}

type networkConfig struct {
	Port string
}

type meteoblueConfig struct {
	URL string
	APIKey string
	TZ string
	Format string
}

func LoadAppConfig() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	} else {
		return &AppConfig{
			MeteoblueConfig: meteoblueConfig{
			URL:     os.Getenv("METEOBLUE_URL"),
			APIKey:  os.Getenv("METEOBLUE_API_KEY"),
			TZ:      os.Getenv("METEOBLUE_TZ"),
			Format:  os.Getenv("METEOBLUE_FORMAT"),
			},
			DatabaseConfig: databaseConfig{
				PathDB: os.Getenv("DATABASE_PATH"),
			},
			NetworkConfig: networkConfig{
				Port: os.Getenv("PORT"),
			},
		}, nil
	}
}