package payload

type ManagerCall struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Number   string `json:"number"`
	HomeName string `json:"home_name"`
}
