package authenticated_user_task

import (
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

type FetchUserIdTask func(credentialService service.CredentialService, user *model.AuthenticatedUser) (uint, error)

func FetchUserIdTaskImpl(credentialService service.CredentialService, user *model.AuthenticatedUser) (uint, error) {
	userId, err := credentialService.Retrieve(user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
