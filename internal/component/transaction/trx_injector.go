//go:build wireinject
// +build wireinject

package transaction

import (
	"mceasy/ent"

	"github.com/google/wire"
)

var provider = wire.NewSet(
	NewTrx,
	wire.Bind(new(Trx), new(*TrxImpl)),
)

func InitializedTxService(dbClient *ent.Client) *TrxImpl {
	wire.Build(provider)
	return nil
}
