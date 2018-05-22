package middleware

import (
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/zhengjianwen/utils/http/render"
	"github.com/zhengjianwen/utils/log"
	"github.com/toolkits/web/errors"
)

// Recovery is a Negroni middleware that recovers from any panics and writes a 500 if there was one.
type Recovery struct {
	Logger     *log.Logger
	PrintStack bool
	StackAll   bool
	StackSize  int
}

// NewRecovery returns a new instance of Recovery
func NewRecovery(out io.Writer) *Recovery {
	return &Recovery{
		Logger:     log.New(out, "", 0),
		PrintStack: true,
		StackAll:   false,
		StackSize:  1024 * 8,
	}
}

func (rec *Recovery) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {

			if er, ok := err.(errors.Error); ok {
				render.ResponeErr(w, er.Msg)
				return
			}

			// Negroni part
			w.WriteHeader(http.StatusInternalServerError)
			stack := make([]byte, rec.StackSize)
			stack = stack[:runtime.Stack(stack, rec.StackAll)]

			f := "PANIC: %s\n%s"
			rec.Logger.Printf(f, err, stack)

			if rec.PrintStack {
				fmt.Fprintf(w, f, err, stack)
			}
		}
	}()

	next(w, r)
}
