package payments

import (
	"time"

	"github.com/1N1Group/go-yookassa/v3/common"
)

type ReceiptRegistrationType string

const (
	ReceiptRegistrationPending   ReceiptRegistrationType = "pending"
	ReceiptRegistrationSucceeded ReceiptRegistrationType = "succeeded"
	ReceiptRegistrationCanceled  ReceiptRegistrationType = "canceled"
)

type Payment struct {
	ID                   string
	Status               PaymentStatus
	Paid                 bool
	Amount               common.Amount
	IncomeAmount         common.Amount
	AuthorizationDetails AuthorizationDetails
	Description          string
	Recipient            struct {
		AccountID string
		GatewayID string
	}
	Metadata            interface{}
	PaymentMethod       PaymentMethod
	CreatedAt           *time.Time
	CapturedAt          *time.Time
	ExpiresAt           *time.Time
	Confirmation        Confirmation
	Test                bool
	RefundedAmount      common.Amount
	ReceiptRegistration ReceiptRegistrationType
	CancellationDetails struct {
		Party  string
		Reason string
	}
	Transfers []struct {
		AccountID         string
		Amount            common.Amount
		Status            PaymentStatus
		PlatformFeeAmount common.Amount
		Description       string
		Metadata          interface{}
	}
	Deal               Deal
	MerchantCustomerId string
	InvoiceDetails     struct {
		ID string
	}
}
