package common

type Cart struct {
	Description   string
	Price         Amount
	DiscountPrice *Amount
	Quantity      float64
}
