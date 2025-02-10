package invoices

import "github.com/1N1Group/go-yookassa/v3/objects/common"

type InvoiceResponse struct {
	Id                  string
	Status              string
	Cart                []common.Cart
	DeliveryMethod      *interface{}
	PaymentDetails      *common.PaymentDetails
	CreatedAt           string
	ExpiresAt           *string
	Description         *string
	CancellationDetails *common.CancellationDetails
	Metadata            *interface{}
}
