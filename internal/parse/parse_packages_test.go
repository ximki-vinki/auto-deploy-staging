package parse

import (
	"testing"
)

func TestParsePackages_ReturnsSlice(t *testing.T) {
	yamlData := []byte(`{}`)

	packages, err := ParsePackages(yamlData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if packages == nil {
		t.Error("expected non-nil slice, got nil")
	}

}
