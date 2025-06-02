package transaction

import (
	"context"
	"hokusai/ent"
)

type Trx interface {
	WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error
}
