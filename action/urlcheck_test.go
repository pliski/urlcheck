package action

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCheckSomeUrlSendsNoCacheHeaders exercises the interactive (TUI) request
// path end-to-end against a local server and asserts the no-cache headers
// actually arrive.
func TestCheckSomeUrlSendsNoCacheHeaders(t *testing.T) {
	var gotCacheControl, gotPragma string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotCacheControl = r.Header.Get("Cache-Control")
		gotPragma = r.Header.Get("Pragma")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	msg := CheckSomeUrl(srv.URL)()
	if _, ok := msg.(StatusMsg); !ok {
		t.Fatalf("expected StatusMsg, got %T (%v)", msg, msg)
	}
	if gotCacheControl != "no-cache" {
		t.Errorf("Cache-Control = %q, want %q", gotCacheControl, "no-cache")
	}
	if gotPragma != "no-cache" {
		t.Errorf("Pragma = %q, want %q", gotPragma, "no-cache")
	}
}
