/*
Package accesslogs provides a simple straight forward access log middleware. The logs are of the
following format:
<timestamp> <HTTP request method> <full URL including query string parameters> <duration of execution> <HTTP response status code>
*/
package accesslog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bnkamalesh/webgo/v4"
)

// AccessLog is a middleware which prints access log to stdout
func AccessLog(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(rw, req)
	end := time.Now()

	webgo.LOGHANDLER.Info(
		fmt.Sprintf(
			"%s %s %s %s %d",
			end.Format("2006-01-02 15:04:05 -0700 MST"),
			req.Method,
			req.URL.String(),
			end.Sub(start).String(),
			webgo.ResponseStatus(rw),
		),
	)
}
