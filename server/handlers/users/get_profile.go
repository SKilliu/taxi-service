package users

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/middlewares"
	"github.com/SKilliu/taxi-service/server/errs"

	"github.com/labstack/echo/v4"
)

const (
	userID     = "user_id"
	ownerView  = "owner_view"
	publicView = "public_view"
)

type GetProfileResp struct {
	TypeView        string `json:"type_view" example:"owner_view"`
	ID              string `json:"id" example:"Yhte-saaiudchadsc-asdvcsf"`
	Name            string `json:"name" example:"Tester"`
	Email           string `json:"email" example:"test@example.com"`
	ProfileImageUrl string `json:"profile_image_url" example:"http://simple-service-backend/simple-service/photo-924y82hde7ce.jpg"`
} //@name GetProfileResp

// GetProfile godoc
// @Security bearerAuth
// @Summary Get profile
// @Tags user
// @Consume application/json
// @Param user_id query string true "user ID for getting profile"
// @Description Get user's profile by ID
// @Accept json
// @Produce json
// @Success 200 {object} GetProfileResp
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /user [get]
func (h *Handler) GetProfile(c echo.Context) error {
	var (
		resp     GetProfileResp
		typeView string
	)

	userID := c.QueryParam(userID)

	reqOwnerID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed ot get user ID from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	user, err := h.usersDB.GetByID(userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(errs.UserDoesntExistErr.ToError()).Error("user doesn't exist")
			return c.JSON(http.StatusBadRequest, errs.UserDoesntExistErr)
		default:
			h.log.WithError(err).Error("failed to get user from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	if userID == reqOwnerID {
		typeView = ownerView
	} else {
		typeView = publicView
	}

	resp = GetProfileResp{
		TypeView:        typeView,
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		ProfileImageUrl: user.ProfileImageUrl,
	}

	return c.JSON(http.StatusOK, resp)
}
