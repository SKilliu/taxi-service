package dto

const (
	StatusAvailable = "available"
	StatusBusy      = "busy"
)

type AddCarReq struct {
	Model  string `json:"model" example:"BMW"`
	Series string `json:"series" example:"M5"`
	Number string `json:"number" example:"AX1234XA"`
	Status string `json:"status" example:"available"`
} //@name AddCarReq

type AssignCarToDriverReq struct {
	DriverID  string `json:"driver_id"`
	CarNumber string `json:"car_id"`
}
