package common

type Item struct {
	Description                   string
	Amount                        Amount
	VatCode                       int
	Quantity                      float64
	Measure                       *string
	MarkQuantity                  *MarkQuantity
	PaymentSubject                *string
	PaymentMode                   *string
	CountryOfOriginCode           *string
	CustomsDeclarationNumber      *string
	Excise                        *string
	ProductCode                   *string
	MarkCodeInfo                  *MarkCodeInfo
	MarkMode                      *string
	PaymentSubjectIndustryDetails []*IndustryDetail
	AdditionalPaymentSubjectProps *string
	Supplier                      *Supplier
	AgentType                     *string
}
