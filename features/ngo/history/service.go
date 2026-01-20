package history

import (
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService() *Service {
	return &Service{repo: NewRepository()}
}

func (s *Service) GetAllByNGOUserID(ngoUserID uuid.UUID) ([]*NGODonationHistory, error) {
	return s.repo.GetAllByNGOUserID(ngoUserID)
}

func (s *Service) GetByID(id uuid.UUID) (*NGODonationHistory, error) {
	return s.repo.GetByID(id)
}
