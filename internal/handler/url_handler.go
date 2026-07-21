package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mjrdev/internal/handler/request"
	"github.com/mjrdev/internal/middleware"
	"github.com/mjrdev/internal/service"
	"github.com/mjrdev/pkg/response"
	"github.com/mjrdev/pkg/validator"
	"gorm.io/gorm"
)

func StoreUrl(w http.ResponseWriter, r *http.Request) {
	req, ok := validator.Validate[request.CreateUrlRequest](w, r)
	if !ok {
		return
	}

	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	url, err := service.CreateUrl(req, userID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, url)
}

func ShowUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "short_url")

	url, err := service.GetUrlByPath(shortUrl)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(w, http.StatusNotFound, "url não encontrada")
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, url)
}

func ListUrl(w http.ResponseWriter, r *http.Request) {
	url, err := service.ListUrl()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(w, http.StatusNotFound, "usuário não encontrado")
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, url)
}

func DeleteUrl(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "id inválido")
		return
	}

	url, err := service.DeleteUrl(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(w, http.StatusNotFound, "url não encontrada")
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, url)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "path")

	url, err := service.GetUrlByPath(path)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(w, http.StatusNotFound, "url não encontrada")
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.Redirect(w, r, url.Destination, http.StatusFound)
}
