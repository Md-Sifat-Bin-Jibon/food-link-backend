package pickups

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

type NGOPickupSchedule struct {
	ID             uuid.UUID `json:"id" db:"id"`
	OfferID        uuid.UUID `json:"offer_id" db:"offer_id"`
	RouteID        *uuid.UUID `json:"route_id,omitempty" db:"route_id"`
	ScheduledFor  time.Time `json:"scheduled_for" db:"scheduled_for"`
	ETAMinutes     *int      `json:"eta_minutes,omitempty" db:"eta_minutes"`
	VolunteerName  string    `json:"volunteer_name" db:"volunteer_name"`
	VolunteerContact string  `json:"volunteer_contact" db:"volunteer_contact"`
	VehicleType    string    `json:"vehicle_type,omitempty" db:"vehicle_type"`
	Status         string    `json:"status" db:"status"`
	Checkpoints    JSONB     `json:"checkpoints,omitempty" db:"checkpoints"`
	Reminders      JSONB     `json:"reminders,omitempty" db:"reminders"`
	Notes          string    `json:"notes,omitempty" db:"notes"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type CreateNGOPickupScheduleRequest struct {
	OfferID        uuid.UUID              `json:"offer_id" validate:"required"`
	ScheduledFor   time.Time              `json:"scheduled_for" validate:"required"`
	ETAMinutes     *int                   `json:"eta_minutes,omitempty"`
	VolunteerName  string                 `json:"volunteer_name" validate:"required,min=1"`
	VolunteerContact string               `json:"volunteer_contact" validate:"required,min=1"`
	VehicleType    string                 `json:"vehicle_type,omitempty" validate:"omitempty,oneof=van bike car on-foot"`
	Checkpoints    map[string]interface{} `json:"checkpoints,omitempty"`
	Reminders      map[string]interface{} `json:"reminders,omitempty"`
	Notes          string                 `json:"notes,omitempty"`
}

type UpdateNGOPickupScheduleRequest struct {
	ScheduledFor   *time.Time             `json:"scheduled_for,omitempty"`
	ETAMinutes     *int                   `json:"eta_minutes,omitempty"`
	VolunteerName  string                 `json:"volunteer_name,omitempty"`
	VolunteerContact string               `json:"volunteer_contact,omitempty"`
	VehicleType    string                 `json:"vehicle_type,omitempty"`
	Checkpoints    map[string]interface{} `json:"checkpoints,omitempty"`
	Reminders      map[string]interface{} `json:"reminders,omitempty"`
	Notes          string                 `json:"notes,omitempty"`
}

type UpdatePickupStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=scheduled en-route picked-up delivered failed"`
}
