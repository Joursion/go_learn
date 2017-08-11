package httplogger

import (
	"log"
	"time"
	"net/http"
)

type loggedRoundTripper struct {
	rt http.RoundTripper
	log HTTPLogger
}

type HTTPLogger interface {
	LogRequest(*http.Request)
	LogResponese(*http.Request, *http.Response, error, time.Duration)
}

func NewLoggedTransport(rt http.RoundTripper, log HTTPLogger) http.RoundTripper {
	return &loggedRoundTripper{rt: rt, log: log}
}

func DefaultLogger struct {
    
}

func (dl DefaultLogger) LogRequest(*http.Request){}

func (dl DefaultLogger) LogResponse(req *http.Request, res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	log.Printf("method:%s host:%s path:%s duration:%s err%q", req.Method, req.Host, req.URL.Path, duration, err.Error())
}

func (c *loggedRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	c.log.LogRequest(req)
	startTime := time.Now()
	res, err := c.rt.RoundTrip(req)
	duration := time.Since(startTime)
	c.log.LogResponese(req, res, err, duration)
	return res, err
}

var DefaultLoggedTransport = NewLoggedTransport(http.DefaultTransport, DefaultLogger{}iiiiiiiii)