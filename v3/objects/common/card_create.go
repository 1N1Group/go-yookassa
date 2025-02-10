package common

type CardCreate struct {
	Number      string
	ExpiryYear  string
	ExpiryMonth string
	CardHolder  *string
	CSC         *string
}
