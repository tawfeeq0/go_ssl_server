package auto

import (
	"log"

	"github.com/tawfeeq0/go_ssl_server/database"
	"github.com/tawfeeq0/go_ssl_server/models"
	"github.com/tawfeeq0/go_ssl_server/utils/console"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal(err)
		}
		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatal(err)
		}

		console.Pretty(users[i])
		console.Pretty(posts[i])
	}
}
