package payments

type AuthorizationDetails struct {
	RRN          string
	AuthCode     string
	ThreeDSecure struct {
		Applied bool
	}
}
