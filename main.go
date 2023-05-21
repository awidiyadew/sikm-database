package main

import (
	"fmt"
	"log"
	"sikm-marketplace-db/db"
	"sikm-marketplace-db/repository"
)

func main() {
	db, err := db.Connect(db.DBCredential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "P@ssw0rd",
		DatabaseName: "marketplace-demo",
		Port:         5432,
	})

	if err != nil && db.Ping() != nil {
		log.Fatal("failed to connect db", err)
	}

	fmt.Println("connected to DB")

	_ = repository.NewProductRepo(db)
}
