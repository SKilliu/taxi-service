package users

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/server/dto"

	"github.com/SKilliu/taxi-service/db/models"
	"github.com/SKilliu/taxi-service/server/errs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) CreateNewUser(c echo.Context) error {
	var req dto.SignUpReq

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(errs.BadParamInBodyErr.ToError()).Error("failed to parse sign up request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	err = req.Validate()
	if err != nil {
		h.log.WithError(errs.NotValidBodyParamErr.ToError()).Error("not valid param in body")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	_, err = h.usersDB.GetByEmail(req.Email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:

			passwordBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
			if err != nil {
				h.log.WithError(err).Error("failed to encode user password")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			userID := uuid.New().String()

			err = h.usersDB.Insert(models.User{
				ID:             userID,
				Name:           req.Name,
				AccountType:    req.AccountType,
				HashedPassword: string(passwordBytes),
				Email:          req.Email,
			})
			if err != nil {
				h.log.WithError(err).Error("failed to insert a new user into db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			return c.NoContent(http.StatusOK)

		default:
			h.log.WithError(err).Error("failed to get user from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	h.log.WithError(errs.UserAlreadyExistErr.ToError()).Error("user already exist")

	return c.JSON(http.StatusForbidden, errs.UserAlreadyExistErr)
}
