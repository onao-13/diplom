package controller

import (
	"admin/internal/app/errors"
	"admin/internal/app/handler"
	"admin/internal/app/payload"
	"admin/internal/app/service"
	"admin/internal/app/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type City struct {
	service service.City
}

func NewCity(service service.City) City {
	return City{service: service}
}

func (c *City) Create(w http.ResponseWriter, r *http.Request) {
	var city payload.City
	var err error

	if err = json.NewDecoder(r.Body).Decode(&city); err != nil {
		handler.HandleDecodeJsonError(w, err)
		return
	}

	if err = c.service.Create(city); err != nil {
		handler.HandlerInternalServerError(w, "Error create city")
		return
	}

	handler.HandleOkMsg(w, "")
}

func (c City) GetAll(w http.ResponseWriter, r *http.Request) {
	cities, err := c.service.GetAll()
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			handler.HandleNotFound(w, err.Error())
			return
		default:
			handler.HandlerInternalServerError(w, "Ошибка получения городов")
			return
		}
	}

	data := map[string]interface{}{
		"cities": cities,
	}

	handler.HandleOkData(w, data)
}

func (c *City) Preview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ids, ok := vars["id"]
	if !ok {
		handler.HandleBadRequest(w, "ID is empty")
		return
	}

	id := utils.ParseInt(w, ids)

	city, err := c.service.Preview(id)
	if err != nil {
		handler.HandleNotFound(w, "")
		return
	}

	handler.HandleOkData(w, city)
}

func (c *City) Update(w http.ResponseWriter, r *http.Request) {
	var city payload.City

	vars := mux.Vars(r)
	ids, ok := vars["id"]
	if !ok {
		handler.HandleBadRequest(w, "ID is empty")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&city); err != nil {
		handler.HandleDecodeJsonError(w, err)
		return
	}

	id := utils.ParseInt(w, ids)

	if err := c.service.Update(id, city); err != nil {
		handler.HandlerInternalServerError(w, "Error update city")
		return
	}

	handler.HandleOkMsg(w, "")
}

func (c *City) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids, ok := vars["id"]
	if !ok {
		handler.HandleBadRequest(w, "ID is empty")
		return
	}

	id := utils.ParseInt(w, ids)

	if err := c.service.Delete(id); err != nil {
		handler.HandlerInternalServerError(w, "Error delete city")
		return
	}

	handler.HandleOkMsg(w, "")
}
