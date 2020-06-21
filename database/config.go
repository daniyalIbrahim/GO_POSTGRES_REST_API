package database

import (
	"fastReturns/models"
	pg "github.com/go-pg/pg"
	"log"
	"os"
)

func Connect() *pg.DB {

	opts := &pg.Options{
		User:     os.Getenv("User"),
		Password: os.Getenv("Password"),
		Addr:     os.Getenv("Addr"),
		Database: os.Getenv("Database"),
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	log.Printf("Connection to database successful.\n")
	models.CreateCustomerTable(db)
	models.CreateSupplierTable(db)
	models.CreateConsignmentTable(db)
	models.CreateVanTable(db)
	return db

}
