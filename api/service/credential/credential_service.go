package credential

type Service interface {
	GetUserId(email, password string) (uint, error)

	StoreCredentials(email, password string, userId uint) error
}
