package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/rushikeshkandekar/model"
	"log"
	"os"
)

func ConnectDB() *pg.DB {

	log.Println("connecting to the database...")

	options := &pg.Options{
		Database: os.Getenv("DATABASE_NAME"),
		Addr:     os.Getenv("DATABASE_ADDRESS"),
		Password: os.Getenv("DATABASE_PASS"),
		User:     os.Getenv("DATABASE_USER"),
	}

	var db *pg.DB = pg.Connect(options)

	if db == nil {
		log.Fatal("failed to connect database")
		os.Exit(100)
	}

	_ = createSchema(db)

	log.Println("Connnected success to the database..........")

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
