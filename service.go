package tipocambiobccr

// BCCRSvc BCCR Service struct
type BCCRSvc struct {
	email string
	token string
	name  string
}

// NewBCCRSvc Return a new BCCR Service
func NewBCCRSvc(email, token, name string) (BCCRSvc, error) {

	// TODO: Add Errors

	return BCCRSvc{
		email: email,
		token: token,
		name:  name,
	}, nil

}
