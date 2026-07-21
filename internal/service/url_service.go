package service

import (
	"math/rand"
	"strings"

	"github.com/mjrdev/internal/config"
	"github.com/mjrdev/internal/handler/request"
	models "github.com/mjrdev/internal/model"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetUrl(id uint) (models.Url, error) {
	db := config.Db()

	var url models.Url

	if err := db.First(&url, "id = ?", id).Error; err != nil {
		return models.Url{}, err
	}

	return url, nil
}

func GetUrlByPath(path string) (models.Url, error) {
	db := config.Db()

	var url models.Url

	if err := db.First(&url, "path = ?", path).Error; err != nil {
		return models.Url{}, err
	}

	return url, nil
}

func ListUrl() ([]models.Url, error) {
	db := config.Db()

	var urls []models.Url

	if err := db.Find(&urls).Error; err != nil {
		return nil, err
	}

	return urls, nil
}

func CreateUrl(req request.CreateUrlRequest, userID uint) (models.Url, error) {
	db := config.Db()

	url := models.Url{
		Path:        generateRandomString(12),
		Destination: normalizeDestination(req.Url),
		UserId:      userID,
	}

	if err := db.Create(&url).Error; err != nil {
		return models.Url{}, err
	}

	return url, nil
}

func DeleteUrl(id uint) (models.Url, error) {
	db := config.Db()

	var url models.Url

	if err := db.First(&url, "id = ?", id).Error; err != nil {
		return models.Url{}, err
	}

	if err := db.Delete(&url).Error; err != nil {
		return models.Url{}, err
	}

	return url, nil
}

func normalizeDestination(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "https://" + url
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
