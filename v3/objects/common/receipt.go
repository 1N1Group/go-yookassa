package common

type Receipt struct {
	Customer                  *Customer
	Items                     []Item
	Phone                     *string
	Email                     *string
	TaxSystemCode             *int
	ReceiptIndustryDetails    []*IndustryDetail
	ReceiptOperationalDetails *OperationalDetails
}
