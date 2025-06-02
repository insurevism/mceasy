package log

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type ContextualLoggerConfig struct {
	DomainName              string
	LogHeaders              []string
	LogRequestTime          bool
	LogMethod               bool
	LogURI                  bool
	LogPath                 bool
	LogRoutePath            bool
	LogReferer              bool
	LogUserAgent            bool
	LogRequestContentLength bool
}

func ContextualLoggerMiddleware(config ContextualLoggerConfig) echo.MiddlewareFunc {
	return config.ToMiddleware()
}

func (config ContextualLoggerConfig) ToMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			fields := map[string]any{}

			if config.DomainName != "" {
				fields["domain"] = config.DomainName
			}
			if config.LogRequestTime {
				fields["requestTime"] = time.Now()
			}
			if config.LogMethod {
				fields["method"] = req.Method
			}
			if config.LogURI {
				fields["uri"] = req.RequestURI
			}
			if config.LogPath {
				reqPath := req.URL.Path
				if reqPath == "" {
					reqPath = "/"
				}
				fields["path"] = reqPath
			}
			if config.LogRoutePath {
				fields["routePath"] = c.Path()
			}
			if config.LogReferer {
				fields["referer"] = req.Referer()
			}
			if config.LogUserAgent {
				fields["userAgent"] = req.UserAgent()
			}
			if config.LogRequestContentLength {
				if contentLen := req.Header.Get(echo.HeaderContentLength); contentLen != "" {
					fields["contentLength"] = contentLen
				}
			}
			for _, header := range config.LogHeaders {
				if values, ok := req.Header[header]; ok {
					fields[header] = values
				}
			}

			newContext := context.WithValue(
				req.Context(),
				contextualFieldKey,
				fields,
			)
			c.SetRequest(req.WithContext(newContext))

			return next(c)
		}
	}
}
