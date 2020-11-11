package orders

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/SKilliu/taxi-service/db/models"

	"github.com/SKilliu/taxi-service/utils"

	"github.com/SKilliu/taxi-service/server/middlewares"

	"github.com/SKilliu/taxi-service/server/dto"
	"github.com/SKilliu/taxi-service/server/errs"

	"github.com/labstack/echo/v4"
)

// CreateOrder godoc
// @Security bearerAuth
// @Summary Create order
// @Tags orders
// @Consume application/json
// @Param JSON body dto.CreateOrderReq true "Body for creating new order"
// @Description Create a new order as a client
// @Accept  json
// @Produce  json
// @Success 200 {} http.StatusOk
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /orders [post]
func (h *Handler) CreateOrder(c echo.Context) error {
	var (
		req         dto.CreateOrderReq
		CoordinateA utils.Coordinates
		CoordinateB utils.Coordinates
	)

	accountType, _, err := middlewares.GetAccountTypeFromJWT(c.Request(), h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get account type from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if accountType != dto.ClientRole {
		h.log.WithError(errs.IncorrectAccountTypeErr.ToError()).Error("incorrect account type in token")
		return c.JSON(http.StatusForbidden, errs.IncorrectAccountTypeErr)
	}

	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get user id from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse create order request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	CoordinateA = utils.Coordinates{
		Latitude:  req.StartingPoint.Latitude,
		Longitude: req.StartingPoint.Longitude,
	}

	CoordinateB = utils.Coordinates{
		Latitude:  req.DestinationPoint.Latitude,
		Longitude: req.DestinationPoint.Longitude,
	}

	_, distance := utils.DistanceCalculator(CoordinateA, CoordinateB)

	tripID := uuid.New().String()

	err = h.tripsDB.Insert(models.Trip{
		ID:                        tripID,
		StartingPointLocation:     req.StartingPoint.Address,
		StartingPointLongitude:    req.StartingPoint.Longitude,
		StartingPointLatitude:     req.StartingPoint.Latitude,
		DestinationPointLocation:  req.DestinationPoint.Address,
		DestinationPointLongitude: req.DestinationPoint.Longitude,
		DestinationPointLatitude:  req.DestinationPoint.Latitude,
		Distance:                  distance,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to insert trip to db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = h.ordersDB.Insert(models.Order{
		ID:             uuid.New().String(),
		ClientID:       userID,
		Price:          distance * dto.PricePerKilometer,
		Status:         dto.StatusAvailable,
		CarArrivalTime: req.CarArrivalTime,
		CreatedAt:      time.Now(),
	})
	if err != nil {
		h.log.WithError(err).Error("failed to insert order to db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
