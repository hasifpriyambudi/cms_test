package app

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hasifpriyambudi/cms_test/helpers"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {

	// Read DotEnv
	err := godotenv.Load("./.env")
	helpers.PanicError(err)

	db, err := sql.Open("mysql", os.Getenv("MYSQL_CON"))
	helpers.PanicError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
