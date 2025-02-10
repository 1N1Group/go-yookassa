package invoices

import "github.com/1N1Group/go-yookassa/v3/objects/common"

type CreateInvoiceRequest struct {
	PaymentData        interface{}
	DeliveryMethodData *interface{}
	Cart               []common.Cart
	ExpiresAt          string
	Locale             *string
	Description        *string
	Metadata           *interface{}
}
