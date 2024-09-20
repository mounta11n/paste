package main

import (
	
	"database/sql"
	"fmt"
	"os"
	"time"
	"strconv"

)

func CheckExpiration(db *sql.DB){

	for{

		tx, err := db.Begin()
		if err != nil {
			fmt.Println(err)
		}

		//Sqlite doesn't support select for update, instead when we begin transaction it locks the whole db file
		rows, err := tx.Query("SELECT id, filePath FROM data WHERE expire <= ?", strconv.FormatInt(time.Now().Unix(), 10))
		if err != nil{
			tx.Rollback()
			fmt.Print(err)
			return
		}
	
		var toDelete = []struct {
			
			Id string
			FilePath string

		}{}

		for rows.Next() {

			var id string
			var filePath string

			err = rows.Scan(&id, &filePath)
			if err != nil {
				tx.Rollback()
				fmt.Println(err)
				rows.Close()
				return
			}

			//dont wanna clutter the file with type struct
			//we cant' directly remove the rows from database because it will be locked before we call rows.Close()
			//so we just put things we want to delete into array
			toDelete = append(toDelete, struct{Id string; FilePath string}{Id: id, FilePath: filePath})
		
		}

		rows.Close()

		for _, v := range toDelete {

			_, err = tx.Exec("DELETE FROM data WHERE id = ?", v.Id)
			if err != nil {
				tx.Rollback()
				fmt.Println(err)
				return
			}

			err = os.Remove(v.FilePath)
			if err != nil {
				fmt.Println(err)
			}
	
		}

		if err := tx.Commit(); err != nil {
			fmt.Println(err)
			return;
	    }

		time.Sleep(10 * time.Second)
		
		
	}

}
