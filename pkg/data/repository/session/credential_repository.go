package session

type CredentialRepository interface {
	GetUserId(email, password string) (uint, error)
}
