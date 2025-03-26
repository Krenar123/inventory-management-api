package models

import "gorm.io/gorm"

// Admins must register and log in via JWT authentication (email/password) to perform write operations.

// Creating Admin model with email, password(which needs to be a hash but will look after to this) and also timestamp
// models/admin.go

type Admin struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}