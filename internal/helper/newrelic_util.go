package helper

import (
	"context"
	"mceasy/middleware"

	"github.com/labstack/gommon/log"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func GetNewRelicTransaction(ctx context.Context) *newrelic.Transaction {

	if ctx != nil {
		if txn, isExist := ctx.Value("newrelic-transaction").(*newrelic.Transaction); isExist && txn != nil {
			return txn
		} else {
			return creatTxnNewRelic()
		}
	} else {
		return creatTxnNewRelic()
	}
}

func creatTxnNewRelic() *newrelic.Transaction {
	app, err := middleware.NewRelicApplication()
	if err != nil {
		log.Warnf("error initialized new relic configuration= %s", err)
	}

	txn := app.StartTransaction("newrelic-transaction")
	defer txn.End()
	return txn
}

func CreateNewRelicSegment(ctx context.Context, dataStore newrelic.DatastoreProduct, collection, operation string) *newrelic.DatastoreSegment {

	// Start a Datastore segment
	trx := GetNewRelicTransaction(ctx)

	// Start a Datastore segment
	datastore := newrelic.DatastoreSegment{
		StartTime:  trx.StartSegmentNow(),
		Product:    dataStore,
		Collection: collection,
		Operation:  operation,
	}
	defer datastore.End()

	return &datastore
}
