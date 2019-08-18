package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/tawfeeq0/go_ssl_server/database"
	"github.com/tawfeeq0/go_ssl_server/models"
	"github.com/tawfeeq0/go_ssl_server/security"
	"github.com/tawfeeq0/go_ssl_server/utils/channels"
)

func SignIn(email, password string) (string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		err = db.Debug().Model(&models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return CreateToken(user.ID)
	}
	return "", err
}
