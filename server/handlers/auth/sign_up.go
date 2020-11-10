package auth

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/server/dto"

	"golang.org/x/crypto/bcrypt"

	"github.com/SKilliu/taxi-service/db/models"
	"github.com/SKilliu/taxi-service/server/errs"
	"github.com/SKilliu/taxi-service/server/handlers"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

// SignUp godoc
// @Summary Sign up
// @Tags authentication
// @Consume application/json
// @Param JSON body dto.SignUpReq true "Body for sign up"
// @Description Sign up with login, password and account type (driver, client or operator)
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.AuthResp
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /sign_up [post]
func (h *Handler) SignUp(c echo.Context) error {
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

			token, err := handlers.GenerateJWT(userID, req.AccountType, h.auth.VerifyKey)
			if err != nil {
				h.log.WithError(err).Error("failed to generate token")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			return c.JSON(http.StatusOK, dto.AuthResp{
				Token: token,
			})

		default:
			h.log.WithError(err).Error("failed to get user from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	h.log.WithError(errs.UserAlreadyExistErr.ToError()).Error("user already exist")

	return c.JSON(http.StatusForbidden, errs.UserAlreadyExistErr)
}
