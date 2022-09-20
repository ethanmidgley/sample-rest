// Package db handles everything related to the database
package db

import (
	"log"

	"github.com/ethanmidgley/sample-rest/pkg/db/models"
	// need to import sqlite driver for xorm to work
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

// DB is a struct which holds a xorm database connection
type DB struct {
	Client *xorm.Engine
}

// Init when called will attempt to connect to the and will either log an error or return a DB struct.
// It will also handle all structs needed to be synced to the database
func Init() *DB {
	engine, err := xorm.NewEngine("sqlite3", "./db.sql")
	if err != nil {
		log.Panic(err)
	}
	engine.Sync(new(models.User))
	engine.Sync(new(models.Auth))

	return &DB{Client: engine}
}
