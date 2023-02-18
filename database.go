package main

import "github.com/go-pg/pg/v10"

func PostgresConnect() *pg.DB {
	options := &pg.Options{
		User:     "postgres",
		Password: "ptpit",
		Addr:     "localhost:5432",
		Database: "cats_2",
	}
	return pg.Connect(options)
}
