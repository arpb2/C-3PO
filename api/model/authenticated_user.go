package model

type AuthenticatedUser struct {
	*User
	Password string `json:"password"`
}
