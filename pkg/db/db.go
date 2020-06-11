package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

type DB struct {
	*gorm.DB
}

func Open() DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	return DB{db}
}

func (db *DB) SyncSchema() {
	//db.AutoMigrate(&document.Document{}, &task.Task{}, &status.Status{}, &errorMessage.ErrorMessage{})
}

func (db *DB) DropSchema() {
	//db.DropTableIfExists(&document.Document{}, &task.Task{}, &status.Status{}, &errorMessage.ErrorMessage{})
}
