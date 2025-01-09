package payments

type ConfirmationType string

const (
	ConfirmationTypeEmbedded          ConfirmationType = "embedded"
	ConfirmationTypeExternal          ConfirmationType = "external"
	ConfirmationTypeMobileApplication ConfirmationType = "mobile_application"
	ConfirmationTypeQR                ConfirmationType = "qr"
	ConfirmationTypeRedirect          ConfirmationType = "redirect"
)

type Confirmation interface{}

type Embedded struct {
	Type              ConfirmationType
	ConfirmationToken string
}

type External struct {
	Type ConfirmationType
}

type MobileApplication struct {
	Type            ConfirmationType
	ConfirmationURL string
}

type QR struct {
	Type             ConfirmationType
	ConfirmationData string
}

type Redirect struct {
	Type            ConfirmationType
	ConfirmationURL string
	Enforce         bool
	ReturnURL       string
}
