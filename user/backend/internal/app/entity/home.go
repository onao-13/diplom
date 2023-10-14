package entity

type Home struct {
	Id            int64
	Name          string
	Price         int64
	HomeFeatureId int64
	CityId        int64
}

type City struct {
	Id   int64
	Name string
}
