package response

import "net/http"

type Article struct {
	Id      int `json:id`
	Name    string `json:name`
	Content string `json:content`
	Images  []http.File `json:images`
}