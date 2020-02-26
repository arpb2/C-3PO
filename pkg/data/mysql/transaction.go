package mysql

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/third_party/ent"
	"golang.org/x/crypto/scrypt"
)

const (
	hashBytes = 64
)

// Rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%v: %v", err, rerr)
	}
	if err == nil {
		return http.CreateInternalError()
	}
	return err
}

// SaltHash hashes a content with a given salt
func SaltHash(content []byte, salt []byte) ([]byte, error) {
	return scrypt.Key(content, salt, 1<<14, 8, 1, hashBytes)
}
