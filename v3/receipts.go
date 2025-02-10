package v3

import (
	"fmt"

	"github.com/1N1Group/go-yookassa/v3/objects/receipts"
	"resty.dev/v3"
)

const (
	baseReceiptsUrl = "/payments"
)

type ReceiptsService interface {
	Create(req *receipts.CreateReceiptRequest) (*receipts.ReceiptResponse, error)
	GetList(req *receipts.GetListReceiptRequest) (*receipts.ReceiptsListResponse, error)
	Get(receptId string) (*receipts.ReceiptResponse, error)
}

type ReceiptsServiceImpl struct {
	client *Client
}

func (s *ReceiptsServiceImpl) Create(req *receipts.CreateReceiptRequest) (*receipts.ReceiptResponse, error) {
	result, err := s.client.request(resty.MethodPost, baseReceiptsUrl, nil, req)

	if err != nil {
		return nil, err
	}

	return result.(*receipts.ReceiptResponse), nil
}

func (s *ReceiptsServiceImpl) GetList(req *receipts.GetListReceiptRequest) (*receipts.ReceiptsListResponse, error) {
	result, err := s.client.request(resty.MethodGet, baseReceiptsUrl, req, nil)

	if err != nil {
		return nil, err
	}

	return result.(*receipts.ReceiptsListResponse), nil
}

func (s *ReceiptsServiceImpl) Get(receptId string) (*receipts.ReceiptResponse, error) {
	url := fmt.Sprintf("%s/%s", baseReceiptsUrl, receptId)

	result, err := s.client.request(resty.MethodGet, url, nil, nil)

	if err != nil {
		return nil, err
	}

	return result.(*receipts.ReceiptResponse), nil
}
