package main

import (
	"log"
	"sikm-marketplace-db/db"
	"sikm-marketplace-db/repository"
)

func main() {
	db, err := db.Connect(db.DBCredential{
		Host:         "",
		Username:     "",
		Password:     "",
		DatabaseName: "",
		Port:         5432,
	})

	if err != nil {
		log.Fatal("failed to connect db", err)
	}

	_ = repository.NewProductRepo(db)
}
