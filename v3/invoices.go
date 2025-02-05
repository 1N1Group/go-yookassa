package v3

import (
	"fmt"

	"github.com/1N1Group/go-yookassa/v3/objects/invoices"
	"resty.dev/v3"
)

const (
	baseInvoiceUrl = "/invoices"
)

type InvoicesService interface {
	Create(req *invoices.CreateInvoiceRequest) (*invoices.InvoiceResponse, error)
	Get(invoiceId string) (*invoices.InvoiceResponse, error)
}

type InvoicesServiceImpl struct {
	client *Client
}

func (s *InvoicesServiceImpl) Create(req *invoices.CreateInvoiceRequest) (*invoices.InvoiceResponse, error) {
	result, err := s.client.request(resty.MethodPost, baseInvoiceUrl, nil, req)

	if err != nil {
		return nil, err
	}

	return result.(*invoices.InvoiceResponse), nil
}

func (s *InvoicesServiceImpl) Get(invoiceId string) (*invoices.InvoiceResponse, error) {
	url := fmt.Sprintf("%s/%s", baseInvoiceUrl, invoiceId)

	result, err := s.client.request(resty.MethodGet, url, nil, nil)

	if err != nil {
		return nil, err
	}

	return result.(*invoices.InvoiceResponse), nil
}
