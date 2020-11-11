package orders

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/taxi-service/server/dto"
	"github.com/SKilliu/taxi-service/server/errs"
	"github.com/SKilliu/taxi-service/server/middlewares"
	"github.com/labstack/echo/v4"
)

// GetAvailableOrders godoc
// @Security bearerAuth
// @Summary Get available orders
// @Tags orders
// @Consume application/json
// @Description Create a new order as a client
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.GetAvailableOrdersResp
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /orders [get]
func (h *Handler) GetAvailableOrders(c echo.Context) error {
	var resp []dto.GetAvailableOrdersResp

	accountType, _, err := middlewares.GetAccountTypeFromJWT(c.Request(), h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get account type from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if accountType != dto.DriverRole {
		h.log.WithError(errs.IncorrectAccountTypeErr.ToError()).Error("incorrect account type in token")
		return c.JSON(http.StatusForbidden, errs.IncorrectAccountTypeErr)
	}

	availableOrders, err := h.ordersDB.GetWithAvailableStatus()
	if err != nil {
		h.log.WithError(err).Error("failed to get orders with available status from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	for _, ao := range availableOrders {
		trip, err := h.tripsDB.GetByID(ao.TripID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				h.log.WithError(err).Error("trip doesn't exist")
				return c.JSON(http.StatusBadRequest, errs.CarDoesntExistErr)
			default:
				h.log.WithError(err).Error("failed to get trip from db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}
		}

		resp = append(resp, dto.GetAvailableOrdersResp{
			OrderID:  ao.ID,
			ClientID: ao.ClientID,
			TripInfo: dto.TripInfo{
				StartingPoint: dto.Location{
					Address:   trip.StartingPointLocation,
					Latitude:  trip.StartingPointLatitude,
					Longitude: trip.StartingPointLongitude,
				},
				DestinationPoint: dto.Location{
					Address:   trip.DestinationPointLocation,
					Latitude:  trip.DestinationPointLatitude,
					Longitude: trip.DestinationPointLongitude,
				},
				Distance: trip.Distance,
			},
			Price:          ao.Price,
			CarArrivalTime: ao.CarArrivalTime,
		})

	}

	return c.JSON(http.StatusOK, resp)
}
