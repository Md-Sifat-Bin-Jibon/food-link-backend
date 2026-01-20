package partners

import (
	"time"

	"github.com/google/uuid"
)

type NGOPartnerProfile struct {
	ID                 uuid.UUID  `json:"id" db:"id"`
	NGOUserID          uuid.UUID  `json:"ngo_user_id" db:"ngo_user_id"`
	Name               string     `json:"name" db:"name"`
	Type               string     `json:"type" db:"type"`
	Location           string     `json:"location" db:"location"`
	DistanceKm         *float64   `json:"distance_km,omitempty" db:"distance_km"`
	ContactName        string     `json:"contact_name" db:"contact_name"`
	ContactPhone       string     `json:"contact_phone" db:"contact_phone"`
	ContactEmail       string     `json:"contact_email,omitempty" db:"contact_email"`
	OperatingHours     string     `json:"operating_hours,omitempty" db:"operating_hours"`
	AcceptanceRate     float64    `json:"acceptance_rate" db:"acceptance_rate"`
	LastDonationAt     *time.Time `json:"last_donation_at,omitempty" db:"last_donation_at"`
	AvgDonationKg      float64    `json:"avg_donation_kg" db:"avg_donation_kg"`
	StorageCapabilities []string   `json:"storage_capabilities,omitempty" db:"storage_capabilities"`
	Notes               string     `json:"notes,omitempty" db:"notes"`
	Avatar              string     `json:"avatar,omitempty" db:"avatar"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateNGOPartnerRequest struct {
	Name               string   `json:"name" validate:"required,min=1,max=255"`
	Type               string   `json:"type" validate:"required,oneof=community-kitchen building restaurant ngo"`
	Location           string   `json:"location" validate:"required,min=1"`
	DistanceKm         *float64 `json:"distance_km,omitempty"`
	ContactName        string   `json:"contact_name" validate:"required,min=1,max=255"`
	ContactPhone       string   `json:"contact_phone" validate:"required,min=1,max=50"`
	ContactEmail       string   `json:"contact_email,omitempty" validate:"omitempty,email"`
	OperatingHours     string   `json:"operating_hours,omitempty"`
	StorageCapabilities []string `json:"storage_capabilities,omitempty"`
	Notes               string   `json:"notes,omitempty"`
	Avatar              string   `json:"avatar,omitempty"`
}

type UpdateNGOPartnerRequest struct {
	Name               string   `json:"name,omitempty" validate:"omitempty,min=1,max=255"`
	Location           string   `json:"location,omitempty" validate:"omitempty,min=1"`
	DistanceKm         *float64 `json:"distance_km,omitempty"`
	ContactName        string   `json:"contact_name,omitempty" validate:"omitempty,min=1,max=255"`
	ContactPhone       string   `json:"contact_phone,omitempty" validate:"omitempty,min=1,max=50"`
	ContactEmail       string   `json:"contact_email,omitempty" validate:"omitempty,email"`
	OperatingHours     string   `json:"operating_hours,omitempty"`
	StorageCapabilities []string `json:"storage_capabilities,omitempty"`
	Notes               string   `json:"notes,omitempty"`
	Avatar              string   `json:"avatar,omitempty"`
}
