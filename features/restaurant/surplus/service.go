package surplus

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

func (s *Service) GetAllByUserID(userID uuid.UUID) ([]*RestaurantSurplusItem, error) {
	return s.repo.GetAllByUserID(userID)
}

func (s *Service) GetByID(id uuid.UUID) (*RestaurantSurplusItem, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Create(userID uuid.UUID, req *CreateRestaurantSurplusItemRequest) (*RestaurantSurplusItem, error) {
	if validationErrors := utils.ValidateStruct(req); len(validationErrors) > 0 {
		return nil, errors.NewAppErrorWithErr(errors.ErrValidationFailed.Code, "Validation failed: "+validationErrors[0], nil)
	}
	item := &RestaurantSurplusItem{
		ID:           uuid.New(),
		UserID:       userID,
		Title:        req.Title,
		Description:  req.Description,
		Quantity:     req.Quantity,
		Unit:         req.Unit,
		Category:     req.Category,
		StorageType:  req.StorageType,
		PickupWindow: JSONB(req.PickupWindow),
		Tags:         req.Tags,
		Image:        req.Image,
		Status:       "pending",
	}
	if err := s.repo.Create(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(id uuid.UUID, userID uuid.UUID, req *UpdateRestaurantSurplusItemRequest) (*RestaurantSurplusItem, error) {
	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if item.UserID != userID {
		return nil, errors.ErrForbidden
	}
	if validationErrors := utils.ValidateStruct(req); len(validationErrors) > 0 {
		return nil, errors.NewAppErrorWithErr(errors.ErrValidationFailed.Code, "Validation failed: "+validationErrors[0], nil)
	}
	if req.Title != "" {
		item.Title = req.Title
	}
	if req.Description != "" {
		item.Description = req.Description
	}
	if req.Quantity != nil {
		item.Quantity = *req.Quantity
	}
	if req.Unit != "" {
		item.Unit = req.Unit
	}
	if req.Category != "" {
		item.Category = req.Category
	}
	if req.StorageType != "" {
		item.StorageType = req.StorageType
	}
	if req.PickupWindow != nil {
		item.PickupWindow = JSONB(req.PickupWindow)
	}
	if req.Tags != nil {
		item.Tags = req.Tags
	}
	if req.Image != "" {
		item.Image = req.Image
	}
	if req.Status != "" {
		item.Status = req.Status
	}
	if err := s.repo.Update(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Assign(id uuid.UUID, userID uuid.UUID, req *AssignSurplusItemRequest) (*RestaurantSurplusItem, error) {
	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if item.UserID != userID {
		return nil, errors.ErrForbidden
	}
	if validationErrors := utils.ValidateStruct(req); len(validationErrors) > 0 {
		return nil, errors.NewAppErrorWithErr(errors.ErrValidationFailed.Code, "Validation failed: "+validationErrors[0], nil)
	}
	item.AssignedTo = req.AssignedTo
	item.RecipientName = req.RecipientName
	if err := s.repo.Update(item); err != nil {
		return nil, err
	}
	return item, nil
}
