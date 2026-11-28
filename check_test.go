package main

import "testing"

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
