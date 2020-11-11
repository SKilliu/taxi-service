package users

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/server/errs"
	"github.com/SKilliu/taxi-service/server/middlewares"

	"github.com/labstack/echo/v4"
)

// Delete godoc
// @Security bearerAuth
// @Summary Delete profile
// @Tags users
// @Consume application/json
// @Description Delete user profile from db
// @Accept  json
// @Produce  json
// @Success 200 {} http.StatusOk
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /user [delete]
func (h *Handler) Delete(c echo.Context) error {
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

	err = h.usersDB.Delete(user)
	if err != nil {
		h.log.WithError(err).Error("failed to delete user from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
