package db

import (
	"errors"
	"net/url"

	"github.com/LyubenGeorgiev/rentals/db/models"
)

// Store represents the interface for interacting with the database
type Storage interface {
	FetchRentalByID(id int) (*models.UserAndRental, error)
	FetchRentalByQuery(queryParams url.Values) ([]*models.UserAndRental, error)
}

var ErrNotFound = errors.New("Resource was not found")
