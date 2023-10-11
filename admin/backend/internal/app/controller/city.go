package controller

import (
	"admin/internal/app/handler"
	"admin/internal/app/middleware/service"
	"admin/internal/payload"
	"admin/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type City struct {
	s service.City
	log logrus.Logger
}

func NewCity(s service.City, log logrus.Logger) City {
	return City{s, log}
}

func (c *City) Create(w http.ResponseWriter, r *http.Request) {
	var city payload.City
	var err error

	if err = json.NewDecoder(r.Body).Decode(&city); err != nil {
		handler.HandleDecodeJsonError(w, err)
		return
	}

	if err = c.s.Create(city); err != nil {
		handler.HandlerInternalServerError(w, "Error create city")
		return
	}

	handler.HandleOkMsg(w, "")
}

func (c *City) Preview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ids, ok := vars["id"]
	if !ok {
		handler.HandleBadRequest(w, "ID is empty")
		return
	}

	id := utils.ParseInt(w, ids)

	city, err := c.s.Preview(id)
	if err != nil {
		handler.HandleNotFound(w, "")
		return
	}

	data := map[string]interface{}{
		"city": city,
	}

	handler.HandleOkData(w, data)
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

	if err := c.s.Update(id, city); err != nil {
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

	if err := c.s.Delete(id); err != nil {
		handler.HandlerInternalServerError(w, "Error delete city")
		return
	}

	handler.HandleOkMsg(w, "")
}