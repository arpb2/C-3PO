package model

type AuthenticatedUser struct {
	*User
	*Credential
}
