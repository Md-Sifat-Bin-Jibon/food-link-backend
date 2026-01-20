package offers

import (
	"database/sql"
	"encoding/json"
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

func (r *Repository) GetAllByNGOUserID(ngoUserID uuid.UUID, status string) ([]*NGODonationOffer, error) {
	if r.db == nil {
		return nil, errors.ErrDatabase
	}
	var query string
	var rows *sql.Rows
	var err error
	if status != "" {
		query = `SELECT id, ngo_user_id, donor_name, donor_type, partner_id, distance_km, location_label, geo_point, offer_title, items, weight_kg, meals_estimated, freshness_score, pickup_window, expires_at, urgency_level, dietary_notes, safety_flags, contact, images, status, match_reason, created_at, updated_at FROM ngo_donation_offers WHERE ngo_user_id = $1 AND status = $2 ORDER BY created_at DESC`
		rows, err = r.db.Query(query, ngoUserID, status)
	} else {
		query = `SELECT id, ngo_user_id, donor_name, donor_type, partner_id, distance_km, location_label, geo_point, offer_title, items, weight_kg, meals_estimated, freshness_score, pickup_window, expires_at, urgency_level, dietary_notes, safety_flags, contact, images, status, match_reason, created_at, updated_at FROM ngo_donation_offers WHERE ngo_user_id = $1 ORDER BY created_at DESC`
		rows, err = r.db.Query(query, ngoUserID)
	}
	if err != nil {
		return nil, errors.WrapError(err, errors.ErrDatabase)
	}
	defer rows.Close()
	var offers []*NGODonationOffer
	for rows.Next() {
		offer := &NGODonationOffer{}
		var geoPointJSON, itemsJSON, pickupWindowJSON, contactJSON []byte
		if err := rows.Scan(&offer.ID, &offer.NGOUserID, &offer.DonorName, &offer.DonorType, &offer.PartnerID, &offer.DistanceKm, &offer.LocationLabel, &geoPointJSON, &offer.OfferTitle, &itemsJSON, &offer.WeightKg, &offer.MealsEstimated, &offer.FreshnessScore, &pickupWindowJSON, &offer.ExpiresAt, &offer.UrgencyLevel, &offer.DietaryNotes, pq.Array(&offer.SafetyFlags), &contactJSON, pq.Array(&offer.Images), &offer.Status, &offer.MatchReason, &offer.CreatedAt, &offer.UpdatedAt); err != nil {
			return nil, errors.WrapError(err, errors.ErrDatabase)
		}
		if len(geoPointJSON) > 0 {
			json.Unmarshal(geoPointJSON, &offer.GeoPoint)
		}
		if len(itemsJSON) > 0 {
			json.Unmarshal(itemsJSON, &offer.Items)
		}
		if len(pickupWindowJSON) > 0 {
			json.Unmarshal(pickupWindowJSON, &offer.PickupWindow)
		}
		if len(contactJSON) > 0 {
			json.Unmarshal(contactJSON, &offer.Contact)
		}
		offers = append(offers, offer)
	}
	return offers, nil
}

func (r *Repository) GetByID(id uuid.UUID) (*NGODonationOffer, error) {
	if r.db == nil {
		return nil, errors.ErrDatabase
	}
	offer := &NGODonationOffer{}
	var geoPointJSON, itemsJSON, pickupWindowJSON, contactJSON []byte
	query := `SELECT id, ngo_user_id, donor_name, donor_type, partner_id, distance_km, location_label, geo_point, offer_title, items, weight_kg, meals_estimated, freshness_score, pickup_window, expires_at, urgency_level, dietary_notes, safety_flags, contact, images, status, match_reason, created_at, updated_at FROM ngo_donation_offers WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&offer.ID, &offer.NGOUserID, &offer.DonorName, &offer.DonorType, &offer.PartnerID, &offer.DistanceKm, &offer.LocationLabel, &geoPointJSON, &offer.OfferTitle, &itemsJSON, &offer.WeightKg, &offer.MealsEstimated, &offer.FreshnessScore, &pickupWindowJSON, &offer.ExpiresAt, &offer.UrgencyLevel, &offer.DietaryNotes, pq.Array(&offer.SafetyFlags), &contactJSON, pq.Array(&offer.Images), &offer.Status, &offer.MatchReason, &offer.CreatedAt, &offer.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound
		}
		return nil, errors.WrapError(err, errors.ErrDatabase)
	}
	if len(geoPointJSON) > 0 {
		json.Unmarshal(geoPointJSON, &offer.GeoPoint)
	}
	if len(itemsJSON) > 0 {
		json.Unmarshal(itemsJSON, &offer.Items)
	}
	if len(pickupWindowJSON) > 0 {
		json.Unmarshal(pickupWindowJSON, &offer.PickupWindow)
	}
	if len(contactJSON) > 0 {
		json.Unmarshal(contactJSON, &offer.Contact)
	}
	return offer, nil
}

func (r *Repository) UpdateStatus(id uuid.UUID, status string) error {
	if r.db == nil {
		return errors.ErrDatabase
	}
	query := `UPDATE ngo_donation_offers SET status=$1, updated_at=CURRENT_TIMESTAMP WHERE id=$2`
	result, err := r.db.Exec(query, status, id)
	if err != nil {
		return errors.WrapError(err, errors.ErrDatabase)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.ErrNotFound
	}
	return nil
}
