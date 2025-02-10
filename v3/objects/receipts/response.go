package receipts

import "github.com/1N1Group/go-yookassa/v3/objects/common"

type ReceiptResponse struct {
	Id                        string
	Type                      string
	PaymentId                 *string
	RefundId                  *string
	Status                    string
	FiscalDocumentNumber      *string
	FiscalStorageNumber       *string
	FiscalAttribute           *string
	RegisteredAt              *string
	FiscalProviderId          *string
	Items                     []common.Item
	Settlements               []*common.Settlement
	OnBehalfOf                *string
	TaxSystemCode             *int
	ReceiptIndustryDetails    []*common.IndustryDetail
	ReceiptOperationalDetails *common.OperationalDetails
}
