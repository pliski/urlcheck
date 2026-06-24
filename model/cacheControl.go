package model

import "net/http"

// SetNoCacheHeaders asks intermediary caches and the origin not to serve a
// stale, stored response. These are request directives and only advisory:
// well-behaved caches revalidate against the origin, but some CDNs ignore
// client cache-control. "Pragma" is the HTTP/1.0 fallback.
func SetNoCacheHeaders(h http.Header) {
	h.Set("Cache-Control", "no-cache")
	h.Set("Pragma", "no-cache")
}
