package web

import (
	"fmt"
	"os"
)

type Page struct {
	Body []byte
}

const (
	Index     = "./../frontend/html/index.html"
	Home      = "./../frontend/html/home.html"
	Articles  = "./../frontend/html/articles.html"
	Locations = "./../frontend/html/locations.html"
)

func LoadPage(name string) (Page, error) {
	body, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		return Page{}, err
	}

	return Page{body}, nil
}
