package security

import (
	"admin/internal/app/database"
	"admin/internal/app/errors"
	"admin/internal/app/payload"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	log    *logrus.Logger
	db     database.Auth
	tokens []string
}

func NewAuth(db database.Auth, log *logrus.Logger) Auth {
	return Auth{log: log, db: db}
}

func (a Auth) Login(auth payload.Auth) (token string, err error) {
	if !a.db.Login(auth) {
		return "", fmt.Errorf("ошибка логина или пароля")
	}

	token = a.generateToken()

	a.log.Infoln("Пользователь авторизован")

	a.tokens = append(a.tokens, token)

	return
}

func (a Auth) IsAuth(token string) (access error) {
	for _, createdToken := range a.tokens {
		if token == createdToken {
			access = &errors.ErrUnauth{}
		}
	}

	return
}

func (a Auth) generateToken() (token string) {
	apiKeyByte := []byte(uuid.New().String())

	var hash = sha512.New()
	hash.Write(apiKeyByte)

	tokenSum := hash.Sum(nil)
	token = hex.EncodeToString(tokenSum)

	return
}
