// Package models holds all the types required for the database
package models

// User is the basic structure of how a user account which will be stored in the database
type User struct {
	ID       int64  `xorm:"pk autoincr"`
	Username string `json:"username"`
	Password string `json:"password"`
}
