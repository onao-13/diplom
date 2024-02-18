package payload

type SendManagerCall struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	HomeId string `json:"home_id"`
}
