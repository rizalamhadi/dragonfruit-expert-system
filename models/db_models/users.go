package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// BeforeCreate hash password
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	password := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(password, 11)
	if err != nil {
		log.Println(err)
		return err
	}

	tx.Model(u).Updates(User{Password: string(hash), ID: u.ID})

	return
}