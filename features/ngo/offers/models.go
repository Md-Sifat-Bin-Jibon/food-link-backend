package offers

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

type NGODonationOffer struct {
	ID            uuid.UUID `json:"id" db:"id"`
	NGOUserID     uuid.UUID `json:"ngo_user_id" db:"ngo_user_id"`
	DonorName     string    `json:"donor_name" db:"donor_name"`
	DonorType     string    `json:"donor_type" db:"donor_type"`
	PartnerID     *uuid.UUID `json:"partner_id,omitempty" db:"partner_id"`
	DistanceKm    float64   `json:"distance_km" db:"distance_km"`
	LocationLabel string    `json:"location_label" db:"location_label"`
	GeoPoint      JSONB     `json:"geo_point,omitempty" db:"geo_point"`
	OfferTitle    string    `json:"offer_title" db:"offer_title"`
	Items         JSONB     `json:"items" db:"items"`
	WeightKg      float64   `json:"weight_kg" db:"weight_kg"`
	MealsEstimated int      `json:"meals_estimated" db:"meals_estimated"`
	FreshnessScore int      `json:"freshness_score" db:"freshness_score"`
	PickupWindow  JSONB     `json:"pickup_window" db:"pickup_window"`
	ExpiresAt     time.Time `json:"expires_at" db:"expires_at"`
	UrgencyLevel  string    `json:"urgency_level" db:"urgency_level"`
	DietaryNotes  string    `json:"dietary_notes,omitempty" db:"dietary_notes"`
	SafetyFlags   []string  `json:"safety_flags,omitempty" db:"safety_flags"`
	Contact       JSONB     `json:"contact" db:"contact"`
	Images        []string  `json:"images,omitempty" db:"images"`
	Status        string    `json:"status" db:"status"`
	MatchReason   string    `json:"match_reason,omitempty" db:"match_reason"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
