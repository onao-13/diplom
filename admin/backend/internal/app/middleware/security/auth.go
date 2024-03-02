package security

import (
	"admin/internal/app/errors"
	"admin/internal/app/payload"
	"admin/internal/config"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	log    *logrus.Logger
	root   config.RootUser
	tokens []string
}

func NewAuth(root config.RootUser, log *logrus.Logger) Auth {
	return Auth{log: log, root: root}
}

func (a Auth) Login(auth payload.Auth) (token string, err error) {
	if auth.Username != a.root.Username {
		return "", fmt.Errorf("неверное имя пользователя")
	}

	if auth.Password != a.root.Password {
		return "", fmt.Errorf("неверный пароль")
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
