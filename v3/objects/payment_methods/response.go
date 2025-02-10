package paymentmethods

import "github.com/1N1Group/go-yookassa/v3/objects/common"

type PaymentMethodResponse struct {
	Id           string
	Type         string
	Card         *common.Card
	Saved        bool
	Status       string
	Holder       common.Holder
	Title        *string
	Confirmation *interface{}
}
