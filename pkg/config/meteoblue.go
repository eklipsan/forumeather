package config

import (
	"os"

	"github.com/joho/godotenv"
)



type MeteoblueConfig struct {
	URL string
	APIKey string
	TZ string
	Format string
}

func LoadMeteoblueConfig() (*MeteoblueConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	} else {
		return &MeteoblueConfig{
			URL:     os.Getenv("METEOBLUE_URL"),
			APIKey:  os.Getenv("METEOBLUE_API_KEY"),
			TZ:      os.Getenv("METEOBLUE_TZ"),
			Format:  os.Getenv("METEOBLUE_FORMAT"),
		}, nil
	}
}