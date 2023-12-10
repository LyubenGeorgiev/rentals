package models

type Rental struct {
	ID              int     `db:"id"`
	Name            string  `db:"name"`
	Type            string  `db:"type"`
	Description     string  `db:"description"`
	Sleeps          int     `db:"sleeps"`
	PricePerDay     int64   `db:"price_per_day"`
	HomeCity        string  `db:"home_city"`
	HomeState       string  `db:"home_state"`
	HomeZip         string  `db:"home_zip"`
	HomeCountry     string  `db:"home_country"`
	VehicleMake     string  `db:"vehicle_make"`
	VehicleModel    string  `db:"vehicle_model"`
	VehicleYear     int     `db:"vehicle_year"`
	VehicleLength   float64 `db:"vehicle_length"`
	Lat             float32 `db:"lat"`
	Lng             float32 `db:"lng"`
	PrimaryImageURL string  `db:"primary_image_url"`
}
