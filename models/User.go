package models

import (
	"time"

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
