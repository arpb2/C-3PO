package user

import (
	"context"

	"github.com/arpb2/C-3PO/third_party/ent"
	"github.com/arpb2/C-3PO/third_party/ent/credential"
	"github.com/arpb2/C-3PO/third_party/ent/user"
)

func del(dbClient *ent.Client, userId uint) error {
	ctx := context.Background()

	_, err := dbClient.Credential.Delete().Where(credential.HasHolderWith(user.ID(userId))).Exec(ctx)
	if err != nil {
		return err
	}

	return dbClient.User.DeleteOneID(userId).Exec(ctx)
}
