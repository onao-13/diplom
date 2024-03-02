package controller

import (
	"admin/internal/app/handler"
	"admin/internal/app/middleware/security"
	"admin/internal/app/payload"
	"encoding/json"
	"net/http"
)

type Auth struct {
	securityAuth security.Auth
}

func NewAuth(securityAuth security.Auth) Auth {
	return Auth{securityAuth: securityAuth}
}

func (a Auth) Login(w http.ResponseWriter, r *http.Request) {
	var auth payload.Auth
	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		handler.HandleBadRequest(w, "Ошибка декодирования JSON")
		return
	}

	token, err := a.securityAuth.Login(auth)
	if err != nil {
		handler.Forbidden(w, err.Error())
		return
	}

	tokenCookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, tokenCookie)
}
