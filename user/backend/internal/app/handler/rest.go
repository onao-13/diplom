package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// HandleOkMsg function   Обертка для отправки ответа со статусом 200 и сообщением
func HandleOkMsg(w http.ResponseWriter, msg string) {
	handle(http.StatusOK, w, nil)
}

// HandleOk function     Обертка для отправки ответа со статусом 200 и телом
func HandleOkData(w http.ResponseWriter, data interface{}) {
	handle(http.StatusOK, w, data)
}

// HandleBadRequest function     Обертка для отправки ответа со статусом 400
func HandleBadRequest(w http.ResponseWriter, err string) {
	//res := createShortRes("error", err)
	handle(http.StatusBadRequest, w, nil)
}

// HandleNotFound function     Обертка для отправки ответа со статусом 404
func HandleNotFound(w http.ResponseWriter, err string) {
	handle(http.StatusNotFound, w, nil)
}

// HandlerInternalServerError function     Обертка для отправки ответа со статусом 500
func HandlerInternalServerError(w http.ResponseWriter, err string) {
	res := map[string]interface{}{
		"error": err,
	}
	handle(http.StatusInternalServerError, w, res)
}

// setContentType установка заголовка Content-Type - application/json
func setContentType(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

// handle function    функция для создания ответа и отправки ее клиенту
func handle(code int, w http.ResponseWriter, data interface{}) {
	setContentType(&w)
	w.WriteHeader(code)
	d, err := json.Marshal(data)
	if err != nil {
		logger.Error("Ошибка: %s", err.Error())
	}

	if _, err := w.Write(d); err != nil {
		logger.Error("Ошибка: %s", err.Error())
	}
}
