package refunds

type GetListRefundRequest struct {
	CreatedAtGte *string
	CreatedAtGt  *string
	CreatedAtLte *string
	CreatedAtLt  *string
	Status       *string
	PaymentId    *string
	Limit        int
	Cursor       *string
}
