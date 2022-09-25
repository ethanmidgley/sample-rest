// Package handlers stores handles for the api routes
package handlers

import (
	"github.com/ethanmidgley/sample-rest/pkg/db"
	"github.com/ethanmidgley/sample-rest/pkg/db/models"
	"github.com/ethanmidgley/sample-rest/pkg/user"
	"github.com/gin-contrib/sessions"
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

// func (u *UH) Login(ctx *gin.Context) {
// }

// TODO: this is how to implement the getting of a user id from the session and checking if it is set and casting it to int64
// if tmp := session.Get("user-id"); tmp != nil {
//	we have a userID
// 	userID = tmp.(int64)
// } else {
//	we do not have a userID
// 	return
// }

// Register is the handler for user registration
func (u *UH) Register(ctx *gin.Context) {

	var data models.User
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(500, gin.H{"error": "unable to parse data"})
		return
	}

	session := sessions.Default(ctx)

	output, err := user.Register(&data, u.DB.Client)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	session.Set("user-id", output.User.ID)
	session.Save()

	ctx.JSON(200, output)
	return

}
