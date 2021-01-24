package logrusmiddleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	var err error
	if err = next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()

	p := req.URL.Path
	if p == "" {
		p = "/"
	}

	var userID string
	useridIn := c.Get("userid")
	if useridIn != nil {
		var ok bool
		userID, ok = useridIn.(string)
		if !ok {
			userID = ""
		}
	}

	bytesIn := req.Header.Get(echo.HeaderContentLength)
	if bytesIn == "" {
		bytesIn = "0"
	}

	xff := req.Header.Get("X-Forwarded-For")
	if xff == "" {
		xff = c.RealIP()
	}

	entry := logrus.WithFields(logrus.Fields{
		"time_rfc3339":  time.Now().Format(time.RFC3339),
		"remoteIP":      xff,
		"remote_ip":     c.RealIP(),
		"userId":        userID,
		"host":          req.Host,
		"uri":           req.RequestURI,
		"method":        req.Method,
		"path":          p,
		"referer":       req.Referer(),
		"user_agent":    req.UserAgent(),
		"status":        res.Status,
		"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		"latency_human": stop.Sub(start).String(),
		"bytes_in":      bytesIn,
		"bytes_out":     strconv.FormatInt(res.Size, 10),
	})

	if err != nil {
		entry = entry.WithError(err)
	}

	entry.Info("Handled request")

	return nil
}

func logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return logrusMiddlewareHandler(c, next)
	}
}

// Hook returns an echo.MiddlewareFunc that logs desired information using the logrus StandardLogger
func Hook() echo.MiddlewareFunc {
	return logger
}
