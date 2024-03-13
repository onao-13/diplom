package payload

type (
	Home struct {
		Id               string      `json:"id"`
		Name             string      `json:"name"`
		Street           string      `json:"street"`
		Price            string      `json:"price"`
		Description      string      `json:"description"`
		CityId           string      `json:"city_id"`
		Images           []HomeImage `json:"images"`
		Transports       string      `json:"transports"`
		PopularLocations string      `json:"popular_locations"`
		Layout           string      `json:"layout"`
		GreenZone        string      `json:"green_zone"`
		Infrastructure   string      `json:"infrastructure"`
		Events           string      `json:"events"`
		Schools          string      `json:"schools"`
	}
	HomeImage struct {
		Id  string `json:"id"`
		URL string `json:"url"`
	}
	HomeTransport struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
	HomePopularLocation struct {
		Id      int64  `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}
)
