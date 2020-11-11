package dto

import "time"

const (
	PricePerKilometer = 1.0
	ActionAcceptOrder = "accept"
	ActionCloseOrder  = "close"
	StatusInProgress  = "in_progress"
	StatusClosed      = "closed"
)

type CreateOrderReq struct {
	StartingPoint    Location  `json:"starting_point"`
	DestinationPoint Location  `json:"destination_point"`
	CarArrivalTime   time.Time `json:"car_arrival_time" example:"2020-11-11T23:30:00Z"`
} //@name CreateOrderReq

type Location struct {
	Address   string  `json:"address" example:"Some City, some Street, 123"`
	Latitude  float64 `json:"latitude" example:"12.12345"`
	Longitude float64 `json:"destination_point_longitude" example:"12.12345"`
} //@name Location

type GetAvailableOrdersResp struct {
	OrderID        string    `json:"id"`
	ClientID       string    `json:"client_id"`
	TripInfo       TripInfo  `json:"trip_info"`
	Price          float64   `json:"price"`
	CarArrivalTime time.Time `json:"car_arrival_time"`
} //@name GetAvailableOrdersResp

type TripInfo struct {
	StartingPoint    Location `json:"starting_point"`
	DestinationPoint Location `json:"destination_point"`
	Distance         float64  `json:"distance" example:"15"`
} //@name TripInfo

type OrderActionsReq struct {
	OrderID string `json:"order_id"`
	Action  string `json:"action"`
}
