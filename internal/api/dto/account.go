package dto

type AuthReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Received struct {
	FromUser string `json:"from_user"`
	Amount   int64  `json:"amount"`
}

type Sent struct {
	ToUser string `json:"to_user"`
	Amount int64  `json:"amount"`
}

type CoinsHistory struct {
	Received Received `json:"received"`
	Sent     Sent     `json:"sent"`
}
