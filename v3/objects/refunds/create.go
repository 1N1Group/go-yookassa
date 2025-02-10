package refunds

import "github.com/1N1Group/go-yookassa/v3/objects/common"

type CreateRefundRequest struct {
	PaymentId        string
	Amount           common.Amount
	Description      *string
	Receipt          *common.Receipt
	Sources          []*common.Source
	Deal             *common.Deal
	RefundMethodData *interface{}
}
