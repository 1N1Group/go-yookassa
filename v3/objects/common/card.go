package common

type Card struct {
	First6        *string
	Last4         string
	ExpiryYear    string
	ExpiryMonth   string
	CardType      string
	CardProduct   *CardProduct
	IssuerCountry *string
	IssuerName    *string
	Source        *string
}
