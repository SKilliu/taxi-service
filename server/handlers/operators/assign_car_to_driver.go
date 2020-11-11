package operators

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/db/models"
	"github.com/google/uuid"

	"github.com/SKilliu/taxi-service/server/middlewares"

	"github.com/SKilliu/taxi-service/server/dto"
	"github.com/SKilliu/taxi-service/server/errs"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AssignCarToDriver(c echo.Context) error {
	var req dto.AssignCarToDriverReq

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
		h.log.WithError(err).Error("failed to parse assign car to driver request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	car, err := h.carsDB.GetByNumber(req.CarNumber)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("car doesn't exist")
			return c.JSON(http.StatusBadRequest, errs.CarDoesntExistErr)
		default:
			h.log.WithError(err).Error("failed ot get car from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	err = h.driverCarsDB.Insert(models.DriverCars{
		ID:       uuid.New().String(),
		DriverID: req.DriverID,
		CarID:    car.ID,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to insert the driver car into db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
