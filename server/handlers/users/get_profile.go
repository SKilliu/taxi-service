package users

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/server/dto"

	"github.com/SKilliu/taxi-service/server/errs"
	"github.com/SKilliu/taxi-service/server/middlewares"

	"github.com/labstack/echo/v4"
)

const userID = "user_id"

type GetProfileResp struct {
	ID              string `json:"id" example:"Yhte-saaiudchadsc-asdvcsf"`
	Name            string `json:"name" example:"Tester"`
	Email           string `json:"email" example:"test@example.com"`
	AccountType     string `json:"account_type"`
	ProfileImageUrl string `json:"profile_image_url" example:"http://simple-service-backend/simple-service/photo-924y82hde7ce.jpg"`
} //@name GetProfileResp

// GetProfile godoc
// @Security bearerAuth
// @Summary Get profile
// @Tags users
// @Consume application/json
// @Description Get your profile
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetProfileResp
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /user [get]
func (h *Handler) GetProfile(c echo.Context) error {
	var resp dto.GetProfileResp

	userID := c.QueryParam(userID)

	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.auth)
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

	resp = dto.GetProfileResp{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		AccountType:     user.AccountType,
		ProfileImageUrl: user.ProfileImageUrl,
	}

	return c.JSON(http.StatusOK, resp)
}
