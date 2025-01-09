package payments

import "github.com/1N1Group/go-yookassa/v3/common"

type SettlmentType string

const (
	SettlmentTypePayout SettlmentType = "payout"
)

type Settlment interface{}

type Payout struct {
	Type   SettlmentType
	Amount common.Amount
}

type Deal struct {
	ID          string
	Sattlements []Settlment
}
