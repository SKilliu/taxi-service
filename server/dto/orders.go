package dto

const pricePerKilometer = "1.0"

type CreateOrderReq struct {
	StartingPointLongitude    float64 `json:"starting_point_longitude" example:"12.12345"`
	StartingPointLatitude     float64 `json:"starting_point_latitude" example:"12.12345"`
	DestinationPointLongitude float64 `json:"destination_point_longitude" example:"12.12345"`
	DestinationPointLatitude  float64 `json:"destination_point_latitude" example:"12.12345"`
	DestinationLocation       string  `json:"destination_location" example:"Some City, some Street, 123"`
} //@name CreateOrderReq
