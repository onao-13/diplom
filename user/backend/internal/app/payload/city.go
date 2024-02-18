package payload

type (
	City struct {
		Id    string  `json:"id"`
		Name  string  `json:"name"`
		Homes []*Home `json:"homes"`
	}
)
