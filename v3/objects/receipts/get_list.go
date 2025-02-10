package receipts

type GetListReceiptRequest struct {
	CreatedAtGte *string
	CreatedAtGt  *string
	CreatedAtLte *string
	CreatedAtLt  *string
	Status       *string
	PaymentId    *string
	RefundId     *string
	Limit        int
	Cursor       *string
}
