package response

type City struct {
	Name string `json:"name"`
}

type Home struct {
	Name  string `json:"name"`
	Price uint64 `json:"price"`
}
