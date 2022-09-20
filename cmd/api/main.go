// Package main here puts all the compnents together creating our rest api
package main

import (
	"github.com/ethanmidgley/sample-rest/pkg/api"
	"github.com/ethanmidgley/sample-rest/pkg/auth"
	"github.com/ethanmidgley/sample-rest/pkg/db"
	"github.com/ethanmidgley/sample-rest/pkg/server"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	database := db.Init()

	r := api.CreateRoutes(database)
	r.Use(auth.Middleware(database))

	s := server.Create(r)
	s.Run()

}
