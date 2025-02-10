package refunds

import "github.com/1N1Group/go-yookassa/v3/objects/common"

type RefundResponse struct {
	Id                  string
	PaymentId           string
	Status              string
	CancellationDetails *common.CancellationDetails
	ReceiptRegistration *string
	CreatedAt           string
	Amount              common.Amount
	Description         *string
	Sources             []*common.Source
	Deal                *common.Deal
	RefundMethod        *interface{}
}
