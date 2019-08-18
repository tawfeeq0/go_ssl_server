package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"

	"github.com/tawfeeq0/go_ssl_server/security"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:20;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:5-;not null;uinque" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	Post      []Post    `gorm:"foreignkey:AuthorID" json:"posts:omitempty"`
}

func (user *User) BeforeSave() error {
	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Prepare() {
	user.ID = 0
	user.Nickname = html.EscapeString(strings.TrimSpace(user.Nickname))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if user.Nickname == "" {
			return errors.New("Nickname is required")
		}
		if user.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Email is invalid")
		}
		return nil
	case "login":
		if user.Email == "" {
			return errors.New("Email is required")
		}
		if user.Password == "" {
			return errors.New("Password is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Email is invalid")
		}
		return nil
	default:
		if user.Nickname == "" {
			return errors.New("Nickname is required")
		}
		if user.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Email is invalid")
		}
		return nil
	}
}
