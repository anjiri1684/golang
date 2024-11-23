package main

import (
	"database/sql"
	"log"

	"github.com/anjiri1684/ecom/cmd/api"
	"github.com/anjiri1684/ecom/config"
	"github.com/anjiri1684/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

db, err := db.NewMySQlStorage(mysql.Config{
	User: config.Envs.DBUser,
	Passwd: config.Envs.DBPassword,
	Addr: config.Envs.DBAddress,
	DBName: config.Envs.DBName,
	Net: "tcp",
	AllowNativePasswords: true,
	ParseTime: true,
})


if err != nil {
	log.Fatal(err)
}

initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func  initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DATABASE CONNECTED!!")
}