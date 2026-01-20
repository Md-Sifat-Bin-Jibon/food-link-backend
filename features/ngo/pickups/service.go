package pickups

import (
	"foodlink_backend/errors"
	"foodlink_backend/utils"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService() *Service {
	return &Service{repo: NewRepository()}
}

func (s *Service) GetAllByOfferID(offerID uuid.UUID) ([]*NGOPickupSchedule, error) {
	return s.repo.GetAllByOfferID(offerID)
}

func (s *Service) GetByID(id uuid.UUID) (*NGOPickupSchedule, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Create(req *CreateNGOPickupScheduleRequest) (*NGOPickupSchedule, error) {
	if validationErrors := utils.ValidateStruct(req); len(validationErrors) > 0 {
		return nil, errors.NewAppErrorWithErr(errors.ErrValidationFailed.Code, "Validation failed: "+validationErrors[0], nil)
	}
	schedule := &NGOPickupSchedule{
		ID:              uuid.New(),
		OfferID:         req.OfferID,
		ScheduledFor:   req.ScheduledFor,
		ETAMinutes:      req.ETAMinutes,
		VolunteerName:   req.VolunteerName,
		VolunteerContact: req.VolunteerContact,
		VehicleType:     req.VehicleType,
		Status:          "scheduled",
		Checkpoints:     JSONB(req.Checkpoints),
		Reminders:       JSONB(req.Reminders),
		Notes:           req.Notes,
	}
	if err := s.repo.Create(schedule); err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *Service) Update(id uuid.UUID, req *UpdateNGOPickupScheduleRequest) (*NGOPickupSchedule, error) {
	schedule, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if validationErrors := utils.ValidateStruct(req); len(validationErrors) > 0 {
		return nil, errors.NewAppErrorWithErr(errors.ErrValidationFailed.Code, "Validation failed: "+validationErrors[0], nil)
	}
	if req.ScheduledFor != nil {
		schedule.ScheduledFor = *req.ScheduledFor
	}
	if req.ETAMinutes != nil {
		schedule.ETAMinutes = req.ETAMinutes
	}
	if req.VolunteerName != "" {
		schedule.VolunteerName = req.VolunteerName
	}
	if req.VolunteerContact != "" {
		schedule.VolunteerContact = req.VolunteerContact
	}
	if req.VehicleType != "" {
		schedule.VehicleType = req.VehicleType
	}
	if req.Checkpoints != nil {
		schedule.Checkpoints = JSONB(req.Checkpoints)
	}
	if req.Reminders != nil {
		schedule.Reminders = JSONB(req.Reminders)
	}
	if req.Notes != "" {
		schedule.Notes = req.Notes
	}
	if err := s.repo.Update(schedule); err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *Service) UpdateStatus(id uuid.UUID, req *UpdatePickupStatusRequest) (*NGOPickupSchedule, error) {
	schedule, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if validationErrors := utils.ValidateStruct(req); len(validationErrors) > 0 {
		return nil, errors.NewAppErrorWithErr(errors.ErrValidationFailed.Code, "Validation failed: "+validationErrors[0], nil)
	}
	schedule.Status = req.Status
	if err := s.repo.Update(schedule); err != nil {
		return nil, err
	}
	return schedule, nil
}
