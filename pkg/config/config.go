package config

import (
	"flag"
)

var WebConfig = loadAppConfig()
var pathDB = "data/forumeather.db"

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
	APIKey string
}

func loadAppConfig() (*AppConfig) {
	api_key := flag.String("api-key", "nXbOjqDSC12ALe3s", "API key of meteoblue api forecast service. The default one is disabled and gives empty data values")
	web_port := flag.String("port", ":8080", "Port to run the web server on.")

	flag.Parse()
	appConfig := &AppConfig{
		MeteoblueConfig: meteoblueConfig{
			APIKey:  *api_key,
		},
		DatabaseConfig: databaseConfig{
			PathDB: pathDB,
		},
		NetworkConfig: networkConfig{
			Port: *web_port,
		},
	}
	return appConfig
}
