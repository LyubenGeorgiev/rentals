package internal

import (
	"net/http"

	"github.com/LyubenGeorgiev/rentals/db"
)

type RentalsHandler struct {
	Storage db.Storage
}

func NewRentalsHandler(storage db.Storage) *RentalsHandler {
	return &RentalsHandler{
		Storage: storage,
	}
}

func (rh *RentalsHandler) GetRentalByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ID"))
}

func (rh *RentalsHandler) GetRentalByQuery(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Query"))
}
