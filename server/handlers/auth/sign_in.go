package auth

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/handlers"

	"golang.org/x/crypto/bcrypt"

	"github.com/SKilliu/taxi-service/server/errs"
	"github.com/labstack/echo/v4"
)

type SignInReq struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"qwerty1234"`
} //@name SignInReq

type SignInResp struct {
	Token string `json:"token"`
} //@name SignInResp

// SignIn godoc
// @Summary Sign in
// @Tags authentication
// @Consume application/json
// @Param JSON body SignInReq true "Body for sign in"
// @Description Sign in with login and password
// @Accept  json
// @Produce  json
// @Success 200 {object} SignInResp
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /sign_in [post]
func (h *Handler) SignIn(c echo.Context) error {
	var req SignInReq

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse sign up request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	user, err := h.usersDB.GetByEmail(req.Email)
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

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		h.log.WithError(err).Error("wrong email or password")
		return c.JSON(http.StatusBadRequest, errs.WrongCredentialsErr)
	}

	token, err := handlers.GenerateJWT(user.ID, h.auth.VerifyKey)
	if err != nil {
		h.log.WithError(err).Error("failed to generate token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, SignInResp{
		Token: token,
	})
}
