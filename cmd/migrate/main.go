package main

import (
	"github.com/taslav8/wb_proj1/internal/database"
	"github.com/taslav8/wb_proj1/internal/migrate"
)

func main() {
	db, err := database.SetConfig("configuration.json")
	if err != nil {
		panic(err)
	}
	db.Open()
	defer db.Close()
	if err = migrate.CreateSchema(db); err != nil {
		panic(err)
	}
}
