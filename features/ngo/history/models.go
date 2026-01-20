package history

import (
	"time"

	"github.com/google/uuid"
)

type NGODonationHistory struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	NGOUserID      uuid.UUID  `json:"ngo_user_id" db:"ngo_user_id"`
	OfferID        *uuid.UUID `json:"offer_id,omitempty" db:"offer_id"`
	DonorName      string     `json:"donor_name" db:"donor_name"`
	DonorType      string     `json:"donor_type" db:"donor_type"`
	ItemsSummary   string     `json:"items_summary" db:"items_summary"`
	WeightKg       float64    `json:"weight_kg" db:"weight_kg"`
	MealsProvided  int        `json:"meals_provided" db:"meals_provided"`
	CO2PreventedKg float64    `json:"co2_prevented_kg" db:"co2_prevented_kg"`
	Beneficiaries  int        `json:"beneficiaries" db:"beneficiaries"`
	PickupTime     time.Time  `json:"pickup_time" db:"pickup_time"`
	DeliveredAt    *time.Time `json:"delivered_at,omitempty" db:"delivered_at"`
	Status         string     `json:"status" db:"status"`
	Tags           []string   `json:"tags,omitempty" db:"tags"`
	Photo          string     `json:"photo,omitempty" db:"photo"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
}
