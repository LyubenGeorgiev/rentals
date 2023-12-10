package models

type price struct {
	Day int64 `json:"day"`
}

type location struct {
	City    string  `json:"city"`
	State   string  `json:"state"`
	Zip     string  `json:"zip"`
	Country string  `json:"country"`
	Lat     float32 `json:"lat"`
	Lng     float32 `json:"lng"`
}

type user struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type RentalResponse struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Type            string   `json:"type"`
	Make            string   `json:"make"`
	Model           string   `json:"model"`
	Year            int      `json:"year"`
	Length          float64  `json:"length"`
	Sleeps          int      `json:"sleeps"`
	PrimaryImageURL string   `json:"primary_image_url"`
	Price           price    `json:"price"`
	Location        location `json:"location"`
	User            user     `json:"user"`
}
