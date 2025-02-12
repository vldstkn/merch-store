package models

type User struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Balance  int64  `db:"balance"`
}
