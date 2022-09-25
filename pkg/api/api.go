package api

import (
	"os"

	"github.com/ethanmidgley/sample-rest/pkg/db"
	"github.com/ethanmidgley/sample-rest/pkg/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// CreateRoutes will take in a database context and will return a gin engine ready to be passed to a server
func CreateRoutes(database *db.DB) *gin.Engine {

	r := gin.Default()

	store := cookie.NewStore([]byte(os.Getenv("SESSIONSECRET")))
	r.Use(sessions.Sessions("sid", store))

	a := r.Group("/user")
	handlers.AttachUser(a, database)

	return r

}
