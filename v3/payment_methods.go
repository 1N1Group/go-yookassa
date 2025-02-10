package v3

import (
	"fmt"

	paymentmethods "github.com/1N1Group/go-yookassa/v3/objects/payment_methods"
	"resty.dev/v3"
)

const (
	basePaymentMethodUrl = "/payment_methods"
)

type PaymentMethodsService interface {
	Create(req *paymentmethods.CreatePaymentMethodRequest) (*paymentmethods.PaymentMethodResponse, error)
	Get(methodId string) (*paymentmethods.PaymentMethodResponse, error)
}

type PaymentMethodsServiceImpl struct {
	client *Client
}

func (s *PaymentMethodsServiceImpl) Create(req *paymentmethods.CreatePaymentMethodRequest) (*paymentmethods.PaymentMethodResponse, error) {
	result, err := s.client.request(resty.MethodPost, basePaymentMethodUrl, nil, req)

	if err != nil {
		return nil, err
	}

	return result.(*paymentmethods.PaymentMethodResponse), nil
}

func (s *PaymentMethodsServiceImpl) Get(methodId string) (*paymentmethods.PaymentMethodResponse, error) {
	url := fmt.Sprintf("%s/%s", baseInvoiceUrl, methodId)

	result, err := s.client.request(resty.MethodGet, url, nil, nil)

	if err != nil {
		return nil, err
	}

	return result.(*paymentmethods.PaymentMethodResponse), nil
}
