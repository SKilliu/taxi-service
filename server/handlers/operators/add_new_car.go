package operators

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/server/middlewares"

	"github.com/SKilliu/taxi-service/db/models"
	"github.com/google/uuid"

	"github.com/SKilliu/taxi-service/server/dto"
	"github.com/SKilliu/taxi-service/server/errs"

	"github.com/labstack/echo/v4"
)

// AddNewCar godoc
// @Security bearerAuth
// @Summary Edit profile
// @Tags operators
// @Consume application/json
// @Param JSON body dto.AddCarReq true "Body for add new car request"
// @Description Add new car to database
// @Accept  json
// @Produce  json
// @Success 200 {} http.StatusOk
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /operators/car [post]
func (h *Handler) AddNewCar(c echo.Context) error {
	var req dto.AddCarReq

	accountType, _, err := middlewares.GetAccountTypeFromJWT(c.Request(), h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get account type from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if accountType != dto.OperatorRole {
		h.log.WithError(errs.IncorrectAccountTypeErr.ToError()).Error("incorrect account type in token")
		return c.JSON(http.StatusForbidden, errs.IncorrectAccountTypeErr)
	}

	err = c.Bind(&req)
	if err != nil {
		h.log.WithError(errs.BadParamInBodyErr.ToError()).Error("failed to parse add new car request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	// We don't add new car to db, if it already exists
	_, err = h.carsDB.GetByNumber(req.Number)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = h.carsDB.Insert(models.Car{
				ID:     uuid.New().String(),
				Model:  req.Model,
				Number: req.Number,
				Status: req.Status,
			})
			if err != nil {
				h.log.WithError(err).Error("failed to insert new car into db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			return c.NoContent(http.StatusOK)
		default:
			h.log.WithError(err).Error("failed to get car by number from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	return c.JSON(http.StatusBadRequest, errs.CarAlreadyExistsErr)
}
