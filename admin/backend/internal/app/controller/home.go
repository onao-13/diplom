package controller

import (
	"admin/internal/app/errors"
	"admin/internal/app/handler"
	"admin/internal/app/payload"
	"admin/internal/app/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Home struct {
	service service.Home
}

func NewHome(service service.Home) Home {
	return Home{service: service}
}

func (l *Home) Create(w http.ResponseWriter, r *http.Request) {
	var (
		home   payload.Home
		err    error
		cityId = mux.Vars(r)["cityId"]
	)

	if err = json.NewDecoder(r.Body).Decode(&home); err != nil {
		handler.HandleDecodeJsonError(w, err)
		return
	}

	if err := l.service.Create(cityId, home); err != nil {
		handler.HandlerInternalServerError(w, "Не удалось создать дом")
		return
	}

	handler.HandleOkMsg(w, "")
}

func (l *Home) Preview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		cityId = vars["cityId"]
		homeId = vars["id"]
	)

	home, err := l.service.GetById(cityId, homeId)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			handler.HandleNotFound(w, err.Error())
			return
		default:
			handler.HandleNotFound(w, "Дом по указанному ID не существует")
			return
		}
	}

	handler.HandleOkData(w, home)
}

func (l Home) GetAll(w http.ResponseWriter, r *http.Request) {
	var cityId = mux.Vars(r)["cityId"]

	homes, err := l.service.GetByCityId(cityId)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			handler.HandleNotFound(w, err.Error())
			return
		default:
			handler.HandlerInternalServerError(w, "Ошибка сервера")
			return
		}
	}

	data := map[string]interface{}{
		"homes": homes,
	}

	handler.HandleOkData(w, data)
}

func (l *Home) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		homeId = vars["id"]
	)

	var (
		home payload.HomeUpdate
		err  error
	)

	if err = json.NewDecoder(r.Body).Decode(&home); err != nil {
		handler.HandleDecodeJsonError(w, fmt.Errorf("Ошибка декодирования JSON"))
		return
	}

	if err = l.service.Update(homeId, home); err != nil {
		handler.HandlerInternalServerError(w, "Ошибка обновления дома")
		return
	}

	handler.HandleOkMsg(w, "")
}

func (l *Home) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		cityId = vars["cityId"]
		homeId = vars["id"]
	)

	if err := l.service.Delete(cityId, homeId); err != nil {
		handler.HandlerInternalServerError(w, "Ошибка удаления дома")
		return
	}

	handler.HandleNoContent(w)
}
