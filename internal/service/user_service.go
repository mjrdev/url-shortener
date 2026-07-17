package service

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/mjrdev/internal/config"
	"github.com/mjrdev/internal/handler/request"
	models "github.com/mjrdev/internal/model"
	"github.com/mjrdev/pkg/bcrypt"
)

var ErrEmailTaken = errors.New("e-mail já está em uso")

func UserServiceGetAll() ([]models.User, error) {
	db := config.Db()

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func UserServiceGet(id string) (models.User, error) {
	db := config.Db()

	var user models.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UserServiceCreate(req request.CreateUserRequest) (models.User, error) {
	db := config.Db()

	password, err := bcrypt.HashPassword(req.Password)

	if err != nil {
		return models.User{}, err
	}

	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: password,
	}

	if err := db.Create(&user).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.User{}, ErrEmailTaken
		}
		return models.User{}, err
	}

	return *user, nil
}
