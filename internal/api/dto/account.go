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
	Received []Received `json:"received"`
	Sent     []Sent     `json:"sent"`
}

type Inventory struct {
	Type     string `json:"type"`
	Quantity int64  `json:"quantity"`
}

type GetInfoRes struct {
	Inventory    []Inventory  `json:"inventory"`
	Coins        int64        `json:"coins"`
	CoinsHistory CoinsHistory `json:"coinsHistory"`
}
