package controller

import (
	"fmt"
	"net/http"
	"os"

	// "backend/app/handler"
)

type Location struct{}

func NewLocation() Location {
	return Location{}
}

func (l *Location) Get(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("./json/locations.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
	// handler.HandleOk(w, data)
}
