package services

import (
	"Shortener/models"
	"Shortener/repository"
	"fmt"
)

// AddUrl добавление ссылки
func AddUrl(url models.Url) {
	err := repository.AddUrl(url)
	if err != nil {
		fmt.Println(err)
	}
}

// GetLongByShort получение длинной ссылки по короткой
func GetLongByShort(shortUrl string) (string, error) {
	var a, err = Get(shortUrl)

	if err != nil {
		fmt.Println(err)
	} else {
		repository.GetByShort(shortUrl)
		return a, nil
	}
	url, err := repository.GetByShort(shortUrl)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return url.Long, nil

}

// GetByShort получений ссылки по короткой
func GetByShort(shortUrl string) (models.Url, error) {

	url, err := repository.GetByShort(shortUrl)
	if err != nil {
		fmt.Println(err)
		return models.Url{}, err
	}
	return url, nil

}
