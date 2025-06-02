package transaction

import (
	"context"
	"errors"
	"mceasy/ent"
	"mceasy/ent/enttest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestTrxImpl_WithSuccessfulTx(t *testing.T) {

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

	// CreateTx a TrxImpl instance with the mock client.
	trxService := NewTrx(client)
	ctx := context.Background()

	// Test a successful transaction.
	err := trxService.WithTx(context.Background(), func(tx *ent.Tx) error {
		// Insert some data into the database.
		_, err := tx.SystemParameter.Create().
			SetKey("John").
			SetValue("john@doe.com").
			Save(ctx)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		t.Errorf("WithTx returned an unexpected error: %v", err)
	}

}

func TestTrxImpl_WithFailedTx(t *testing.T) {

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

	// CreateTx a TrxImpl instance with the mock client.
	trxService := NewTrx(client)
	ctx := context.Background()

	// Test a failed transaction.
	expectedErr := errors.New("an error occurred during the transaction")
	err := trxService.WithTx(context.Background(), func(tx *ent.Tx) error {
		_, err := tx.SystemParameter.Create().
			SetKey("John cena").
			SetValue("john@cena.com").
			Save(ctx)

		if err != nil {
			return err
		}
		return expectedErr
	})

	if err == nil {
		t.Error("WithTx did not return an expected error")
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("WithTx returned an unexpected error: %v", err)
	}
}
