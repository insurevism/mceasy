package database

import (
	"mceasy/ent"
	// "mceasy/ent/intercept"

	"github.com/labstack/gommon/log"
)

func SetupHooks(dbConnection *ent.Client) {
	// dbConnection.Intercept(intercept.NewRelicSegmentDb())

	log.Info("initialized SetupHooks configuration=")
}
