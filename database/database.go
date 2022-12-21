package database

import (
	"github.com/Saurabhkanawade/golang-practice/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"os"
)

// making connection with pg

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect ...........\n")
		os.Exit(100)
	}
	log.Printf("Connected sucesss...\n")

	err := createSchema(db)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.Customer)(nil),
	}

	for _, mod := range models {
		err := db.Model(mod).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
