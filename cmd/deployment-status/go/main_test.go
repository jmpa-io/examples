package main

import (
	"encoding/json"
	"testing"
)

func TestValidateID(t *testing.T) {
	tests := []struct {
		id      string
		wantErr bool
	}{
		{"d-abc123", false},
		{"d-abc", false},
		{"d-123", false},
		{"d-", true},
		{"d-ABC", true},
		{"INVALID", true},
		{"x-abc123", true},
		{"", true},
	}
	for _, tt := range tests {
		err := validateID(tt.id)
		if (err != nil) != tt.wantErr {
			t.Errorf("validateID(%q) error = %v, wantErr %v", tt.id, err, tt.wantErr)
		}
	}
}

func TestDeploymentUnmarshal(t *testing.T) {
	t.Run("status present", func(t *testing.T) {
		var d Deployment
		if err := json.Unmarshal([]byte(`{"id":"d-abc","status":"active"}`), &d); err != nil {
			t.Fatal(err)
		}
		if d.Status == nil {
			t.Fatal("expected status to be non-nil")
		}
		if *d.Status != Active {
			t.Errorf("got %q, want %q", *d.Status, Active)
		}
	})

	t.Run("status absent", func(t *testing.T) {
		var d Deployment
		if err := json.Unmarshal([]byte(`{"id":"d-abc"}`), &d); err != nil {
			t.Fatal(err)
		}
		if d.Status != nil {
			t.Errorf("expected nil status, got %q", *d.Status)
		}
	})

	t.Run("status null", func(t *testing.T) {
		var d Deployment
		if err := json.Unmarshal([]byte(`{"id":"d-abc","status":null}`), &d); err != nil {
			t.Fatal(err)
		}
		if d.Status != nil {
			t.Errorf("expected nil status, got %q", *d.Status)
		}
	})
}
