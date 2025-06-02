package database

import (
	"hokusai/ent"
	// "hokusai/ent/intercept"

	"github.com/labstack/gommon/log"
)

func SetupHooks(dbConnection *ent.Client) {
	// dbConnection.Intercept(intercept.NewRelicSegmentDb())

	log.Info("initialized SetupHooks configuration=")
}
