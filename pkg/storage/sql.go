package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)


type ForumDB struct {
	*sql.DB
}

func (fdb *ForumDB) NewForumDB() error {
	db, err := sql.Open("sqlite3", "forums.db")
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	fdb.DB = db
	return nil
}


func (fdb *ForumDB) CreateTable() error {
	return nil
}