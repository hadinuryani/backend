package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB

// database config
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getEnv("DB_HOST", "127.0.0.1"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASS", ""),
		Database: getEnv("DB_NAME", "rems"),
	}
}

// koneksi database
func ConnectDB() {
	config := NewDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	db,err := sql.Open("mysql",dsn)
	if err != nil {
		log.Fatalf("Failed to create database connection:%v\n",err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database : %v\n",err)
	}
	DB = db
	log.Println("database connected successfully")
}


func getEnv(key string, defaultFalue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultFalue
}
