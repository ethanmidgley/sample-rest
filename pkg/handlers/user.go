// Package handlers stores handles for the api routes
package handlers

import (
	"github.com/ethanmidgley/sample-rest/pkg/db"
	"github.com/ethanmidgley/sample-rest/pkg/db/models"
	"github.com/ethanmidgley/sample-rest/pkg/user"
	"github.com/gin-gonic/gin"
)

// UH is a struct which holds a database connection as well as methods for each user handler
type UH struct {
	DB *db.DB
}

// AttachUser will take in a gin router and database and will attach the handles to the specified routes
func AttachUser(router *gin.RouterGroup, db *db.DB) {

	handles := UH{DB: db}

	router.POST("/register", handles.Register)

}

// func Login() {}

// Register is the handler for user registration
func (u *UH) Register(ctx *gin.Context) {

	var data models.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(500, gin.H{"error": "unable to parse data"})
		return
	}

	_, err := user.Register(&data, u.DB.Client)
	if err != nil {
		// ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"hell": "how are you"})
	return

}
