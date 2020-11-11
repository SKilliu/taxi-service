package orders

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/server/errs"
	"github.com/SKilliu/taxi-service/server/middlewares"

	"github.com/SKilliu/taxi-service/server/dto"
	"github.com/labstack/echo/v4"
)

// OrderActions godoc
// @Security bearerAuth
// @Summary Actions with order
// @Tags orders
// @Consume application/json
// @Param JSON body dto.OrderActionsReq true "Body for order actions"
// @Description Accept or close the order
// @Accept  json
// @Produce  json
// @Success 200 {} http.StatusOk
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /orders [patch]
func (h *Handler) OrderActions(c echo.Context) error {
	var req dto.OrderActionsReq

	accountType, _, err := middlewares.GetAccountTypeFromJWT(c.Request(), h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get account type from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if accountType != dto.DriverRole {
		h.log.WithError(errs.IncorrectAccountTypeErr.ToError()).Error("incorrect account type in token")
		return c.JSON(http.StatusForbidden, errs.IncorrectAccountTypeErr)
	}

	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	order, err := h.ordersDB.GetByID(req.OrderID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("order doesn't exist")
			return c.JSON(http.StatusBadRequest, errs.CarDoesntExistErr)
		default:
			h.log.WithError(err).Error("failed to get order from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	switch req.Action {
	case dto.ActionAcceptOrder:
		if order.Status != dto.StatusAvailable {
			h.log.WithError(errs.IncorrectOrderErr.ToError()).Error("incorrect order")
			return c.JSON(http.StatusBadRequest, errs.IncorrectOrderErr)
		}

		car, err := h.carsDB.GetAvailableCar(userID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				h.log.WithError(err).Error("user hasn't available cars")
				return c.JSON(http.StatusForbidden, errs.HasNoAvailableCarsErr)
			default:
				h.log.WithError(err).Error("failed to get available car from db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}
		}

		order.DriverID = userID
		order.CarID = car.ID
		order.Status = dto.StatusInProgress

		err = h.ordersDB.Update(order)
		if err != nil {
			h.log.WithError(err).Error("failed to update an order in db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}

		car.Status = dto.StatusBusy
		err = h.carsDB.Update(car)
		if err != nil {
			h.log.WithError(err).Error("failed to update a car in db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}

	case dto.ActionCloseOrder:
		if order.Status != dto.StatusInProgress {
			h.log.WithError(errs.IncorrectOrderErr.ToError()).Error("incorrect order")
			return c.JSON(http.StatusBadRequest, errs.IncorrectOrderErr)
		}

		order.Status = dto.StatusClosed
		err = h.ordersDB.Update(order)
		if err != nil {
			h.log.WithError(err).Error("failed to update an order in db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}

		car, err := h.carsDB.GetByID(order.CarID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				h.log.WithError(err).Error("user hasn't available cars")
				return c.JSON(http.StatusForbidden, errs.HasNoAvailableCarsErr)
			default:
				h.log.WithError(err).Error("failed to get available car from db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}
		}

		car.Status = dto.StatusAvailable
		err = h.carsDB.Update(car)
		if err != nil {
			h.log.WithError(err).Error("failed to update a car in db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}

	default:
		h.log.WithError(err).Error("empty action for order")
		return c.JSON(http.StatusForbidden, errs.EmptyActionForOrderErr)
	}

	return c.NoContent(http.StatusOK)
}
