package main

import (
	"log"
	"os"

	"github.com/eklipsan/forumeather/pkg/config"
	"github.com/eklipsan/forumeather/pkg/storage"
)


func getInfoLog() *log.Logger {
	InfoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	return 	InfoLog
}

func getErrorLog() *log.Logger {
	ErrorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Llongfile)
	return ErrorLog
}

func getDatabaseConnection(pathDB string) storage.ForumDB {
	forumDB := storage.ForumDB{}
	forumDB.NewForumDB(pathDB)
	return forumDB
}


type Application struct {
	InfoLog *log.Logger
	ErrorLog *log.Logger
	Config *config.AppConfig
	Database storage.ForumDB
}

func NewApplication() *Application {
	errorLog := getErrorLog()
	infoLog := getInfoLog()
	AppConfig, err := config.LoadAppConfig()
	if err != nil {
		errorLog.Println(err)
	}
	return &Application{
		InfoLog: infoLog,
		ErrorLog: errorLog,
		Config: AppConfig,
		Database: getDatabaseConnection(AppConfig.DatabaseConfig.PathDB),
	}
}