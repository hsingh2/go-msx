package config

import (
	"context"
	"testing"
)

func TestCachedLoader(t *testing.T) {
	static := NewStatic("static", map[string]string{"enabled": "true"})
	loader := NewCachedLoader(static)

	settings, err := loader.Load(context.Background())
	if err != nil {
		t.Error(err)
	}

	if settings["enabled"] != "true" {
		t.Errorf("Enabled was '%s', expected 'true'", settings["enabled"])
	}

	static.Set("enabled", "false")

	settings, err = loader.Load(context.Background())
	if err != nil {
		t.Error(err)
	}

	if settings["enabled"] != "true" {
		t.Errorf("Enabled was '%s', expected 'true'", settings["enabled"])
	}

	loader.Invalidate()
	static.Set("enabled", "false")

	settings, err = loader.Load(context.Background())
	if err != nil {
		t.Error(err)
	}

	if settings["enabled"] != "false" {
		t.Errorf("Enabled was '%s', expected 'false'", settings["enabled"])
	}
}