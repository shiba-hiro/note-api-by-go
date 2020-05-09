package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func OpenDbConnection() (*gorm.DB, error) {
	accessInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		getEnv("DB_USER", "note_db_user"),
		getEnv("DB_PASS", "note_db_pass"),
		getEnv("DB_HOST", "127.0.0.1"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "note_db"),
	)
	db, err := gorm.Open("mysql", accessInfo)
	if err != nil {
		log.Fatalf("DB connection error: %v\n", err)
		return db, err
	}
	db.LogMode(getEnv("DB_LOG_MODE", "true") == "true")
	DB = db
	return db, nil
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
