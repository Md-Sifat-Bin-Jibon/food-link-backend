package surplus

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

type RestaurantSurplusItem struct {
	ID           uuid.UUID `json:"id" db:"id"`
	UserID       uuid.UUID `json:"user_id" db:"user_id"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	Quantity     float64   `json:"quantity" db:"quantity"`
	Unit         string    `json:"unit" db:"unit"`
	Category     string    `json:"category" db:"category"`
	StorageType  string    `json:"storage_type" db:"storage_type"`
	PickupWindow JSONB     `json:"pickup_window" db:"pickup_window"`
	Tags         []string  `json:"tags,omitempty" db:"tags"`
	Image        string    `json:"image,omitempty" db:"image"`
	AssignedTo   string    `json:"assigned_to,omitempty" db:"assigned_to"`
	RecipientName string   `json:"recipient_name,omitempty" db:"recipient_name"`
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type CreateRestaurantSurplusItemRequest struct {
	Title        string                 `json:"title" validate:"required,min=1,max=255"`
	Description  string                 `json:"description" validate:"required,min=1"`
	Quantity     float64                `json:"quantity" validate:"required,gt=0"`
	Unit         string                 `json:"unit" validate:"required,min=1,max=50"`
	Category     string                 `json:"category" validate:"required,min=1,max=100"`
	StorageType  string                 `json:"storage_type" validate:"required,oneof=fresh chilled frozen"`
	PickupWindow map[string]interface{} `json:"pickup_window" validate:"required"`
	Tags         []string               `json:"tags,omitempty"`
	Image        string                 `json:"image,omitempty"`
}

type UpdateRestaurantSurplusItemRequest struct {
	Title        string                 `json:"title,omitempty" validate:"omitempty,min=1,max=255"`
	Description  string                 `json:"description,omitempty" validate:"omitempty,min=1"`
	Quantity     *float64               `json:"quantity,omitempty" validate:"omitempty,gt=0"`
	Unit         string                 `json:"unit,omitempty" validate:"omitempty,min=1,max=50"`
	Category     string                 `json:"category,omitempty" validate:"omitempty,min=1,max=100"`
	StorageType  string                 `json:"storage_type,omitempty" validate:"omitempty,oneof=fresh chilled frozen"`
	PickupWindow map[string]interface{} `json:"pickup_window,omitempty"`
	Tags         []string               `json:"tags,omitempty"`
	Image        string                 `json:"image,omitempty"`
	Status       string                 `json:"status,omitempty"`
}

type AssignSurplusItemRequest struct {
	AssignedTo   string `json:"assigned_to" validate:"required,oneof=ngo kitchen"`
	RecipientName string `json:"recipient_name" validate:"required,min=1"`
}
