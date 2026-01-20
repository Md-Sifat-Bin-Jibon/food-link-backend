package partners

import (
	"database/sql"
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

func (r *Repository) GetAllByNGOUserID(ngoUserID uuid.UUID) ([]*NGOPartnerProfile, error) {
	if r.db == nil {
		return nil, errors.ErrDatabase
	}
	query := `SELECT id, ngo_user_id, name, type, location, distance_km, contact_name, contact_phone, contact_email, operating_hours, acceptance_rate, last_donation_at, avg_donation_kg, storage_capabilities, notes, avatar, created_at, updated_at FROM ngo_partner_profiles WHERE ngo_user_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.Query(query, ngoUserID)
	if err != nil {
		return nil, errors.WrapError(err, errors.ErrDatabase)
	}
	defer rows.Close()
	var partners []*NGOPartnerProfile
	for rows.Next() {
		partner := &NGOPartnerProfile{}
		if err := rows.Scan(&partner.ID, &partner.NGOUserID, &partner.Name, &partner.Type, &partner.Location, &partner.DistanceKm, &partner.ContactName, &partner.ContactPhone, &partner.ContactEmail, &partner.OperatingHours, &partner.AcceptanceRate, &partner.LastDonationAt, &partner.AvgDonationKg, pq.Array(&partner.StorageCapabilities), &partner.Notes, &partner.Avatar, &partner.CreatedAt, &partner.UpdatedAt); err != nil {
			return nil, errors.WrapError(err, errors.ErrDatabase)
		}
		partners = append(partners, partner)
	}
	return partners, nil
}

func (r *Repository) GetByID(id uuid.UUID) (*NGOPartnerProfile, error) {
	if r.db == nil {
		return nil, errors.ErrDatabase
	}
	partner := &NGOPartnerProfile{}
	query := `SELECT id, ngo_user_id, name, type, location, distance_km, contact_name, contact_phone, contact_email, operating_hours, acceptance_rate, last_donation_at, avg_donation_kg, storage_capabilities, notes, avatar, created_at, updated_at FROM ngo_partner_profiles WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&partner.ID, &partner.NGOUserID, &partner.Name, &partner.Type, &partner.Location, &partner.DistanceKm, &partner.ContactName, &partner.ContactPhone, &partner.ContactEmail, &partner.OperatingHours, &partner.AcceptanceRate, &partner.LastDonationAt, &partner.AvgDonationKg, pq.Array(&partner.StorageCapabilities), &partner.Notes, &partner.Avatar, &partner.CreatedAt, &partner.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound
		}
		return nil, errors.WrapError(err, errors.ErrDatabase)
	}
	return partner, nil
}

func (r *Repository) Create(partner *NGOPartnerProfile) error {
	if r.db == nil {
		return errors.ErrDatabase
	}
	query := `INSERT INTO ngo_partner_profiles (id, ngo_user_id, name, type, location, distance_km, contact_name, contact_phone, contact_email, operating_hours, acceptance_rate, last_donation_at, avg_donation_kg, storage_capabilities, notes, avatar, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) RETURNING id, ngo_user_id, name, type, location, distance_km, contact_name, contact_phone, contact_email, operating_hours, acceptance_rate, last_donation_at, avg_donation_kg, storage_capabilities, notes, avatar, created_at, updated_at`
	now := time.Now()
	return r.db.QueryRow(query, partner.ID, partner.NGOUserID, partner.Name, partner.Type, partner.Location, partner.DistanceKm, partner.ContactName, partner.ContactPhone, partner.ContactEmail, partner.OperatingHours, partner.AcceptanceRate, partner.LastDonationAt, partner.AvgDonationKg, pq.Array(partner.StorageCapabilities), partner.Notes, partner.Avatar, now, now).Scan(&partner.ID, &partner.NGOUserID, &partner.Name, &partner.Type, &partner.Location, &partner.DistanceKm, &partner.ContactName, &partner.ContactPhone, &partner.ContactEmail, &partner.OperatingHours, &partner.AcceptanceRate, &partner.LastDonationAt, &partner.AvgDonationKg, pq.Array(&partner.StorageCapabilities), &partner.Notes, &partner.Avatar, &partner.CreatedAt, &partner.UpdatedAt)
}

func (r *Repository) Update(partner *NGOPartnerProfile) error {
	if r.db == nil {
		return errors.ErrDatabase
	}
	query := `UPDATE ngo_partner_profiles SET name=$1, location=$2, distance_km=$3, contact_name=$4, contact_phone=$5, contact_email=$6, operating_hours=$7, storage_capabilities=$8, notes=$9, avatar=$10, updated_at=$11 WHERE id=$12 RETURNING id, ngo_user_id, name, type, location, distance_km, contact_name, contact_phone, contact_email, operating_hours, acceptance_rate, last_donation_at, avg_donation_kg, storage_capabilities, notes, avatar, created_at, updated_at`
	return r.db.QueryRow(query, partner.Name, partner.Location, partner.DistanceKm, partner.ContactName, partner.ContactPhone, partner.ContactEmail, partner.OperatingHours, pq.Array(partner.StorageCapabilities), partner.Notes, partner.Avatar, time.Now(), partner.ID).Scan(&partner.ID, &partner.NGOUserID, &partner.Name, &partner.Type, &partner.Location, &partner.DistanceKm, &partner.ContactName, &partner.ContactPhone, &partner.ContactEmail, &partner.OperatingHours, &partner.AcceptanceRate, &partner.LastDonationAt, &partner.AvgDonationKg, pq.Array(&partner.StorageCapabilities), &partner.Notes, &partner.Avatar, &partner.CreatedAt, &partner.UpdatedAt)
}
