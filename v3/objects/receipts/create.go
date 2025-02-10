package receipts

import "github.com/1N1Group/go-yookassa/v3/objects/common"

type CreateReceiptRequest struct {
	Type                      string
	PaymentId                 *string
	RefundId                  *string
	Customer                  common.Customer
	Items                     []common.Item
	Send                      bool
	TaxSystemCode             *int
	AdditionalUserProps       *common.AdditionalUserProps
	ReceiptIndustryDetails    []*common.IndustryDetail
	ReceiptOperationalDetails *common.OperationalDetails
	Settlements               []common.Settlement
	OnBehalfOf                *string
}
