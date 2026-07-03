package stdlib

import (
	"runtime/debug"
	"testing"
)

func TestBuildVersionNoBuildInfo(t *testing.T) {
	orig := readBuildInfo
	t.Cleanup(func() { readBuildInfo = orig })
	readBuildInfo = func() (*debug.BuildInfo, bool) { return nil, false }
	if v, ok := BuildVersion("github.com/go-ruby-set/set"); ok || v != "" {
		t.Errorf("BuildVersion with no build info = (%q, %v), want (\"\", false)", v, ok)
	}
}

func TestBuildVersionFound(t *testing.T) {
	orig := readBuildInfo
	t.Cleanup(func() { readBuildInfo = orig })
	readBuildInfo = func() (*debug.BuildInfo, bool) {
		return &debug.BuildInfo{Deps: []*debug.Module{
			{Path: "github.com/go-ruby-set/set", Version: "v1.2.3"},
		}}, true
	}
	if v, ok := BuildVersion("github.com/go-ruby-set/set"); !ok || v != "v1.2.3" {
		t.Errorf("BuildVersion = (%q, %v), want (v1.2.3, true)", v, ok)
	}
	if v, ok := BuildVersion("github.com/go-ruby-json/json"); ok || v != "" {
		t.Errorf("BuildVersion(absent) = (%q, %v), want (\"\", false)", v, ok)
	}
}

func TestVersionOfReplace(t *testing.T) {
	info := &debug.BuildInfo{Deps: []*debug.Module{
		{Path: "github.com/go-ruby-a/a", Version: "v0.1.0"},
		{
			Path:    "github.com/go-ruby-b/b",
			Version: "v0.2.0",
			Replace: &debug.Module{Path: "example.com/fork", Version: "v9.9.9"},
		},
	}}
	if v, ok := versionOf(info, "github.com/go-ruby-a/a"); !ok || v != "v0.1.0" {
		t.Errorf("versionOf(a) = (%q, %v), want (v0.1.0, true)", v, ok)
	}
	if v, ok := versionOf(info, "github.com/go-ruby-b/b"); !ok || v != "v9.9.9" {
		t.Errorf("versionOf(b, replaced) = (%q, %v), want (v9.9.9, true)", v, ok)
	}
	if v, ok := versionOf(info, "github.com/missing/missing"); ok || v != "" {
		t.Errorf("versionOf(missing) = (%q, %v), want (\"\", false)", v, ok)
	}
}
