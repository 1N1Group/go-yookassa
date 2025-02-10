package refunds

type RefundsListResponse struct {
	Type       string
	Items      []*RefundResponse
	NextCursor *string
}
