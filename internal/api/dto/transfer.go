package dto

type SendCoinsReq struct {
	ToUser string `json:"toUser" validate:"required"`
	//TODO: number tag
	Amount int64 `json:"amount" validate:"required"`
}
