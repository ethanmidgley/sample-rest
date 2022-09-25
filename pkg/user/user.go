// Package user handles everything user related
package user

import (
	"errors"
	"log"

	"github.com/ethanmidgley/sample-rest/pkg/db/models"
	"github.com/ethanmidgley/sample-rest/types"
	"golang.org/x/crypto/bcrypt"
	"xorm.io/xorm"
)

// RegisterOutput is the struct which will be returned from the register method
type RegisterOutput struct {
	User  *models.User      `json:"user"`
	Error *types.FieldError `json:"error"`
}

// func Login() {}

// Register will take in user details perform validation checks if the details pass the user will be added to the database
func Register(details *models.User, database *xorm.Engine) (*RegisterOutput, error) {

	otherUser := models.User{
		Username: details.Username,
	}
	has, err := database.Get(&otherUser)
	if err != nil {
		return nil, errors.New("failed to query database")
	}
	if has {
		return &RegisterOutput{Error: &types.FieldError{Field: "username", Message: "username has already been taken"}}, nil
	}

	if len(details.Password) < 8 {
		return &RegisterOutput{Error: &types.FieldError{Field: "password", Message: "password has to be longer 8 characters"}}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(details.Password), bcrypt.MinCost)
	if err != nil {
		return nil, errors.New("unable to secure password")
	}
	details.Password = string(hashedPassword)

	_, err = database.Insert(details)
	if err != nil {
		log.Panic(err)
	}

	return &RegisterOutput{User: details}, nil

}
