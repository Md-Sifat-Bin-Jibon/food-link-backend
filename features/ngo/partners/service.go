package partners

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

func (s *Service) GetAllByNGOUserID(ngoUserID uuid.UUID) ([]*NGOPartnerProfile, error) {
	return s.repo.GetAllByNGOUserID(ngoUserID)
}

func (s *Service) GetByID(id uuid.UUID) (*NGOPartnerProfile, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Create(ngoUserID uuid.UUID, req *CreateNGOPartnerRequest) (*NGOPartnerProfile, error) {
	if validationErrors := utils.ValidateStruct(req); len(validationErrors) > 0 {
		return nil, errors.NewAppErrorWithErr(errors.ErrValidationFailed.Code, "Validation failed: "+validationErrors[0], nil)
	}
	partner := &NGOPartnerProfile{
		ID:                 uuid.New(),
		NGOUserID:          ngoUserID,
		Name:               req.Name,
		Type:               req.Type,
		Location:           req.Location,
		DistanceKm:         req.DistanceKm,
		ContactName:        req.ContactName,
		ContactPhone:       req.ContactPhone,
		ContactEmail:       req.ContactEmail,
		OperatingHours:     req.OperatingHours,
		AcceptanceRate:     0,
		AvgDonationKg:      0,
		StorageCapabilities: req.StorageCapabilities,
		Notes:               req.Notes,
		Avatar:              req.Avatar,
	}
	if err := s.repo.Create(partner); err != nil {
		return nil, err
	}
	return partner, nil
}

func (s *Service) Update(id uuid.UUID, ngoUserID uuid.UUID, req *UpdateNGOPartnerRequest) (*NGOPartnerProfile, error) {
	partner, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if partner.NGOUserID != ngoUserID {
		return nil, errors.ErrForbidden
	}
	if validationErrors := utils.ValidateStruct(req); len(validationErrors) > 0 {
		return nil, errors.NewAppErrorWithErr(errors.ErrValidationFailed.Code, "Validation failed: "+validationErrors[0], nil)
	}
	if req.Name != "" {
		partner.Name = req.Name
	}
	if req.Location != "" {
		partner.Location = req.Location
	}
	if req.DistanceKm != nil {
		partner.DistanceKm = req.DistanceKm
	}
	if req.ContactName != "" {
		partner.ContactName = req.ContactName
	}
	if req.ContactPhone != "" {
		partner.ContactPhone = req.ContactPhone
	}
	if req.ContactEmail != "" {
		partner.ContactEmail = req.ContactEmail
	}
	if req.OperatingHours != "" {
		partner.OperatingHours = req.OperatingHours
	}
	if req.StorageCapabilities != nil {
		partner.StorageCapabilities = req.StorageCapabilities
	}
	if req.Notes != "" {
		partner.Notes = req.Notes
	}
	if req.Avatar != "" {
		partner.Avatar = req.Avatar
	}
	if err := s.repo.Update(partner); err != nil {
		return nil, err
	}
	return partner, nil
}
