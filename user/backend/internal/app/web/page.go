package web

import "os"

type Page struct {
	Body []byte
}

const (
	Index     = "./html/index.html"
	Home      = "./html/home.html"
	Articles  = "./html/articles.html"
	Locations = "./html/locations.html"
)

func LoadPage(name string) (Page, error) {
	body, err := os.ReadFile(name)
	if err != nil {
		return Page{}, err
	}

	return Page{body}, nil
}
