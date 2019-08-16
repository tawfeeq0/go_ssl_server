package auto

import (
	"github.com/tawfeeq0/go_ssl_server/models"
)

var users = []models.User{
	models.User{Nickname: "Tawfeeq", Email: "tawfeeq@outlook.com", Password: "123456789"},
}

var posts = []models.Post{
	{
		Title:   "Golang Tutorial",
		Content: "Golang Tutorial illustrated by Tawfeeq",
	},
}
