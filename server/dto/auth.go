package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	ClientRole   = "client"
	DriverRole   = "driver"
	OperatorRole = "operator"
)

type SignUpReq struct {
	Email       string `json:"email" example:"test@example.com"`
	Password    string `json:"password" example:"qwerty1234"`
	Name        string `json:"name" example:"TestName"`
	AccountType string `json:"account_type" example:"client"`
} //@name SignUpReq

type AuthResp struct {
	Token string `json:"token" example:"nausdgtGTGAjndfsKijIYbsgfsuadfe34r"`
} //@name SignUpResp

type SignInReq struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"qwerty1234"`
} //@name SignInReq

func (c SignUpReq) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.AccountType, validation.Required, validation.In(DriverRole, OperatorRole, ClientRole)),
	)
}
