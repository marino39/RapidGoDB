package main

import (
	"RapidGoDB/db"
	"log"
)

func main() {
	rgdb := db.GetRapidGoDB(db.RGDBConfig)
	rgdb.Set("Test", "Hello world!")
	log.Println(rgdb.Get("Test"))
}
