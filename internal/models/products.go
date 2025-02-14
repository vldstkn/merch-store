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

func IsValid(item string) bool {
	switch item {
	case string(TSirt):
		return true
	case string(Cup):
		return true
	case string(Book):
		return true
	case string(Pen):
		return true
	case string(Powerbank):
		return true
	case string(Hoody):
		return true
	case string(Umbrella):
		return true
	case string(Socks):
		return true
	case string(Wallet):
		return true
	case string(PinkHoody):
		return true
	default:
		return false
	}
}

type Product struct {
	Id    int64       `db:"id"`
	Type  ProductType `db:"type"`
	Price int64       `db:"price"`
}

type Inventory struct {
	Type     ProductType `db:"type"`
	Quantity int64       `db:"quantity"`
}
