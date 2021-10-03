package middleware1

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		logStartOfRequest(r)
		then := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Now().Sub(then)
		logEndOfRequest(duration)
	}
	return http.HandlerFunc(fn)
}

func logStartOfRequest(r *http.Request) {
	fields := logrus.Fields{
		"path":   r.Host + r.URL.String(),
		"method": r.Method,
	}

	logrus.WithFields(fields).Info("Starting request")
}

func logEndOfRequest(duration time.Duration) {
	fields := logrus.Fields{
		"status":  " mw.Status()",
		"bytes":    "mw.BytesWritten()",
		"duration": duration,
	}

	logrus.WithFields(fields).Info("Finished request")
}
