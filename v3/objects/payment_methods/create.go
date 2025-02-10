package paymentmethods

import "github.com/1N1Group/go-yookassa/v3/objects/common"

type CreatePaymentMethodRequest struct {
	Type         string
	Card         *common.CardCreate
	Holder       *common.HolderCreate
	ClientIp     *string
	Confirmation *interface{}
}
