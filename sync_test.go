package stdlib

import (
	"go/parser"
	"go/token"
	"sort"
	"strconv"
	"testing"
)

// TestAllPackageInSync guards the invariant that the all sub-package
// blank-imports exactly the modules in the catalog — no more, no less. If a
// module is added to or removed from the catalog without updating all/all.go
// (or vice versa), this fails.
func TestAllPackageInSync(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "all/all.go", nil, parser.ImportsOnly)
	if err != nil {
		t.Fatalf("parsing all/all.go: %v", err)
	}
	var imported []string
	for _, spec := range f.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			t.Fatalf("unquoting import %s: %v", spec.Path.Value, err)
		}
		imported = append(imported, path)
	}
	sort.Strings(imported)

	want := ImportPaths()
	if len(imported) != len(want) {
		t.Fatalf("all/all.go imports %d modules, catalog has %d", len(imported), len(want))
	}
	for i := range want {
		if imported[i] != want[i] {
			t.Errorf("import[%d] = %q, catalog has %q", i, imported[i], want[i])
		}
	}
}
