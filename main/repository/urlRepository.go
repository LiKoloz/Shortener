package repository

import (
	"Shortener/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error = nil

// GetByShort получение по короткой ссылке из бд
func GetByShort(shortUrl string) (models.Url, error) {
	if db == nil {
		dsn := "host=postgresql user=demo_user password=user_password dbname=url_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return models.Url{}, err
		}

	}

	db.AutoMigrate(&models.Url{})

	url := models.Url{}

	db.Where("Short = ?", shortUrl).First(&url)
	url.Transitions++

	db.Save(&url)
	return url, nil
}

// AddUrl добавление нового url в бд
func AddUrl(url models.Url) error {
	if db == nil {
		dsn := "host=postgresql user=demo_user password=user_password dbname=url_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}

	}

	db.AutoMigrate(&models.Url{})

	result := db.Create(&url)

	return result.Error
}
