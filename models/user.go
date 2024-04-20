package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"-"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(providedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword)); err != nil {
		return err
	}

	return nil
}
