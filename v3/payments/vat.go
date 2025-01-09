package payments

import "github.com/1N1Group/go-yookassa/v3/common"

type VATType string

const (
	VATTypeUntaxed    VATType = "untaxed"
	VATTypeCalculated VATType = "calculated"
	VATTypeMixed      VATType = "mixed"
)

type VAT interface{}

type vat struct {
	Type VATType
}

type Untaxed struct {
	vat
}

type Calculated struct {
	vat
	Amount common.Amount
	Rate   string
}

type Mixed struct {
	vat
	Amount common.Amount
}
