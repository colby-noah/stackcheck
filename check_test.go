package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	tests := []struct {
		name string
		up   bool
		want string
	}{
		{"up", true, "UP"},
		{"down", false, "DOWN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Result{Up: tt.up}

			got := r.Status()

			if got != tt.want {
				t.Errorf("Status() with Up=%v = %q, want %q", tt.up, got, tt.want)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	srv := newTestServer(t, 200)

	target := Target{Name: "Test", URL: srv.URL}
	got := check(target)

	wantUp := true
	wantStatus := 200

	if got.Up != wantUp {
		t.Errorf("check().Up = %v, want %v", got.Up, wantUp)
	}

	if got.StatusCode != wantStatus {
		t.Errorf("check().StatusCode = %d, want %d", got.StatusCode, wantStatus)
	}
}

func newTestServer(t *testing.T, status int) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	}))
	t.Cleanup(srv.Close) // closes automatically when the test ends, no defer needed
	return srv
}
