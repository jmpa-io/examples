package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"regexp"
	"strings"
)

type Status string

const (
	Pending  Status = "pending"
	Active   Status = "active"
	Inactive Status = "inactive"
	Failed   Status = "failed"
)

type Deployment struct {
	ID     string  `json:"id"`
	Status *Status `json:"status"`
}

var idPattern = regexp.MustCompile(`^d-[a-z0-9]+$`)

func validateID(id string) error {
	if !strings.HasPrefix(id, "d-") {
		return fmt.Errorf("id must start with 'd-'")
	}
	if !idPattern.MatchString(id) {
		return fmt.Errorf("id must be lowercase alphanumeric after 'd-'")
	}
	return nil
}

func main() {
	examples := []string{
		`{"id":"d-abc123","status":"active"}`,
		`{"id":"d-xyz789"}`,
		`{"id":"INVALID","status":"pending"}`,
	}

	for _, raw := range examples {
		var d Deployment
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			slog.Error("unmarshal failed", "error", err)
			continue
		}

		if err := validateID(d.ID); err != nil {
			slog.Warn("invalid deployment", "id", d.ID, "error", err)
			continue
		}

		if d.Status == nil {
			slog.Info("deployment has no status", "id", d.ID)
		} else {
			slog.Info("deployment", "id", d.ID, "status", *d.Status)
		}
	}
}
