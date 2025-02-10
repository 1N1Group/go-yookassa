package v3

import (
	"fmt"

	"github.com/1N1Group/go-yookassa/v3/objects/refunds"
	"resty.dev/v3"
)

const (
	baseRefundsUrl = "/payments"
)

type RefundsService interface {
	Create(req *refunds.CreateRefundRequest) (*refunds.RefundResponse, error)
	GetList(req *refunds.GetListRefundRequest) (*refunds.RefundsListResponse, error)
	Get(refundId string) (*refunds.RefundResponse, error)
}

type RefundsServiceImpl struct {
	client *Client
}

func (s *RefundsServiceImpl) Create(req *refunds.CreateRefundRequest) (*refunds.RefundResponse, error) {
	result, err := s.client.request(resty.MethodPost, baseRefundsUrl, nil, req)

	if err != nil {
		return nil, err
	}

	return result.(*refunds.RefundResponse), nil
}

func (s *RefundsServiceImpl) GetList(req *refunds.GetListRefundRequest) (*refunds.RefundsListResponse, error) {
	result, err := s.client.request(resty.MethodGet, baseRefundsUrl, req, nil)

	if err != nil {
		return nil, err
	}

	return result.(*refunds.RefundsListResponse), nil
}

func (s *RefundsServiceImpl) Get(refundId string) (*refunds.RefundResponse, error) {
	url := fmt.Sprintf("%s/%s", baseRefundsUrl, refundId)

	result, err := s.client.request(resty.MethodGet, url, nil, nil)

	if err != nil {
		return nil, err
	}

	return result.(*refunds.RefundResponse), nil
}
