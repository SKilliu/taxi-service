package orders

import (
	"net/http"

	"github.com/SKilliu/taxi-service/server/middlewares"

	"github.com/SKilliu/taxi-service/server/dto"
	"github.com/SKilliu/taxi-service/server/errs"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateOrder(c echo.Context) error {
	var (
		req dto.CreateOrderReq
	)

	accountType, _, err := middlewares.GetAccountTypeFromJWT(c.Request(), h.auth)
	if err != nil {
		h.log.WithError(err).Error("failed to get account type from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if accountType == dto.OperatorRole {
		h.log.WithError(errs.IncorrectAccountTypeErr.ToError()).Error("incorrect account type in token")
		return c.JSON(http.StatusForbidden, errs.IncorrectAccountTypeErr)
	}

	err = c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse create order request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	return c.NoContent(http.StatusOK)
}
