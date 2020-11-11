package users

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/SKilliu/taxi-service/server/errs"
	"github.com/SKilliu/taxi-service/server/middlewares"

	"github.com/labstack/echo/v4"
)

// UpdateProfileImage godoc
// @Security bearerAuth
// @Summary Update profile image
// @Tags users
// @Consume multipart/form-data
// @Param profile_image formData file true "select image"
// @Description Update user's profile image
// @Accept  multipart/form-data
// @Produce json
// @Success 200 {} http.StatusOk
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /user [patch]
func (h *Handler) UpdateProfileImage(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		h.log.WithError(err).Error("failed to get data from the multipart form")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
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

	if len(form.File) == 0 {
		h.log.WithError(err).Error("multipart form is empty")
		return c.JSON(http.StatusBadRequest, errs.NoDataInFormErr)
	}

	file := form.File["profile_image"][0]
	photo, err := file.Open()
	defer photo.Close()

	urlByElements := strings.Split(user.ProfileImageUrl, "/")

	fileName := urlByElements[len(urlByElements)-1]

	if user.ProfileImageUrl != "" {
		err = h.s3.DropFile(fileName)
		if err != nil {
			h.log.WithError(err).Error("failed to delete an image from s3")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	user.ProfileImageUrl, err = h.s3.PutObject(photo, file.Size, user.ID)
	if err != nil {
		h.log.WithError(err).Error("failed to put new profile image into bucket")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = h.usersDB.Update(user)
	if err != nil {
		h.log.WithError(err).Error("failed to update user's profile photo url in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
