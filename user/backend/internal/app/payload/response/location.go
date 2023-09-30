package response

type Location struct {
	Name        string `json:"name"`
	Position    string `json:"position"`
	Description string `json:"description"`
}

type LocationList struct {
	List []Location `json:"location_list"`
}
