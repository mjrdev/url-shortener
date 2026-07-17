package handler

import (
	"net/http"

	"github.com/mjrdev/pkg/response"
)

func UrlCreate(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, nil)
}

func UrlGet(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, nil)
}
