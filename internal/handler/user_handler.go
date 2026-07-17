package handler

import (
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/mjrdev/internal/handler/request"
	"github.com/mjrdev/internal/middleware"
	"github.com/mjrdev/internal/service"
	"github.com/mjrdev/pkg/response"
	"github.com/mjrdev/pkg/validator"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {

	req, ok := validator.Validate[request.CreateUserRequest](w, r)
	if !ok {
		return
	}

	user, err := service.UserServiceCreate(req)

	if err != nil {
		if errors.Is(err, service.ErrEmailTaken) {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Context().Value(middleware.UserIDKey).(string)

	user, err := service.UserServiceGet(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(w, http.StatusNotFound, "usuário não encontrado")
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, user)
}
