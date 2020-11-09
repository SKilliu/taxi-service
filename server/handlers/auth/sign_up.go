package auth

import (
	"database/sql"
	"net/http"
	"regexp"
	"simple-service/db/models"
	"simple-service/server/errs"
	"simple-service/server/handlers"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/labstack/echo/v4"
)

type SignUpReq struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"qwerty1234"`
} //@name SignUpReq

type SignUpResp struct {
	Token string `json:"token" example:"nausdgtGTGAjndfs.KijIYbsgfsuadfe34r"`
} //@name SignUpResp

// SignUp godoc
// @Summary Sign up
// @Tags authentication
// @Consume application/json
// @Param JSON body SignUpReq true "Body for sign up"
// @Description Sign up with login and password
// @Accept  json
// @Produce  json
// @Success 200 {object} SignUpResp
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /sign_up [post]
func (h *Handler) SignUp(c echo.Context) error {
	var req SignUpReq

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
				HashedPassword: string(passwordBytes),
				Email:          req.Email,
			})
			if err != nil {
				h.log.WithError(err).Error("failed to insert a new user into db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			token, err := handlers.GenerateJWT(userID, h.auth.VerifyKey)
			if err != nil {
				h.log.WithError(err).Error("failed to generate token")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			return c.JSON(http.StatusOK, SignUpResp{
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

func (c SignUpReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.Password, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9'-]{8,18}$"))),
		validation.Field(&c.Email, validation.Required, validation.Length(5, 70)),
	)
}
