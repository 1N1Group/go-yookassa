package receipts

type ReceiptsListResponse struct {
	Type       string
	Items      []*ReceiptResponse
	NextCursor *string
}
