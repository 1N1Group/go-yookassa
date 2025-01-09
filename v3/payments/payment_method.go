package payments

import "github.com/1N1Group/go-yookassa/v3/common"

type PaymentMethodType string

const (
	PaymentTypeAlfabank              PaymentMethodType = "alfabank"
	PaymentTypeApplePay              PaymentMethodType = "apple_pay"
	PaymentTypeBankCard              PaymentMethodType = "bank_card"
	PaymentTypeB2BSberbank           PaymentMethodType = "b2b_sberbank"
	PaymentTypeCash                  PaymentMethodType = "cash"
	PaymentTypeElectronicCertificate PaymentMethodType = "electronic_certificate"
	PaymentTypeGooglePay             PaymentMethodType = "google_pay"
	PaymentTypeInstallments          PaymentMethodType = "installments"
	PaymentTypeMobileBalance         PaymentMethodType = "mobile_balance"
	PaymentTypeSberPay               PaymentMethodType = "sberbank"
	PaymentTypeSberbankLoan          PaymentMethodType = "sber_loan"
	PaymentTypeSBP                   PaymentMethodType = "sbp"
	PaymentTypeTinkoffBank           PaymentMethodType = "tinkoff_bank"
	PaymentTypeQiwi                  PaymentMethodType = "qiwi"
	PaymentTypeYooMoney              PaymentMethodType = "yoo_money"
	PaymentTypeWebmoney              PaymentMethodType = "webmoney"
	PaymentTypeWeChat                PaymentMethodType = "wechat"
)

type PaymentMethod interface{}

type paymentMethod struct {
	Type  PaymentMethodType
	ID    string
	Saved bool
	Title string
}

type SberbankLoan struct {
	paymentMethod
	DiscountAmount common.Amount
	LoanOption     string
}

type Alfabank struct {
	paymentMethod
	Login string
}

type MobileBalance struct {
	paymentMethod
}

type BankCard struct {
	paymentMethod
	Card Card
}

type Installments struct {
	paymentMethod
}

type Cash struct {
	paymentMethod
}

type SBP struct {
	paymentMethod
	PayerBankDetails struct {
		BankID string
		BIC    string
	}
	SBPOperationID string
}

type SberbankB2B struct {
	paymentMethod
	PayerBankDetails struct {
		FullName   string
		ShortName  string
		Address    string
		INN        string
		BankName   string
		BankBranch string
		BankBik    string
		Account    string
		KPP        string
	}
	PaymentPurpose string
	VatData        VAT
}

type ElectronicCertificate struct {
	paymentMethod
	Articles []struct {
		ArticleNumber uint
		TRUCode       string
		ArticleCode   string
		Certificates  []struct {
			CertificateID         string
			TRUQuantity           uint
			AvailableCompensation common.Amount
			AppliedCompensation   common.Amount
		}
	}
	Card                  Card
	ElectronicCertificate struct {
		Amount   common.Amount
		BasketID string
	}
}

type YooMoney struct {
	paymentMethod
	AccountNumber string
}

type ApplePay struct {
	paymentMethod
}

type GooglePay struct {
	paymentMethod
}

type Qiwi struct {
	paymentMethod
}

type SberPay struct {
	paymentMethod
	Card  Card
	Phone string
}

type TinkoffBank struct {
	paymentMethod
	Card Card
}

type WeChat struct {
	paymentMethod
}

type Webmoney struct {
	paymentMethod
}
