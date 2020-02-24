package credential

import (
	"bytes"
	"context"

	"github.com/arpb2/C-3PO/pkg/data/mysql/service"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	credentialservice "github.com/arpb2/C-3PO/pkg/domain/service/credential"
	"github.com/arpb2/C-3PO/third_party/ent"
	"github.com/arpb2/C-3PO/third_party/ent/credential"
	"github.com/arpb2/C-3PO/third_party/ent/user"
)

func CreateService(dbClient *ent.Client) credentialservice.Service {
	return &credentialService{
		dbClient: dbClient,
	}
}

type credentialService struct {
	dbClient *ent.Client
}

func (c credentialService) GetUserId(email, password string) (uint, error) {
	ctx := context.Background()
	cred, err := c.dbClient.Credential.
		Query().
		WithHolder().
		Where(credential.HasHolderWith(user.Email(email))).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return uint(0), http.CreateUnauthorizedError()
		}
		return uint(0), err
	}

	if cred == nil {
		return uint(0), http.CreateInternalError()
	}

	inputPwHash, err := service.SaltHash([]byte(password), cred.Salt)

	if err != nil {
		return uint(0), err
	}

	if bytes.Equal(inputPwHash, cred.PasswordHash) {
		return cred.Edges.Holder.ID, nil
	}
	return uint(0), http.CreateUnauthorizedError()
}
