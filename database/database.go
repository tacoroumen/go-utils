package database

import (
	"database/sql"
	"fmt"
	//"time"
	"beemer/logger"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUsername   = "Admin"
	dbPassword   = "Afslag01!"
	dbHost       = "127.0.0.1"
	dbPort       = 3306
	dbName       = "api-keys"
	dbTLSEnabled = false
)

func InitializeDB() (*sql.DB, error) {
	//currentTime := time.Now()
	//currentDate := currentTime.Format("2006-01-02")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=%t", dbUsername, dbPassword, dbHost, dbPort, dbName, dbTLSEnabled)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.LogErrorToFile("", err)
		return nil, err
	}

	// Ping the database to check the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	// Initialize the database tables and data
	//if err := initializeDatabaseData(db, currentDate); err != nil {
	//	db.Close()
	//	return nil, err
	//}

	return db, nil
}


func InitializeDatabaseData(db *sql.DB, currentDate string) error {
	// ... (existing code)
	return nil
}
