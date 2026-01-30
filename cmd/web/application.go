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

func getDatabaseConnection() storage.ForumDB {
	forumDB := storage.ForumDB{}
	forumDB.NewForumDB()
	return forumDB
}


type Application struct {
	InfoLog *log.Logger
	ErrorLog *log.Logger
	MeteoblueConfig *config.MeteoblueConfig
	Database storage.ForumDB
}

func NewApplication() *Application {
	errorLog := getErrorLog()
	infoLog := getInfoLog()
	meteoblueConfig, err := config.LoadMeteoblueConfig()
	if err != nil {
		errorLog.Println(err)
	}
	return &Application{
		InfoLog: infoLog,
		ErrorLog: errorLog,
		MeteoblueConfig: meteoblueConfig,
		Database: getDatabaseConnection(),
	}
}