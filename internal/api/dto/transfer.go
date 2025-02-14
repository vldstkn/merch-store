package dto

type SendCoinsReq struct {
	ToUser string `json:"toUser" validate:"required"`
	Amount int64  `json:"amount" validate:"required,number"`
}
