package dto

type (
	AuthenticateParams struct {
		Username string
		Password string
	}
	Token struct {
		TokenString string
	}
)
