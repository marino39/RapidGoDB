package main

import (
	"RapidGoDB/db"
	"log"
)

func main() {
	db := db.GetRapidGoDB(db.RGDBConfig)
	db.Set("Test", "Hello world!")
	log.Println(db.Get("Test"))
}
