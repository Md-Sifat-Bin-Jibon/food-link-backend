package surplus

import (
	"encoding/json"
	"foodlink_backend/errors"
	"foodlink_backend/features/auth"
	"foodlink_backend/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) getUserID(r *http.Request) (uuid.UUID, error) {
	user, ok := r.Context().Value("user").(*auth.User)
	if !ok || user == nil {
		return uuid.Nil, errors.ErrUnauthorized
	}
	return user.ID, nil
}

// GetAll handles GET /api/v1/restaurant/surplus
// @Summary      List surplus items
// @Description  Get all surplus items for the authenticated restaurant
// @Tags         restaurant-surplus
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   RestaurantSurplusItem
// @Failure      401  {object}  errors.AppError
// @Router       /restaurant/surplus [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.BadRequestResponse(w, "Method not allowed", nil)
		return
	}
	userID, err := h.getUserID(r)
	if err != nil {
		utils.UnauthorizedResponse(w, "Authentication required")
		return
	}
	items, err := h.service.GetAllByUserID(userID)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			utils.ErrorResponse(w, appErr.Code, appErr.Message, nil)
			return
		}
		utils.InternalServerErrorResponse(w, "Failed to retrieve surplus items", err.Error())
		return
	}
	utils.OKResponse(w, "Surplus items retrieved successfully", items)
}

// Create handles POST /api/v1/restaurant/surplus
// @Summary      Create surplus item
// @Description  Create a new surplus item
// @Tags         restaurant-surplus
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CreateRestaurantSurplusItemRequest  true  "Surplus item data"
// @Success      201      {object}  RestaurantSurplusItem
// @Failure      400      {object}  errors.AppError
// @Failure      401      {object}  errors.AppError
// @Router       /restaurant/surplus [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.BadRequestResponse(w, "Method not allowed", nil)
		return
	}
	userID, err := h.getUserID(r)
	if err != nil {
		utils.UnauthorizedResponse(w, "Authentication required")
		return
	}
	var req CreateRestaurantSurplusItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.BadRequestResponse(w, "Invalid request body", err.Error())
		return
	}
	item, err := h.service.Create(userID, &req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			utils.ErrorResponse(w, appErr.Code, appErr.Message, nil)
			return
		}
		utils.InternalServerErrorResponse(w, "Failed to create surplus item", err.Error())
		return
	}
	utils.CreatedResponse(w, "Surplus item created successfully", item)
}

// Update handles PUT /api/v1/restaurant/surplus/:id
// @Summary      Update surplus item
// @Description  Update an existing surplus item
// @Tags         restaurant-surplus
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string                        true  "Surplus Item ID"
// @Param        request  body      UpdateRestaurantSurplusItemRequest true  "Surplus item data"
// @Success      200      {object}  RestaurantSurplusItem
// @Failure      400      {object}  errors.AppError
// @Failure      401      {object}  errors.AppError
// @Failure      403      {object}  errors.AppError
// @Failure      404      {object}  errors.AppError
// @Router       /restaurant/surplus/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.BadRequestResponse(w, "Method not allowed", nil)
		return
	}
	userID, err := h.getUserID(r)
	if err != nil {
		utils.UnauthorizedResponse(w, "Authentication required")
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/restaurant/surplus/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.BadRequestResponse(w, "Invalid ID format", nil)
		return
	}
	var req UpdateRestaurantSurplusItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.BadRequestResponse(w, "Invalid request body", err.Error())
		return
	}
	item, err := h.service.Update(id, userID, &req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			utils.ErrorResponse(w, appErr.Code, appErr.Message, nil)
			return
		}
		utils.InternalServerErrorResponse(w, "Failed to update surplus item", err.Error())
		return
	}
	utils.OKResponse(w, "Surplus item updated successfully", item)
}

// Assign handles PUT /api/v1/restaurant/surplus/:id/assign
// @Summary      Assign surplus item
// @Description  Assign a surplus item to NGO or community kitchen
// @Tags         restaurant-surplus
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string                  true  "Surplus Item ID"
// @Param        request  body      AssignSurplusItemRequest true  "Assignment data"
// @Success      200      {object}  RestaurantSurplusItem
// @Failure      400      {object}  errors.AppError
// @Failure      401      {object}  errors.AppError
// @Failure      403      {object}  errors.AppError
// @Router       /restaurant/surplus/{id}/assign [put]
func (h *Handler) Assign(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.BadRequestResponse(w, "Method not allowed", nil)
		return
	}
	userID, err := h.getUserID(r)
	if err != nil {
		utils.UnauthorizedResponse(w, "Authentication required")
		return
	}
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/v1/restaurant/surplus/"), "/")
	if len(pathParts) < 2 || pathParts[1] != "assign" {
		utils.BadRequestResponse(w, "Invalid path", nil)
		return
	}
	id, err := uuid.Parse(pathParts[0])
	if err != nil {
		utils.BadRequestResponse(w, "Invalid ID format", nil)
		return
	}
	var req AssignSurplusItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.BadRequestResponse(w, "Invalid request body", err.Error())
		return
	}
	item, err := h.service.Assign(id, userID, &req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			utils.ErrorResponse(w, appErr.Code, appErr.Message, nil)
			return
		}
		utils.InternalServerErrorResponse(w, "Failed to assign surplus item", err.Error())
		return
	}
	utils.OKResponse(w, "Surplus item assigned successfully", item)
}
