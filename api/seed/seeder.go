package seed

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/charleschrist21/golang_first/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "admin",
		Email:    "admin@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello World 1",
	},
	models.Post{
		Title:   "title 2",
		Content: "Hello World 2",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
