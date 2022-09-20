package api

import (
	"github.com/ethanmidgley/sample-rest/pkg/db"
	"github.com/ethanmidgley/sample-rest/pkg/handlers"
	"github.com/gin-gonic/gin"
)

// CreateRoutes will take in a database context and will return a gin engine ready to be passed to a server
func CreateRoutes(database *db.DB) *gin.Engine {

	r := gin.Default()

	a := r.Group("/user")
	handlers.AttachUser(a, database)

	return r

}
