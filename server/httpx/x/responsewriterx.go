/// copy from go-zero

package x

import (
	"bufio"
	"net"
	"net/http"
)

// A CodeResponseWriter is a helper to delay sealing a http.ResponseWriter on writing code.
type CodeResponseWriter struct {
	Writer http.ResponseWriter
	Code   int
}

// Flush flushes the response writer.
func (w *CodeResponseWriter) Flush() {
	if flusher, ok := w.Writer.(http.Flusher); ok {
		flusher.Flush()
	}
}

// Header returns the http header.
func (w *CodeResponseWriter) Header() http.Header {
	return w.Writer.Header()
}

// Hijack implements the http.Hijacker interface.
// This expands the Response to fulfill http.Hijacker if the underlying http.ResponseWriter supports it.
func (w *CodeResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.Writer.(http.Hijacker).Hijack()
}

// Write writes bytes into w.
func (w *CodeResponseWriter) Write(bytes []byte) (int, error) {
	return w.Writer.Write(bytes)
}

// WriteHeader writes code into w, and not sealing the writer.
func (w *CodeResponseWriter) WriteHeader(code int) {
	w.Writer.WriteHeader(code)
	w.Code = code
}
