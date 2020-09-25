package routers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //database

func init() {

	connStr := "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=example"

	conn, err := gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error On Connect DB", err)
	}
	db = conn
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
