package stdlib

import (
	"sort"
	"strings"
	"testing"
)

func TestCatalogIntegrity(t *testing.T) {
	if Count() != len(catalog) {
		t.Fatalf("Count() = %d, want %d", Count(), len(catalog))
	}
	if Count() == 0 {
		t.Fatal("catalog is empty")
	}
	seen := map[string]bool{}
	for _, m := range catalog {
		switch {
		case m.Name == "":
			t.Errorf("module has empty Name: %+v", m)
		case m.Ruby == "":
			t.Errorf("module %q has empty Ruby name", m.Name)
		case m.Description == "":
			t.Errorf("module %q has empty Description", m.Name)
		case m.Wave < 1 || m.Wave > 4:
			t.Errorf("module %q has out-of-range Wave %d", m.Name, m.Wave)
		}
		if seen[m.Name] {
			t.Errorf("duplicate module Name %q", m.Name)
		}
		seen[m.Name] = true
	}
}

func TestModulesSortedCopy(t *testing.T) {
	got := Modules()
	if len(got) != Count() {
		t.Fatalf("Modules() len = %d, want %d", len(got), Count())
	}
	if !sort.SliceIsSorted(got, func(i, j int) bool { return got[i].Name < got[j].Name }) {
		t.Error("Modules() is not sorted by Name")
	}
	// Mutating the returned slice must not affect the catalog.
	got[0] = Module{Name: "tampered"}
	if _, ok := Lookup("tampered"); ok {
		t.Error("mutating Modules() result leaked into the catalog")
	}
}

func TestLookup(t *testing.T) {
	m, ok := Lookup("set")
	if !ok {
		t.Fatal("Lookup(set) not found")
	}
	if m.Ruby != "Set" || m.Wave != 3 {
		t.Errorf("Lookup(set) = %+v, unexpected", m)
	}
	if _, ok := Lookup("nonesuch"); ok {
		t.Error("Lookup(nonesuch) unexpectedly found")
	}
}

func TestByWaveAndWaves(t *testing.T) {
	waves := Waves()
	want := []int{1, 3, 4}
	if len(waves) != len(want) {
		t.Fatalf("Waves() = %v, want %v", waves, want)
	}
	for i := range want {
		if waves[i] != want[i] {
			t.Fatalf("Waves() = %v, want %v", waves, want)
		}
	}
	var total int
	for _, w := range waves {
		ms := ByWave(w)
		if len(ms) == 0 {
			t.Errorf("ByWave(%d) is empty", w)
		}
		if !sort.SliceIsSorted(ms, func(i, j int) bool { return ms[i].Name < ms[j].Name }) {
			t.Errorf("ByWave(%d) not sorted", w)
		}
		for _, m := range ms {
			if m.Wave != w {
				t.Errorf("ByWave(%d) returned wave-%d module %q", w, m.Wave, m.Name)
			}
		}
		total += len(ms)
	}
	if total != Count() {
		t.Errorf("waves partition sums to %d, want %d", total, Count())
	}
	if ms := ByWave(99); ms != nil {
		t.Errorf("ByWave(99) = %v, want nil", ms)
	}
}

func TestImportPathsUniqueSorted(t *testing.T) {
	paths := ImportPaths()
	if len(paths) != Count() {
		t.Fatalf("ImportPaths() len = %d, want %d", len(paths), Count())
	}
	if !sort.StringsAreSorted(paths) {
		t.Error("ImportPaths() not sorted")
	}
	seen := map[string]bool{}
	for _, p := range paths {
		if seen[p] {
			t.Errorf("duplicate import path %q", p)
		}
		seen[p] = true
		if !strings.HasPrefix(p, "github.com/go-ruby-") {
			t.Errorf("import path %q has unexpected prefix", p)
		}
	}
}

func TestModuleDerivedURLs(t *testing.T) {
	m := Module{Name: "did-you-mean"}
	cases := map[string]string{
		"Org":        m.Org(),
		"ImportPath": m.ImportPath(),
		"RepoURL":    m.RepoURL(),
		"LandingURL": m.LandingURL(),
		"DocsURL":    m.DocsURL(),
	}
	want := map[string]string{
		"Org":        "go-ruby-did-you-mean",
		"ImportPath": "github.com/go-ruby-did-you-mean/did-you-mean",
		"RepoURL":    "https://github.com/go-ruby-did-you-mean/did-you-mean",
		"LandingURL": "https://go-ruby-did-you-mean.github.io/",
		"DocsURL":    "https://go-ruby-did-you-mean.github.io/docs/",
	}
	for k, got := range cases {
		if got != want[k] {
			t.Errorf("%s() = %q, want %q", k, got, want[k])
		}
	}
}
