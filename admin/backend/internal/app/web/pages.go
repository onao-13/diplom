package web

import (
	"fmt"
	"os"
)

type Page struct {
	Body []byte
}

const (
	pagesPath = "./frontend/web"
	Auth      = pagesPath + "/auth.html"
	Panel     = pagesPath + "/panel.html"
	Calls     = pagesPath + "/calls.html"
	Cities    = pagesPath + "/cities.html"
	Homes     = pagesPath + "/homes.html"
	Home      = pagesPath + "/home.html"
)

func LoadPage(name string) (Page, error) {
	body, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		return Page{}, err
	}

	return Page{body}, nil
}
