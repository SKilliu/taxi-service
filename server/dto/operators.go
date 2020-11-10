package dto

const (
	StatusAvailable = "available"
	StatusBusy      = "busy"
)

type AddCarReq struct {
	Model    string `json:"model"`
	Number   string `json:"number"`
	Status   string `json:"status"`
	ImageUrl string `json:"image_url"`
} //@name AddCarReq
