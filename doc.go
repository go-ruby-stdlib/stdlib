// Package stdlib is the capstone of the go-ruby-* ecosystem: a single,
// pure-Go (CGO-free) aggregation of the whole MRI-compatible Ruby standard
// library, and the authoritative, machine-readable catalog of every module in
// the family.
//
// # The ecosystem
//
// Each Ruby standard-library module is implemented as its own independent,
// MRI-faithful, pure-Go module under its own github.com/go-ruby-<name>
// organization, at the import path github.com/go-ruby-<name>/<name>. Every one
// of them is CGO-free, keeps its behaviour byte-faithful to MRI, holds its
// coverage at 100%, and is green across the six 64-bit Go targets (amd64,
// arm64, riscv64, loong64, ppc64le, s390x). They are consumed individually by
// go-embedded-ruby (rbgo) — the pure-Go Ruby VM — which binds each as a native
// module.
//
// This package ties the family together in two complementary ways:
//
//   - The catalog. This package exposes a programmatic index of every module —
//     [Modules], [Lookup], [ByWave], [ImportPaths] — with each [Module]'s Ruby
//     name, wave, one-line description, and derived repository, landing, and
//     documentation URLs. Tooling (the ecosystem portal, CI conformance checks,
//     rbgo's binding generator) consumes this as the single source of truth.
//
//   - The aggregate. The sub-package [github.com/go-ruby-stdlib/stdlib/all]
//     blank-imports every module in the catalog, so a consumer gets the entire
//     pure-Go Ruby standard library behind a single
//
//     go get github.com/go-ruby-stdlib/stdlib
//
//     and a single import. Building that sub-package is also the family's
//     integration test: it proves that every module compiles together, at its
//     pinned pseudo-version, on every supported architecture.
//
// # Versions
//
// This module does not hard-code the version of any sibling — the pinned
// versions live in go.mod (the aggregate's dependency graph). When a consumer's
// binary has the sibling modules in its build graph (for example because it
// imports the all sub-package), [BuildVersion] reports the exact version each
// was pinned to at build time, read from the embedded build info.
package stdlib
