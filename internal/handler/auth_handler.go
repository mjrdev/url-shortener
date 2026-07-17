package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mjrdev/internal/config"
	"github.com/mjrdev/internal/handler/request"
	responser "github.com/mjrdev/internal/handler/response"
	"github.com/mjrdev/internal/middleware"
	models "github.com/mjrdev/internal/model"
	"github.com/mjrdev/pkg/bcrypt"
	"github.com/mjrdev/pkg/response"
	"github.com/mjrdev/pkg/validator"
	"gorm.io/gorm"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	db := config.Db()

	req, ok := validator.Validate[request.CreateUserRequest](w, r)
	if !ok {
		return
	}
	var user models.User
	if err := db.First(&user, "email = ?", req.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(w, http.StatusNotFound, "usuário não encontrado")
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	verify := bcrypt.VerifyPassword(req.Password, user.PasswordHash)

	if !verify {
		response.Error(w, http.StatusInternalServerError, "Email ou senha incorretos.")
		return
	}

	s := fmt.Sprintf("%d", user.ID)
	token, err := middleware.GenerateToken(s)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, &responser.AuthResponse{
		Token: token,
	})
}
