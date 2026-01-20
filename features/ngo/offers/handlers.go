package offers

import (
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

// GetAll handles GET /api/v1/ngo/offers
// @Summary      List donation offers
// @Description  Get all donation offers for the authenticated NGO
// @Tags         ngo-offers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        status  query     string  false  "Filter by status"
// @Success      200     {array}   NGODonationOffer
// @Failure      401     {object}  errors.AppError
// @Router       /ngo/offers [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.BadRequestResponse(w, "Method not allowed", nil)
		return
	}
	ngoUserID, err := h.getUserID(r)
	if err != nil {
		utils.UnauthorizedResponse(w, "Authentication required")
		return
	}
	status := r.URL.Query().Get("status")
	offers, err := h.service.GetAllByNGOUserID(ngoUserID, status)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			utils.ErrorResponse(w, appErr.Code, appErr.Message, nil)
			return
		}
		utils.InternalServerErrorResponse(w, "Failed to retrieve offers", err.Error())
		return
	}
	utils.OKResponse(w, "Offers retrieved successfully", offers)
}

// GetByID handles GET /api/v1/ngo/offers/:id
// @Summary      Get offer details
// @Description  Get details of a specific donation offer
// @Tags         ngo-offers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Offer ID"
// @Success      200  {object}  NGODonationOffer
// @Failure      401  {object}  errors.AppError
// @Failure      404  {object}  errors.AppError
// @Router       /ngo/offers/{id} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.BadRequestResponse(w, "Method not allowed", nil)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/ngo/offers/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.BadRequestResponse(w, "Invalid ID format", nil)
		return
	}
	offer, err := h.service.GetByID(id)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			utils.ErrorResponse(w, appErr.Code, appErr.Message, nil)
			return
		}
		utils.InternalServerErrorResponse(w, "Failed to retrieve offer", err.Error())
		return
	}
	utils.OKResponse(w, "Offer retrieved successfully", offer)
}

// Accept handles PUT /api/v1/ngo/offers/:id/accept
// @Summary      Accept offer
// @Description  Accept a donation offer
// @Tags         ngo-offers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Offer ID"
// @Success      200  {object}  NGODonationOffer
// @Failure      401  {object}  errors.AppError
// @Failure      403  {object}  errors.AppError
// @Failure      404  {object}  errors.AppError
// @Router       /ngo/offers/{id}/accept [put]
func (h *Handler) Accept(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.BadRequestResponse(w, "Method not allowed", nil)
		return
	}
	ngoUserID, err := h.getUserID(r)
	if err != nil {
		utils.UnauthorizedResponse(w, "Authentication required")
		return
	}
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/v1/ngo/offers/"), "/")
	if len(pathParts) < 2 || pathParts[1] != "accept" {
		utils.BadRequestResponse(w, "Invalid path", nil)
		return
	}
	id, err := uuid.Parse(pathParts[0])
	if err != nil {
		utils.BadRequestResponse(w, "Invalid ID format", nil)
		return
	}
	offer, err := h.service.Accept(id, ngoUserID)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			utils.ErrorResponse(w, appErr.Code, appErr.Message, nil)
			return
		}
		utils.InternalServerErrorResponse(w, "Failed to accept offer", err.Error())
		return
	}
	utils.OKResponse(w, "Offer accepted successfully", offer)
}

// Decline handles PUT /api/v1/ngo/offers/:id/decline
// @Summary      Decline offer
// @Description  Decline a donation offer
// @Tags         ngo-offers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Offer ID"
// @Success      200  {object}  NGODonationOffer
// @Failure      401  {object}  errors.AppError
// @Failure      403  {object}  errors.AppError
// @Failure      404  {object}  errors.AppError
// @Router       /ngo/offers/{id}/decline [put]
func (h *Handler) Decline(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.BadRequestResponse(w, "Method not allowed", nil)
		return
	}
	ngoUserID, err := h.getUserID(r)
	if err != nil {
		utils.UnauthorizedResponse(w, "Authentication required")
		return
	}
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/v1/ngo/offers/"), "/")
	if len(pathParts) < 2 || pathParts[1] != "decline" {
		utils.BadRequestResponse(w, "Invalid path", nil)
		return
	}
	id, err := uuid.Parse(pathParts[0])
	if err != nil {
		utils.BadRequestResponse(w, "Invalid ID format", nil)
		return
	}
	offer, err := h.service.Decline(id, ngoUserID)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			utils.ErrorResponse(w, appErr.Code, appErr.Message, nil)
			return
		}
		utils.InternalServerErrorResponse(w, "Failed to decline offer", err.Error())
		return
	}
	utils.OKResponse(w, "Offer declined successfully", offer)
}
