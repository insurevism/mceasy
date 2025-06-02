package log

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	wrapper *logWrapper
)

const (
	contextualFieldKey = "elbafContextualLogger"
)

func newLogger(w io.Writer, lvl zerolog.Level) zerolog.Logger {
	return zerolog.New(w).
		Level(lvl).
		With().
		Timestamp().
		Logger()
}

func init() {
	zerolog.ErrorStackMarshaler = zerolog.ErrorMarshalFunc
	zerolog.TimeFieldFormat = time.RFC3339Nano

	wrapper = &logWrapper{
		logger: newLogger(os.Stderr, zerolog.TraceLevel),
	}
}

func colorize(val any, color int) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", color, val)
}

func withTimestampFormatter() func(i interface{}) string {
	return func(i interface{}) string {
		t, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", i))

		return colorize(t.Format("15:04:05.000"), 0x5a)
	}
}

func withLevelFormatter(appName string) func(i interface{}) string {
	return func(i interface{}) string {
		levelStr := fmt.Sprintf("%s", i)

		parsedLevel, _ := zerolog.ParseLevel(levelStr)
		formattedLevel := zerolog.FormattedLevels[parsedLevel]
		levelColor := zerolog.LevelColors[parsedLevel]

		coloredAppName := colorize(appName, 0x24)

		return fmt.Sprintf("%s %s >", coloredAppName, colorize(formattedLevel, levelColor))
	}
}

func extractFields(ctx context.Context) map[string]any {
	if fields, ok := ctx.Value(contextualFieldKey).(map[string]any); ok {
		return fields
	}

	return nil
}

func ApplyConfig(logLevel zerolog.Level, isDevelopmentMode bool, appName string) {
	consoleWriter := zerolog.ConsoleWriter{
		Out:             os.Stdout,
		FormatTimestamp: withTimestampFormatter(),
		FormatLevel:     withLevelFormatter(appName),
	}

	wrapper.
		SetAppName(appName).
		SetIsDevelopmentMode(isDevelopmentMode).
		SetLogLevel(logLevel).
		SetConsoleWriter(consoleWriter).
		Apply()
}

func Trace() *zerolog.Event {
	return wrapper.logger.Trace()
}

func TraceCtx(ctx context.Context) *zerolog.Event {
	return wrapper.logger.Trace().Fields(extractFields(ctx))
}

func Debug() *zerolog.Event {
	return wrapper.logger.Debug()
}

func DebugCtx(ctx context.Context) *zerolog.Event {
	return wrapper.logger.Debug().Fields(extractFields(ctx))
}

func Info() *zerolog.Event {
	return wrapper.logger.Info()
}

func InfoCtx(ctx context.Context) *zerolog.Event {
	return wrapper.logger.Info().Fields(extractFields(ctx))
}

func Warn() *zerolog.Event {
	return wrapper.logger.Warn()
}

func WarnCtx(ctx context.Context) *zerolog.Event {
	return wrapper.logger.Warn().Fields(extractFields(ctx))
}

func Error() *zerolog.Event {
	return wrapper.logger.Error()
}

func ErrorCtx(ctx context.Context) *zerolog.Event {
	return wrapper.logger.Error().Fields(extractFields(ctx))
}

func Fatal() *zerolog.Event {
	return wrapper.logger.Fatal()
}

func FatalCtx(ctx context.Context) *zerolog.Event {
	return wrapper.logger.Fatal().Fields(extractFields(ctx))
}

func Panic() *zerolog.Event {
	return wrapper.logger.Panic()
}

func PanicCtx(ctx context.Context) *zerolog.Event {
	return wrapper.logger.Panic().Fields(extractFields(ctx))
}
