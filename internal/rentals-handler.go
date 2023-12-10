package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LyubenGeorgiev/rentals/db"
	dbmodels "github.com/LyubenGeorgiev/rentals/db/models"
	"github.com/LyubenGeorgiev/rentals/models"
	"github.com/gorilla/mux"
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
	// Get the rental id
	vars := mux.Vars(r)
	rentalIDStr := vars["id"]

	// Convert the rental id to int
	rentalID, err := strconv.Atoi(rentalIDStr)
	if err != nil {
		http.Error(w, "Invalid rental ID", http.StatusBadRequest)
		return
	}

	// Fetch rental data using ID
	dbResponse, err := rh.Storage.FetchRentalByID(rentalID)
	if err != nil {
		if err == db.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Convert to rental response
	rentalResponse := convertDBResponseToRentalResponse(dbResponse)

	// Marshal the rental struct to JSON
	response, err := json.Marshal(rentalResponse)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (rh *RentalsHandler) GetRentalByQuery(w http.ResponseWriter, r *http.Request) {
	// Fetch rental data using query
	dbResponses, err := rh.Storage.FetchRentalByQuery(r.URL.Query())
	if err != nil {
		if err == db.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Convert to rental responses
	rentalResponses := []*models.RentalResponse{}
	for _, dbResponse := range dbResponses {
		rentalResponses = append(rentalResponses, convertDBResponseToRentalResponse(&dbResponse))
	}

	// Marshal the rentals to JSON
	response, err := json.Marshal(rentalResponses)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func convertDBResponseToRentalResponse(r *dbmodels.UserAndRental) *models.RentalResponse {
	return &models.RentalResponse{
		ID:              r.Rental.ID,
		Name:            r.Rental.Name,
		Description:     r.Rental.Description,
		Type:            r.Rental.Type,
		Make:            r.Rental.VehicleMake,
		Model:           r.Rental.VehicleModel,
		Year:            r.Rental.VehicleYear,
		Length:          r.Rental.VehicleLength,
		Sleeps:          r.Rental.Sleeps,
		PrimaryImageURL: r.Rental.PrimaryImageURL,
		Price: models.Price{
			Day: r.Rental.PricePerDay,
		},
		Location: models.Location{
			City:    r.Rental.HomeCity,
			State:   r.Rental.HomeState,
			Zip:     r.Rental.HomeZip,
			Country: r.Rental.HomeCountry,
			Lat:     r.Rental.Lat,
			Lng:     r.Rental.Lng,
		},
		User: models.User{
			ID:        r.User.ID,
			FirstName: r.User.FirstName,
			LastName:  r.User.LastName,
		},
	}
}
