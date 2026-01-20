package surplus

import (
	"database/sql"
	"encoding/json"
	"foodlink_backend/database"
	"foodlink_backend/errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() *Repository {
	return &Repository{db: database.GetDB()}
}

func (r *Repository) GetAllByUserID(userID uuid.UUID) ([]*RestaurantSurplusItem, error) {
	if r.db == nil {
		return nil, errors.ErrDatabase
	}
	query := `SELECT id, user_id, title, description, quantity, unit, category, storage_type, pickup_window, tags, image, assigned_to, recipient_name, status, created_at, updated_at FROM restaurant_surplus_items WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, errors.WrapError(err, errors.ErrDatabase)
	}
	defer rows.Close()
	var items []*RestaurantSurplusItem
	for rows.Next() {
		item := &RestaurantSurplusItem{}
		var pickupWindowJSON []byte
		if err := rows.Scan(&item.ID, &item.UserID, &item.Title, &item.Description, &item.Quantity, &item.Unit, &item.Category, &item.StorageType, &pickupWindowJSON, pq.Array(&item.Tags), &item.Image, &item.AssignedTo, &item.RecipientName, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, errors.WrapError(err, errors.ErrDatabase)
		}
		if len(pickupWindowJSON) > 0 {
			json.Unmarshal(pickupWindowJSON, &item.PickupWindow)
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) GetByID(id uuid.UUID) (*RestaurantSurplusItem, error) {
	if r.db == nil {
		return nil, errors.ErrDatabase
	}
	item := &RestaurantSurplusItem{}
	var pickupWindowJSON []byte
	query := `SELECT id, user_id, title, description, quantity, unit, category, storage_type, pickup_window, tags, image, assigned_to, recipient_name, status, created_at, updated_at FROM restaurant_surplus_items WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&item.ID, &item.UserID, &item.Title, &item.Description, &item.Quantity, &item.Unit, &item.Category, &item.StorageType, &pickupWindowJSON, pq.Array(&item.Tags), &item.Image, &item.AssignedTo, &item.RecipientName, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound
		}
		return nil, errors.WrapError(err, errors.ErrDatabase)
	}
	if len(pickupWindowJSON) > 0 {
		json.Unmarshal(pickupWindowJSON, &item.PickupWindow)
	}
	return item, nil
}

func (r *Repository) Create(item *RestaurantSurplusItem) error {
	if r.db == nil {
		return errors.ErrDatabase
	}
	pickupWindowJSON, _ := json.Marshal(item.PickupWindow)
	query := `INSERT INTO restaurant_surplus_items (id, user_id, title, description, quantity, unit, category, storage_type, pickup_window, tags, image, assigned_to, recipient_name, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id, user_id, title, description, quantity, unit, category, storage_type, pickup_window, tags, image, assigned_to, recipient_name, status, created_at, updated_at`
	now := time.Now()
	var pickupWindowJSONOut []byte
	err := r.db.QueryRow(query, item.ID, item.UserID, item.Title, item.Description, item.Quantity, item.Unit, item.Category, item.StorageType, pickupWindowJSON, pq.Array(item.Tags), item.Image, item.AssignedTo, item.RecipientName, item.Status, now, now).Scan(&item.ID, &item.UserID, &item.Title, &item.Description, &item.Quantity, &item.Unit, &item.Category, &item.StorageType, &pickupWindowJSONOut, pq.Array(&item.Tags), &item.Image, &item.AssignedTo, &item.RecipientName, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return errors.WrapError(err, errors.ErrDatabase)
	}
	if len(pickupWindowJSONOut) > 0 {
		json.Unmarshal(pickupWindowJSONOut, &item.PickupWindow)
	}
	return nil
}

func (r *Repository) Update(item *RestaurantSurplusItem) error {
	if r.db == nil {
		return errors.ErrDatabase
	}
	pickupWindowJSON, _ := json.Marshal(item.PickupWindow)
	query := `UPDATE restaurant_surplus_items SET title=$1, description=$2, quantity=$3, unit=$4, category=$5, storage_type=$6, pickup_window=$7, tags=$8, image=$9, assigned_to=$10, recipient_name=$11, status=$12, updated_at=$13 WHERE id=$14 RETURNING id, user_id, title, description, quantity, unit, category, storage_type, pickup_window, tags, image, assigned_to, recipient_name, status, created_at, updated_at`
	var pickupWindowJSONOut []byte
	err := r.db.QueryRow(query, item.Title, item.Description, item.Quantity, item.Unit, item.Category, item.StorageType, pickupWindowJSON, pq.Array(item.Tags), item.Image, item.AssignedTo, item.RecipientName, item.Status, time.Now(), item.ID).Scan(&item.ID, &item.UserID, &item.Title, &item.Description, &item.Quantity, &item.Unit, &item.Category, &item.StorageType, &pickupWindowJSONOut, pq.Array(&item.Tags), &item.Image, &item.AssignedTo, &item.RecipientName, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.ErrNotFound
		}
		return errors.WrapError(err, errors.ErrDatabase)
	}
	if len(pickupWindowJSONOut) > 0 {
		json.Unmarshal(pickupWindowJSONOut, &item.PickupWindow)
	}
	return nil
}
