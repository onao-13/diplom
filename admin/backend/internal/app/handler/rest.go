package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// HandleDecodeJsonError function     Обертка для отправки ответа со статусом 500 и записью в логи ошибки об декодировании
func HandleDecodeJsonError(w http.ResponseWriter, err error) {
	HandlerInternalServerError(w, "")
	logger.Error("Ошибка декодирования json: %s", err.Error())
}

// HandleOk function     Обертка для отправки ответа со статусом 200 и телом
func HandleOkData(w http.ResponseWriter, data map[string]interface{}) {
	handle(http.StatusOK, w, data)
}

// HandleOkMsg function   Обертка для отправки ответа со статусом 200 и сообщением
func HandleOkMsg(w http.ResponseWriter, msg string) {
	res := createShortRes("response", msg)
	handle(http.StatusOK, w, res)
}

// HandleCreate function Обертка для отправки ответа со статусом 201
func HandleCreate(w http.ResponseWriter, msg string) {
	res := createShortRes("response", msg)
	handle(http.StatusCreated, w, res)
}

func HandleNoContent(w http.ResponseWriter) {
	handle(http.StatusNoContent, w, nil)
}

// HandleDelete function     Обертка для отправки ответа со статусом 204
func HandleDelete(w http.ResponseWriter, msg string) {
	res := createShortRes("response", msg)
	handle(http.StatusNoContent, w, res)
}

// HandleConflict function     Обертка для отправки ответа со статусом 409
func HandleConflict(w http.ResponseWriter, msg string) {
	res := createShortRes("response", msg)
	handle(http.StatusConflict, w, res)
}

// HandleBadRequest function     Обертка для отправки ответа со статусом 400
func HandleBadRequest(w http.ResponseWriter, err string) {
	res := createShortRes("error", err)
	handle(http.StatusBadRequest, w, res)
}

// HandleNotFound function     Обертка для отправки ответа со статусом 404
func HandleNotFound(w http.ResponseWriter, err string) {
	res := createShortRes("error", err)
	handle(http.StatusNotFound, w, res)
}

// HandlerInternalServerError function     Обертка для отправки ответа со статусом 500
func HandlerInternalServerError(w http.ResponseWriter, err string) {
	res := createShortRes("error", err)
	handle(http.StatusInternalServerError, w, res)
}

// handle function    функция для создания ответа и отправки ее клиенту
func handle(code int, w http.ResponseWriter, data map[string]interface{}) {
	setContentType(&w)
	w.WriteHeader(code)
	if data != nil {
		d, err := json.Marshal(data)
		if err != nil {
			logger.Error("Ошибка: %s", err.Error())
		}

		if _, err := w.Write(d); err != nil {
			logger.Error("Ошибка: %s", err.Error())
		}
	} else {
		w.Write([]byte{})
	}
}

// setContentType установка заголовка Content-Type - application/json
func setContentType(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

// createShortRes создание короткого ответа для отправки
func createShortRes(name string, msg string) map[string]interface{} {
	r := make(map[string]interface{}, 1)
	r[name] = msg
	return r
}
