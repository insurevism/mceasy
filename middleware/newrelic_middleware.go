package middleware

import (
	"context"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/spf13/viper"
)

type Singleton struct {
	Value *newrelic.Transaction
}

var (
	instanceNRApp *newrelic.Application
	instance      *Singleton
	once          sync.Once
	obj           Singleton
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Value: obj.Value,
		}
	})
	return instance
}

func NewRelicConfig(e *echo.Echo) {
	app, err := NewRelicApplication()
	if err != nil {
		log.Warnf("error initialized new relic configuration=", err)
	}

	e.Use(newRelicMiddleware(app))
}

// NewRelicApplication doc new relic:
// https://docs.newrelic.com/docs/apm/agents/go-agent/configuration/go-agent-code-level-metrics-config/
// NewRelicApplication menginisialisasi dan mengembalikan instance singleton dari newrelic.Application.
func NewRelicApplication() (*newrelic.Application, error) {
	var err error
	once.Do(func() {
		instanceNRApp, err = newrelic.NewApplication(
			newrelic.ConfigAppName(viper.GetString("newrelic.name")),
			newrelic.ConfigLicense(viper.GetString("newrelic.key")),
			newrelic.ConfigCodeLevelMetricsEnabled(true),
			newrelic.ConfigCustomInsightsEventsEnabled(true),
		)
		if err != nil {
			log.Errorf("error initialized new relic configuration: %v\n", err)
		}
	})
	return instanceNRApp, err
}

func newRelicMiddleware(app *newrelic.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			txn := app.StartTransaction(c.Request().URL.Path)
			if txn != nil {
				obj = Singleton{
					Value: txn,
				}

				//set context for new relic transaction:
				requestCtx := c.Request().Context()
				requestCtx = context.WithValue(requestCtx, "newrelic-transaction", txn)
				c.SetRequest(c.Request().WithContext(requestCtx))
			}
			defer txn.End()

			err := next(c)
			if err != nil {
				txn.NoticeError(err)
				log.Debugf("catch exception in new relic middleware: %s", err)
			}
			return err
		}
	}
}
