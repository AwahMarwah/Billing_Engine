package main

import (
	"billing_engine/database"
	"billing_engine/router"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	log.Println("Starting server...")

	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.SqlDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err = router.Run(db); err != nil {
		log.Fatal(err)
	}
}
