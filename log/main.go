package main 

//reference "http://blog.csdn.net/wangshubo1989/article/details/70668916"

import (
	"time"
	"os"
	"log"
	"net/http"
	"go_learn/log/httplogger"
)

func main() {
	client := http.Client{
		Transport: httplogger.NewLoggedTransport(http.DefaultTransport, newLogger())
	}

	client.Get("https://google.com")
}

type httpLogger struct {
	log *log.Logger
}

func newLogger() *httpLogger {
	return &httpLogger{
		log: log.New(os.Stderr, " log - ", log.LstdFlags)
	}
}

func (l *httpLogger) LogRequest(req *http.Request) {
	l.log.Printf(
		"Request %s %s",
		req.Method,
		req.URL.String()
	)
}

func (l *httpLogger) LogResponse(req *http.Request, res *http.Response, err error , duration time.Duration) {
	duration /= time.Microsecond
	if err != nil {
		l.log.Println(err)
	} else {
		l.log.Printf(
			"Response method=%s status=%s durationMs=%d %s",
			req.Method,
			res.StatusCode,
			duration,
			req.URL.String()
		)
	}
}