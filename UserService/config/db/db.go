package config

import (
	"database/sql"
	"fmt"
	env "userservice/config/env"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASS", "root")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("DB_Addr", "127.0.0.1:3306")
	cfg.DBName = env.GetString("DBName", "userService")

	fmt.Println("Connecting to database:", cfg.DBName, cfg.FormatDSN())

	// Open database connection
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}

	fmt.Println("trying to connect the database....")

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Error pinging th database:", pingErr)
		return nil, pingErr
	}
	fmt.Println("Connected to database successfully:", cfg.DBName)

	return db, nil
}
