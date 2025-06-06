package main

import (
	"database/sql"
	"log"

	"github.com/CuongDepay/go-ecom/cmd/api"
	"github.com/CuongDepay/go-ecom/config"
	"github.com/CuongDepay/go-ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatalf("failed to open mysql connection: %v", err)
	}

	initStorage(db)

	server := api.NewAPIServer(config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("failed to ping mysql: %v", err)
	}

	log.Println("connected to mysql")
}
