package proxy

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// ProxyHandler
// This function returns a http handler  func that can be used in almost all of the we frameworks
func ProxyHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		val := ctx.Value("url")
		v, ok := val.(*url.URL)
		if !ok {
			return
		}
		// Get the params with gorilla mux
		// Because the http handlers are managed in the gorilla/mux router
		r := mux.Vars(req)
		// Reading required value of the map
		path := r["rest"]

		// Constructing  a new url according to
		// the incoming request and the target url
		req.URL = &url.URL{
			Scheme:      v.Scheme,
			Host:        v.Host,
			Path:        "/" + path,
			RawPath:     v.RawPath,
			RawQuery:    v.RawQuery,
			Fragment:    v.Fragment,
			RawFragment: v.RawFragment,
		}

		resp, err := http.DefaultTransport.RoundTrip(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		// Copy the incoming headers to the outgoing headers
		copyHeader(w.Header(), resp.Header)

		// Set http status code
		w.WriteHeader(resp.StatusCode)

		// copy the incoming response body to the outgoing response writer
		_, _ = io.Copy(w, resp.Body)

		// Logging the progress
		defer func() {
			logrus.StandardLogger().
				WithField("request_body", req.Body).
				WithField("request_url", req.URL.String()).
				WithField("request_headers", req.Header).
				WithField("method", req.Method).
				WithField("response", resp.Body).
				Infoln()
			_ = resp.Body.Close()
		}()
	}
}

// Copy header of src into the destination
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
