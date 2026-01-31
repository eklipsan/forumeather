package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)


type Forum struct {
	ID int64
	Name string
	Place string
	Topics string
	Latitude float64
	Longitude float64
}


type ForumDB struct {
	DB *sql.DB
}

func (fdb *ForumDB) NewForumDB() error {
	db, err := sql.Open("sqlite3", "forumeather.db")
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	fdb.DB = db
	return nil
}


func (fdb *ForumDB) Close() error {
	err := fdb.DB.Close()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (fdb *ForumDB) GetAllForums() ([]Forum, error) {
	var result []Forum
	rows, err := fdb.DB.Query(`
	SELECT
		id, name, place, topics, latitude, longitude
	FROM forums;`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var forum Forum
		err := rows.Scan(
			&forum.ID,
			&forum.Name,
			&forum.Place,
			&forum.Topics,
			&forum.Latitude,
			&forum.Longitude,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, forum)
	}
	return result, nil
}
