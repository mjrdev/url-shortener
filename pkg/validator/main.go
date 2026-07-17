package validator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mjrdev/pkg/response"
)

var validate = validator.New()

var tagMessages = map[string]string{
	"required": "campo obrigatório",
	"email":    "e-mail inválido",
	"min":      "tamanho mínimo de %s caracteres",
	"max":      "tamanho máximo de %s caracteres",
}

func fieldMessage(fe validator.FieldError) string {
	msg, ok := tagMessages[fe.Tag()]
	if !ok {
		return fmt.Sprintf("valor inválido para a regra '%s'", fe.Tag())
	}
	if fe.Param() != "" {
		return fmt.Sprintf(msg, fe.Param())
	}
	return msg
}

func Validate[T any](w http.ResponseWriter, r *http.Request) (T, bool) {
	var req T

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "JSON inválido")
		return req, false
	}

	if err := validate.Struct(req); err != nil {
		fields := make(map[string]string)
		for _, fe := range err.(validator.ValidationErrors) {
			fields[strings.ToLower(fe.Field())] = fieldMessage(fe)
		}
		response.ValidationError(w, fields)
		return req, false
	}

	return req, true
}
