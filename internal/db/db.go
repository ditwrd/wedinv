package db

import "github.com/pocketbase/dbx"

func NewDB() (*dbx.DB, error) {
	db, err := dbx.Open("sqlite3", "pb_data/data.db")
	print(db)
	print(err)
	return db, err
}
