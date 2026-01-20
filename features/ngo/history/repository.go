package history

import (
	"database/sql"
	"foodlink_backend/database"
	"foodlink_backend/errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() *Repository {
	return &Repository{db: database.GetDB()}
}

func (r *Repository) GetAllByNGOUserID(ngoUserID uuid.UUID) ([]*NGODonationHistory, error) {
	if r.db == nil {
		return nil, errors.ErrDatabase
	}
	query := `SELECT id, ngo_user_id, offer_id, donor_name, donor_type, items_summary, weight_kg, meals_provided, co2_prevented_kg, beneficiaries, pickup_time, delivered_at, status, tags, photo, created_at FROM ngo_donation_history WHERE ngo_user_id = $1 ORDER BY pickup_time DESC, created_at DESC`
	rows, err := r.db.Query(query, ngoUserID)
	if err != nil {
		return nil, errors.WrapError(err, errors.ErrDatabase)
	}
	defer rows.Close()
	var histories []*NGODonationHistory
	for rows.Next() {
		history := &NGODonationHistory{}
		if err := rows.Scan(&history.ID, &history.NGOUserID, &history.OfferID, &history.DonorName, &history.DonorType, &history.ItemsSummary, &history.WeightKg, &history.MealsProvided, &history.CO2PreventedKg, &history.Beneficiaries, &history.PickupTime, &history.DeliveredAt, &history.Status, pq.Array(&history.Tags), &history.Photo, &history.CreatedAt); err != nil {
			return nil, errors.WrapError(err, errors.ErrDatabase)
		}
		histories = append(histories, history)
	}
	return histories, nil
}

func (r *Repository) GetByID(id uuid.UUID) (*NGODonationHistory, error) {
	if r.db == nil {
		return nil, errors.ErrDatabase
	}
	history := &NGODonationHistory{}
	query := `SELECT id, ngo_user_id, offer_id, donor_name, donor_type, items_summary, weight_kg, meals_provided, co2_prevented_kg, beneficiaries, pickup_time, delivered_at, status, tags, photo, created_at FROM ngo_donation_history WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&history.ID, &history.NGOUserID, &history.OfferID, &history.DonorName, &history.DonorType, &history.ItemsSummary, &history.WeightKg, &history.MealsProvided, &history.CO2PreventedKg, &history.Beneficiaries, &history.PickupTime, &history.DeliveredAt, &history.Status, pq.Array(&history.Tags), &history.Photo, &history.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound
		}
		return nil, errors.WrapError(err, errors.ErrDatabase)
	}
	return history, nil
}
