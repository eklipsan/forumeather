package storage

import (
	"database/sql"
	"strings"

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
	db *sql.DB
}

func (fdb *ForumDB) NewForumDB() error {
	db, err := sql.Open("sqlite3", "forumeather.db")
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	fdb.db = db
	return nil
}


func (fdb *ForumDB) SearchForums(query string) ([]Forum, error) {
	var result []Forum
	query = strings.TrimSpace(query)

	searchPattern := "%" + query + "%"

	rows, err := fdb.db.Query(`
		SELECT id, name, place, topics, latitude, longitude
		FROM forums
		WHERE name like ?
		OR place like ?
		OR topics like ?;`, searchPattern, searchPattern, searchPattern,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func (fdb *ForumDB) GetAllForums() ([]Forum, error) {
	var result []Forum
	rows, err := fdb.db.Query(`
	SELECT
		id, name, place, topics, latitude, longitude
	FROM forums;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
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

func (fdb *ForumDB) GetForum(id int) (Forum, error) {
	var forum Forum
	rows := fdb.db.QueryRow(`
	SELECT
		id, name, place, topics, latitude, longitude
	FROM forums
	WHERE id = ?;`, id)

	err := rows.Scan(
		&forum.ID,
		&forum.Name,
		&forum.Place,
		&forum.Topics,
		&forum.Latitude,
		&forum.Longitude,
	)
	if err != nil {
		return forum, err
	}

	return forum, nil
}
