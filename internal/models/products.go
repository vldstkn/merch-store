package models

type ProductType string

const (
	TSirt     ProductType = "t-shirt"
	Cup       ProductType = "cup"
	Book      ProductType = "book"
	Pen       ProductType = "pen"
	Powerbank ProductType = "powerbank"
	Hoody     ProductType = "hoody"
	Umbrella  ProductType = "umbrella"
	Socks     ProductType = "socks"
	Wallet    ProductType = "wallet"
	PinkHoody ProductType = "pink-hoody"
)

type Product struct {
	Id    int64       `db:"id"`
	Type  ProductType `db:"type"`
	Price int64       `db:"price"`
}

type Inventory struct {
	Type     ProductType `db:"type"`
	Quantity int64       `db:"quantity"`
}
