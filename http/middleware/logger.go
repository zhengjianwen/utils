package middleware

import (
	"io"
	"net/http"
	"time"

	"github.com/zhengjianwen/utils/log"
)
var BaseDubug bool

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// Logger inherits from log.Logger used to log messages with the Logger middleware
	*log.Logger
}

// NewLogger returns a new Logger instance
func NewLogger(out io.Writer) *Logger {
	l := log.New(out, "", 0)
	l.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return &Logger{l}
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(rw, r)
	use := time.Since(start)
	if use > time.Second {
		if BaseDubug{
			l.Printf("[middleware logger]%s %s in %v", r.Method, r.URL.Path, use)
		}
	}
}
