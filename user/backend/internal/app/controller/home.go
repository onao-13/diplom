package controller

import (
	"net/http"
	"os"
)

type Home struct {
	
}

func NewHome() Home {
	return Home{}
}

func (*Home) Data(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("./json/home.json")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(data)
}