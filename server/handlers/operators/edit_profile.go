package operators

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/server/errs"
	"github.com/SKilliu/taxi-service/server/middlewares"
	"github.com/labstack/echo/v4"
)

type EditProfileReq struct {
	Name  string `json:"name" example:"Tester"`
	Email string `json:"email" example:"new-email@example.com"`
} //@name EditProfileReq

// EditProfile godoc
// @Security bearerAuth
// @Summary Edit profile
// @Tags user
// @Consume application/json
// @Param JSON body EditProfileReq true "Body for edit profile request"
// @Description Edit user's name and email in profile info
// @Accept  json
// @Produce  json
// @Success 200 {} http.StatusOk
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /user [post]
func (h *Handler) EditProfile(c echo.Context) error {
	var req EditProfileReq

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse edit profile request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

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

	if req.Email != "" {
		_, err := h.usersDB.GetByEmail(req.Email)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				user.Email = req.Email
			default:
				h.log.WithError(err).Error("failed to get user from db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}
		} else {
			h.log.WithError(errs.UserDoesntExistErr.ToError()).Error("Another user has this email")
			return c.JSON(http.StatusForbidden, errs.UserAlreadyExistErr)
		}
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	err = h.usersDB.Update(user)
	if err != nil {
		h.log.WithError(err).Error("failed to update the user in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
