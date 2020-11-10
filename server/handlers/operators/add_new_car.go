package operators

import (
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
}
