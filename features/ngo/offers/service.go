package offers

import (
	"foodlink_backend/errors"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService() *Service {
	return &Service{repo: NewRepository()}
}

func (s *Service) GetAllByNGOUserID(ngoUserID uuid.UUID, status string) ([]*NGODonationOffer, error) {
	return s.repo.GetAllByNGOUserID(ngoUserID, status)
}

func (s *Service) GetByID(id uuid.UUID) (*NGODonationOffer, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Accept(id uuid.UUID, ngoUserID uuid.UUID) (*NGODonationOffer, error) {
	offer, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if offer.NGOUserID != ngoUserID {
		return nil, errors.ErrForbidden
	}
	if err := s.repo.UpdateStatus(id, "accepted"); err != nil {
		return nil, err
	}
	offer.Status = "accepted"
	return offer, nil
}

func (s *Service) Decline(id uuid.UUID, ngoUserID uuid.UUID) (*NGODonationOffer, error) {
	offer, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if offer.NGOUserID != ngoUserID {
		return nil, errors.ErrForbidden
	}
	if err := s.repo.UpdateStatus(id, "declined"); err != nil {
		return nil, err
	}
	offer.Status = "declined"
	return offer, nil
}
