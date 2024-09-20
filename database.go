package main

import (

	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

)

func initDatabase() *sql.DB {

	db, err := sql.Open("sqlite3", "file:./data/database.db?cache=shared")
    if err != nil {
        fmt.Println(err)
		return nil
    }

	db.SetMaxOpenConns(1)

    // Create a table
    createTableSQL := `CREATE TABLE IF NOT EXISTS data (
        id TEXT NOT NULL,
        type TEXT NOT NULL,
        fileName TEXT NOT NULL,
		filePath TEXT NOT NULL,
		burn TEXT NOT NULL,
		expire TEXT NOT NULL,
		passwordHash TEXT NOT NULL,
		passwordSalt TEXT NOT NULL,
		encryptSalt TEXT NOT NULL
    );`

    _, err = db.Exec(createTableSQL)
    if err != nil {
        fmt.Println(err)
		return nil
    }	

	return db

}
