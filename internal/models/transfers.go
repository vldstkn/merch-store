package models

type Transfer struct {
	Id       int64  `json:"id"`
	FromUser string `json:"from_user" db:"from_user"`
	ToUser   string `json:"to_user" db:"to_user"`
	Amount   int64  `json:"amount" db:"amount"`
}

type Received struct {
	FromUser string `db:"from_user_name" json:"fromUser"`
	Amount   int64  `db:"amount" json:"amount"`
}
type Sent struct {
	ToUser string `db:"to_user_name" json:"toUser"`
	Amount int64  `db:"amount" json:"amount"`
}

type History struct {
	Received []Received `json:"received"`
	Sent     []Sent     `json:"sent"`
}
