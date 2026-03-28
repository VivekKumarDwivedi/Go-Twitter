package config

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"TweetService/config/env"
)

func SetupDB() (*sql.DB, error){
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASS", "root")
	cfg.DBName = env.GetString("DB_NAME", "userservice")
	cfg.Addr = env.GetString("DB_ADDR", "localhost:3306")
	cfg.Net = env.GetString("DB_NET","tcp")

	fmt.Println("Connecting to database:",cfg.DB_NAME,cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error connecting the database:",err)
		return nil,err
	}

	fmt.Println("trying to connect the database...")
	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Error connecting the database:",pingErr)
		return nil,pingErr
	}

	fmt.Println("Database connected successfully",cfg.DB_NAME)
	return db,nil
}
