package repository

import (
	"github.com/tawfeeq0/go_ssl_server/models"
)

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindById(uint32) (models.User, error)
	Update(uint32, models.User) (int64, error)
	Delete(uint32) (int64, error)
}
