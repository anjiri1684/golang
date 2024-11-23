package main

import (
	"log"
	"os"

	"github.com/anjiri1684/ecom/config"
	"github.com/anjiri1684/ecom/db"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	db, err := db.NewMySQlStorage(mysqlCfg.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
		
	}


	m , err := migrate.NewWithDatabaseInstance(
		"file://cms/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
	}
	if cmd == "down"{
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}