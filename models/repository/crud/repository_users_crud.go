package crud

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tawfeeq0/go_ssl_server/models"
	"github.com/tawfeeq0/go_ssl_server/utils"
)

type repositoryUsersCRUD struct {
	db *gorm.DB
}

/*
ERROR
func (r *repositoryUsersCRUD) NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD{db}
}
*/
func NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD{db}
}

func (r *repositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if utils.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

func (r *repositoryUsersCRUD) FindAll() ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if utils.OK(done) {
		return users, nil
	}
	return []models.User{}, err
}

func (r *repositoryUsersCRUD) FindById(userId uint32) (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("id = ?", userId).Find(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if utils.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

func (r *repositoryUsersCRUD) Update(userId uint32, user models.User) (int64, error) {
	var rs *gorm.DB
	//user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", userId).Take(&models.User{}).UpdateColumns(
			map[string]interface{}{
				"nickname":   user.Nickname,
				"email":      user.Email,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)

	if utils.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

func (r *repositoryUsersCRUD) Delete(userId uint32) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", userId).Take(&models.User{}).Delete(&models.User{})
		ch <- true
	}(done)

	if utils.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}
