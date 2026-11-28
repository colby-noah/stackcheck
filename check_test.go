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
	tests := []struct {
		name   string
		status int
		wantUp bool
	}{
		{name: "ok", status: 200, wantUp: true},
		{name: "client error", status: 400, wantUp: false},
		{name: "server error", status: 500, wantUp: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := newTestServer(t, tt.status)

			target := Target{Name: "Test", URL: srv.URL}
			got := check(target)

			if got.Up != tt.wantUp {
				t.Errorf("check().Up = %v, want %v", got.Up, tt.wantUp)
			}

			if got.StatusCode != tt.status {
				t.Errorf("check().StatusCode = %d, want %d", got.StatusCode, tt.status)
			}
		})
	}
}

func TestCheckUnreachable(t *testing.T) {
	got := check(Target{Name: "Unreachable", URL: "https://127.0.0.1:1"})

	if got.Err == nil {
		t.Errorf("check().Err = nil, want an error")
	}

	if got.Up {
		t.Errorf("check().Up = true, want false")
	}
}

func newTestServer(t *testing.T, status int) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	}))
	t.Cleanup(srv.Close)
	return srv
}
