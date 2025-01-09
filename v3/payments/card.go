package payments

type Card struct {
	First6      string
	Last4       string
	ExpiryYear  string
	ExpiryMonth string
	CardType    string
	CardProduct struct {
		Code string
		Name string
	}
	IssuerCountry string
	IssuerName    string
	Source        string
}
