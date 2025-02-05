package v3

import (
	"fmt"

	"github.com/1N1Group/go-yookassa/v3/objects/payments"
	"resty.dev/v3"
)

const (
	basePaymentUrl = "/payments"
)

type PaymentsService interface {
	Create(req *payments.CreatePaymentRequest) (*payments.PaymentResponse, error)
	GetList(req *payments.GetListPaymentRequest) (*payments.PaymentsListResponse, error)
	Get(paymentId string) (*payments.PaymentResponse, error)
	Capture(paymentId string, req *payments.CapturePaymentRequest) (*payments.PaymentResponse, error)
	Cancel(paymentId string, req *payments.CancelPaymentRequest) (*payments.PaymentResponse, error)
}

type PaymentsServiceImpl struct {
	client *Client
}

func (s *PaymentsServiceImpl) Create(req *payments.CreatePaymentRequest) (*payments.PaymentResponse, error) {
	result, err := s.client.request(resty.MethodPost, basePaymentUrl, nil, req)

	if err != nil {
		return nil, err
	}

	return result.(*payments.PaymentResponse), nil
}

func (s *PaymentsServiceImpl) GetList(req *payments.GetListPaymentRequest) (*payments.PaymentsListResponse, error) {
	result, err := s.client.request(resty.MethodGet, basePaymentUrl, req, nil)

	if err != nil {
		return nil, err
	}

	return result.(*payments.PaymentsListResponse), nil
}

func (s *PaymentsServiceImpl) Get(paymentId string) (*payments.PaymentResponse, error) {
	url := fmt.Sprintf("%s/%s", basePaymentUrl, paymentId)

	result, err := s.client.request(resty.MethodGet, url, nil, nil)

	if err != nil {
		return nil, err
	}

	return result.(*payments.PaymentResponse), nil
}

func (s *PaymentsServiceImpl) Capture(paymentId string, req *payments.CapturePaymentRequest) (*payments.PaymentResponse, error) {
	url := fmt.Sprintf("%s/%s/capture", basePaymentUrl, paymentId)

	result, err := s.client.request(resty.MethodPost, url, nil, req)

	if err != nil {
		return nil, err
	}

	return result.(*payments.PaymentResponse), nil
}

func (s *PaymentsServiceImpl) Cancel(paymentId string, req *payments.CancelPaymentRequest) (*payments.PaymentResponse, error) {
	url := fmt.Sprintf("%s/%s/cancel", basePaymentUrl, paymentId)

	result, err := s.client.request(resty.MethodPost, url, nil, req)

	if err != nil {
		return nil, err
	}

	return result.(*payments.PaymentResponse), nil
}
