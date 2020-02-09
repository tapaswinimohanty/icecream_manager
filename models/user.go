package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int    `gorm:"primary_key" json:"id"`
	Username       string `gorm:"unique; not null"`
	DisplayName    string `gorm:"not null" json:"display_name"`
	HashedPassword string `gorm:"not null" json:"-"`
}

func (u *User) ComparePassword(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(plainPassword))
	if err != nil {
		return false
	}
	return true
}
