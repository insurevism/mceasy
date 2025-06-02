package log

import "github.com/rs/zerolog"

type logWrapper struct {
	consoleWriter     zerolog.ConsoleWriter
	logger            zerolog.Logger
	appName           string
	level             zerolog.Level
	isDevelopmentMode bool
}

func (w *logWrapper) SetAppName(name string) *logWrapper {
	w.appName = name
	return w
}

func (w *logWrapper) SetIsDevelopmentMode(isDevelopmentMode bool) *logWrapper {
	w.isDevelopmentMode = isDevelopmentMode
	return w
}

func (w *logWrapper) SetLogLevel(level zerolog.Level) *logWrapper {
	w.level = level
	return w
}

func (w *logWrapper) SetConsoleWriter(consoleWriter zerolog.ConsoleWriter) *logWrapper {
	w.consoleWriter = consoleWriter
	return w
}

func (w *logWrapper) UpdateLogger(newLogger zerolog.Logger) *logWrapper {
	w.logger = newLogger
	return w
}

func (w *logWrapper) Apply() {
	w.logger = w.logger.Level(w.level)
	if w.isDevelopmentMode {
		w.logger = w.logger.Output(w.consoleWriter)
	}
}
