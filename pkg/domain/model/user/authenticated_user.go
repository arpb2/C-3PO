package user

type AuthenticatedUser struct {
	User
	Password string `json:"password"`
}
