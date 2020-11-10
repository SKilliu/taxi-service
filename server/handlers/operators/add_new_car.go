package operators

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/db/models"
	"github.com/google/uuid"

	"github.com/SKilliu/taxi-service/server/dto"
	"github.com/SKilliu/taxi-service/server/errs"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AddNewCar(c echo.Context) error {
	var req dto.AddCarReq

	err := c.Bind(&req)
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
