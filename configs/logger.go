package configs

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	logger "mceasy/internal/component/log"
)

func SetupLogger(e *echo.Echo) {
	levelFromEnv := viper.GetString("application.log.level")
	level := toLogLevel(levelFromEnv)

	e.Logger.SetLevel(level)
	log.SetLevel(level)
}

func SetupZeroLogger() {
	levelFromEnv := viper.GetString("application.log.level")
	applicationName := viper.GetString("application.name")
	applicationMode := viper.GetString("application.mode")

	logger.ApplyConfig(
		toZeroLogLevel(levelFromEnv),
		applicationMode == "dev",
		applicationName,
	)
}

func toZeroLogLevel(level string) zerolog.Level {
	levelMap := map[string]zerolog.Level{
		"trace":    zerolog.TraceLevel,
		"debug":    zerolog.DebugLevel,
		"info":     zerolog.InfoLevel,
		"warn":     zerolog.WarnLevel,
		"error":    zerolog.ErrorLevel,
		"fatal":    zerolog.FatalLevel,
		"panic":    zerolog.PanicLevel,
		"disabled": zerolog.Disabled,
	}

	zerologLevel, ok := levelMap[strings.ToLower(level)]
	if !ok {
		return zerolog.InfoLevel
	}

	return zerologLevel
}

func toLogLevel(level string) log.Lvl {
	switch strings.ToLower(level) {
	case "debug":
		return log.DEBUG
	case "info":
		return log.INFO
	case "warn":
		return log.WARN
	case "error":
		return log.ERROR
	default:
		return log.INFO
	}
}
