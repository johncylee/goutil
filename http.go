package goutil

import (
	"net/http"
)

// HeadResponseWriter is a wrapper around a given http.ResponseWriter,
// but does not send any payload.
type HeadResponseWriter struct {
	ResponseWriter http.ResponseWriter
}

func (t HeadResponseWriter) Header() http.Header {
	return t.ResponseWriter.Header()
}

func (t HeadResponseWriter) Write(b []byte) (int, error) {
	_, err := t.ResponseWriter.Write(nil)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (t HeadResponseWriter) WriteHeader(code int) {
	t.ResponseWriter.WriteHeader(code)
}

type headHandler struct {
	Handler http.Handler
}

func (t headHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "HEAD" {
		r.Method = "GET"
		head := HeadResponseWriter{
			ResponseWriter: w,
		}
		t.Handler.ServeHTTP(head, r)
		return
	}
	t.Handler.ServeHTTP(w, r)
}

// HeadHandler wraps around an http.Handler to handle requests with
// method "HEAD" by changing it to "GET" without sending any payload.
func HeadHandler(h http.Handler) http.Handler {
	return headHandler{
		Handler: h,
	}
}
