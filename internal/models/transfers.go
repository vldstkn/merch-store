package models

type Transfer struct {
	Id       int64  `json:"id"`
	FromUser string `json:"from_user" db:"from_user"`
	ToUser   string `json:"to_user" db:"to_user"`
	Amount   int64  `json:"amount" db:"amount"`
}
