package models

type UserAndRental struct {
	User   User   `db:"users"`   // tag as the "users" table
	Rental Rental `db:"rentals"` // tag as "rentals" table
}
