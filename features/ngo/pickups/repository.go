package pickups

import (
	"database/sql"
	"encoding/json"
	"foodlink_backend/database"
	"foodlink_backend/errors"
	"time"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() *Repository {
	return &Repository{db: database.GetDB()}
}

func (r *Repository) GetAllByOfferID(offerID uuid.UUID) ([]*NGOPickupSchedule, error) {
	if r.db == nil {
		return nil, errors.ErrDatabase
	}
	query := `SELECT id, offer_id, route_id, scheduled_for, eta_minutes, volunteer_name, volunteer_contact, vehicle_type, status, checkpoints, reminders, notes, created_at, updated_at FROM ngo_pickup_schedules WHERE offer_id = $1 ORDER BY scheduled_for ASC`
	rows, err := r.db.Query(query, offerID)
	if err != nil {
		return nil, errors.WrapError(err, errors.ErrDatabase)
	}
	defer rows.Close()
	var schedules []*NGOPickupSchedule
	for rows.Next() {
		schedule := &NGOPickupSchedule{}
		var checkpointsJSON, remindersJSON []byte
		if err := rows.Scan(&schedule.ID, &schedule.OfferID, &schedule.RouteID, &schedule.ScheduledFor, &schedule.ETAMinutes, &schedule.VolunteerName, &schedule.VolunteerContact, &schedule.VehicleType, &schedule.Status, &checkpointsJSON, &remindersJSON, &schedule.Notes, &schedule.CreatedAt, &schedule.UpdatedAt); err != nil {
			return nil, errors.WrapError(err, errors.ErrDatabase)
		}
		if len(checkpointsJSON) > 0 {
			json.Unmarshal(checkpointsJSON, &schedule.Checkpoints)
		}
		if len(remindersJSON) > 0 {
			json.Unmarshal(remindersJSON, &schedule.Reminders)
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (r *Repository) GetByID(id uuid.UUID) (*NGOPickupSchedule, error) {
	if r.db == nil {
		return nil, errors.ErrDatabase
	}
	schedule := &NGOPickupSchedule{}
	var checkpointsJSON, remindersJSON []byte
	query := `SELECT id, offer_id, route_id, scheduled_for, eta_minutes, volunteer_name, volunteer_contact, vehicle_type, status, checkpoints, reminders, notes, created_at, updated_at FROM ngo_pickup_schedules WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&schedule.ID, &schedule.OfferID, &schedule.RouteID, &schedule.ScheduledFor, &schedule.ETAMinutes, &schedule.VolunteerName, &schedule.VolunteerContact, &schedule.VehicleType, &schedule.Status, &checkpointsJSON, &remindersJSON, &schedule.Notes, &schedule.CreatedAt, &schedule.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound
		}
		return nil, errors.WrapError(err, errors.ErrDatabase)
	}
	if len(checkpointsJSON) > 0 {
		json.Unmarshal(checkpointsJSON, &schedule.Checkpoints)
	}
	if len(remindersJSON) > 0 {
		json.Unmarshal(remindersJSON, &schedule.Reminders)
	}
	return schedule, nil
}

func (r *Repository) Create(schedule *NGOPickupSchedule) error {
	if r.db == nil {
		return errors.ErrDatabase
	}
	checkpointsJSON, _ := json.Marshal(schedule.Checkpoints)
	remindersJSON, _ := json.Marshal(schedule.Reminders)
	query := `INSERT INTO ngo_pickup_schedules (id, offer_id, route_id, scheduled_for, eta_minutes, volunteer_name, volunteer_contact, vehicle_type, status, checkpoints, reminders, notes, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id, offer_id, route_id, scheduled_for, eta_minutes, volunteer_name, volunteer_contact, vehicle_type, status, checkpoints, reminders, notes, created_at, updated_at`
	now := time.Now()
	var checkpointsJSONOut, remindersJSONOut []byte
	err := r.db.QueryRow(query, schedule.ID, schedule.OfferID, schedule.RouteID, schedule.ScheduledFor, schedule.ETAMinutes, schedule.VolunteerName, schedule.VolunteerContact, schedule.VehicleType, schedule.Status, checkpointsJSON, remindersJSON, schedule.Notes, now, now).Scan(&schedule.ID, &schedule.OfferID, &schedule.RouteID, &schedule.ScheduledFor, &schedule.ETAMinutes, &schedule.VolunteerName, &schedule.VolunteerContact, &schedule.VehicleType, &schedule.Status, &checkpointsJSONOut, &remindersJSONOut, &schedule.Notes, &schedule.CreatedAt, &schedule.UpdatedAt)
	if err != nil {
		return errors.WrapError(err, errors.ErrDatabase)
	}
	if len(checkpointsJSONOut) > 0 {
		json.Unmarshal(checkpointsJSONOut, &schedule.Checkpoints)
	}
	if len(remindersJSONOut) > 0 {
		json.Unmarshal(remindersJSONOut, &schedule.Reminders)
	}
	return nil
}

func (r *Repository) Update(schedule *NGOPickupSchedule) error {
	if r.db == nil {
		return errors.ErrDatabase
	}
	checkpointsJSON, _ := json.Marshal(schedule.Checkpoints)
	remindersJSON, _ := json.Marshal(schedule.Reminders)
	query := `UPDATE ngo_pickup_schedules SET scheduled_for=$1, eta_minutes=$2, volunteer_name=$3, volunteer_contact=$4, vehicle_type=$5, status=$6, checkpoints=$7, reminders=$8, notes=$9, updated_at=$10 WHERE id=$11 RETURNING id, offer_id, route_id, scheduled_for, eta_minutes, volunteer_name, volunteer_contact, vehicle_type, status, checkpoints, reminders, notes, created_at, updated_at`
	var checkpointsJSONOut, remindersJSONOut []byte
	err := r.db.QueryRow(query, schedule.ScheduledFor, schedule.ETAMinutes, schedule.VolunteerName, schedule.VolunteerContact, schedule.VehicleType, schedule.Status, checkpointsJSON, remindersJSON, schedule.Notes, time.Now(), schedule.ID).Scan(&schedule.ID, &schedule.OfferID, &schedule.RouteID, &schedule.ScheduledFor, &schedule.ETAMinutes, &schedule.VolunteerName, &schedule.VolunteerContact, &schedule.VehicleType, &schedule.Status, &checkpointsJSONOut, &remindersJSONOut, &schedule.Notes, &schedule.CreatedAt, &schedule.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.ErrNotFound
		}
		return errors.WrapError(err, errors.ErrDatabase)
	}
	if len(checkpointsJSONOut) > 0 {
		json.Unmarshal(checkpointsJSONOut, &schedule.Checkpoints)
	}
	if len(remindersJSONOut) > 0 {
		json.Unmarshal(remindersJSONOut, &schedule.Reminders)
	}
	return nil
}
