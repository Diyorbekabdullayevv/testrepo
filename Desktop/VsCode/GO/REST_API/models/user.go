package models

type User struct {
	Id       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
