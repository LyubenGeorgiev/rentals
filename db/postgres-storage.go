package db

import (
	"database/sql"
	"net/url"
	"strconv"
	"strings"

	dbmodels "github.com/LyubenGeorgiev/rentals/db/models"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

// PostgresStorage represents the PostgreSQL implementation of Storage
// wrapped around with goqu Database
type PostgresStorage struct {
	DB *goqu.Database
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		DB: goqu.New("postgres", db),
	}
}

// FetchRentalByID fetches a rental by ID
func (ps *PostgresStorage) FetchRentalByID(id int) (*dbmodels.UserAndRental, error) {
	joinQuery := ps.DB.From("users").Join(
		ps.DB.From("rentals").Where(goqu.C("id").Eq(id)).As("rentals"),
		goqu.On(goqu.Ex{
			"users.id": goqu.I("rentals.user_id"),
		}),
	)

	var userAndRental dbmodels.UserAndRental
	found, err := joinQuery.ScanStruct(&userAndRental)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ErrNotFound
	}

	return &userAndRental, nil
}

// FetchRentalByQuery fetches multiple rentals based on query parameters
func (ps *PostgresStorage) FetchRentalByQuery(queryParams url.Values) ([]dbmodels.UserAndRental, error) {
	filters := []exp.Expression{}

	if minPrice, err := strconv.Atoi(queryParams.Get("price_min")); err == nil {
		filters = append(filters, goqu.C("price_per_day").Gte(minPrice))
	}
	if maxPrice, err := strconv.Atoi(queryParams.Get("price_max")); err == nil {
		filters = append(filters, goqu.C("price_per_day").Lte(maxPrice))
	}
	if cords := strings.Split(queryParams.Get("near"), ","); len(cords) == 2 {
		lat, err := strconv.ParseFloat(cords[0], 32)
		if err == nil {
			lng, err := strconv.ParseFloat(cords[1], 32)
			if err == nil {
				distance := goqu.L("ST_Distance(ST_Point(?, ?)::geography, ST_Point(lat, lng)::geography)", float32(lat), float32(lng))
				filters = append(filters, distance.Lte(100*1609.34))
			}
		}
	}

	ids := []int{}
	for _, idStr := range strings.Split(queryParams.Get("ids"), ",") {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	if len(ids) > 0 {
		filters = append(filters, goqu.C("id").In(ids))
	}

	queryBuilder := ps.DB.From("users").Join(
		ps.DB.From("rentals").Where(filters...).As("rentals"),
		goqu.On(goqu.Ex{
			"users.id": goqu.I("rentals.user_id"),
		}),
	)

	if column := queryParams.Get("sort"); column != "" {
		queryBuilder = queryBuilder.Order(goqu.I(column).Asc())
	}
	if limit, err := strconv.Atoi(queryParams.Get("limit")); err == nil && limit >= 0 {
		queryBuilder = queryBuilder.Limit(uint(limit))
	}
	if offset, err := strconv.Atoi(queryParams.Get("offset")); err == nil && offset >= 0 {
		queryBuilder = queryBuilder.Offset(uint(offset))
	}

	var userAndRentals []dbmodels.UserAndRental
	if err := queryBuilder.ScanStructs(&userAndRentals); err != nil {
		return nil, err
	}

	return userAndRentals, nil
}
