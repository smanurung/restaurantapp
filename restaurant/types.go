package restaurant

type Point struct {
	Lat, Long int
}

type RatingType int

type Restaurant struct {
	ID          int
	Name        string
	CuisineType string
	Location    string
	Rating      RatingType
}
