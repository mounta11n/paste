package main

import (

	"math/rand"
	"time"
	"strconv"
	"database/sql"
	"os"

)

//we want to avoid ambiguous characters like i, I, l, 1, etc
const charset = "abcdefghkmnpqrstwxyzABCDEFGHJKLMNPQRSTWXYZ123456789"

var seed rand.Source
var random *rand.Rand

func init(){

    seed = rand.NewSource(time.Now().UnixNano())
    random = rand.New(seed)
	
}


func GenRandFileName(basePath string, extension string) string {

    for {
        fileName := strconv.FormatInt(time.Now().UnixMilli(), 10) + GenRandString(5) + extension
        filePath := basePath + fileName

        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            return filePath
        }
    }
}

func GenRandPath(length int, db * sql.DB) string{
	
	for {
		
		randPath := GenRandString(6)
		var id string

		db.QueryRow("SELECT id FROM data WHERE id = ?", randPath).Scan(&id)

		if(id == ""){
			return randPath	
		}

	}	

}

// Function to generate a random string
func GenRandString(length int) string{
    result := make([]byte, length)
    for i := range result {
        result[i] = charset[random.Intn(len(charset))]
    }
    return string(result)
}
