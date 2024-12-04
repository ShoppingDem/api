package models

type User struct {
	ID           int    `json:"id"`
	PhoneNumber  string `json:"phoneNumber" binding:"required"`
	PasswordHash string `json:"passwordHash" binding:"required"`
}

type UserRegistration struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required,min=6"`
}

var Users []User
var NextID = 1
