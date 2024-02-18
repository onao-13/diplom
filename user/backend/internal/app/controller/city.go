package controller

import (
	"backend/internal/app/handler"
	"backend/internal/app/service"
	"net/http"
)

type City struct {
	s service.City
}

func NewCity(s service.City) City {
	return City{s: s}
}

func (c City) GetAll(w http.ResponseWriter, r *http.Request) {
	cities, err := c.s.GetAll()
	if err != nil {
		handler.HandlerInternalServerError(w, "Ошибка получения городов")
		return
	}

	body := map[string]interface{}{
		"cities": cities,
	}

	handler.HandleOkData(w, body)
}
