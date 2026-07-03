<p align="center"><img src="https://raw.githubusercontent.com/go-ruby-stdlib/brand/main/social/go-ruby-stdlib-stdlib.png" alt="go-ruby-stdlib/stdlib" width="720"></p>

# stdlib — go-ruby-stdlib

[![Docs](https://img.shields.io/badge/docs-mkdocs--material-DC2626)](https://go-ruby-stdlib.github.io/docs/)
[![License](https://img.shields.io/badge/license-BSD--3--Clause-blue)](LICENSE)
[![Go](https://img.shields.io/badge/go-1.26.4%2B-00ADD8)](https://go.dev/dl/)
[![Coverage](https://img.shields.io/badge/coverage-100%25-1a7f37)](#tests--coverage)

**The capstone of the go-ruby-\* ecosystem: one pure-Go (no cgo) import for the
whole MRI-compatible Ruby standard library** — plus the authoritative,
machine-readable catalog of every module in the family.

Each Ruby standard-library module is implemented as its own independent,
MRI-faithful, pure-Go module under its own `github.com/go-ruby-<name>`
organization (for example [`go-ruby-set/set`](https://github.com/go-ruby-set/set),
[`go-ruby-json/json`](https://github.com/go-ruby-json/json),
[`go-ruby-matrix/matrix`](https://github.com/go-ruby-matrix/matrix)). Every one is
CGO-free, byte-faithful to MRI, held at 100% coverage, and green across the six
64-bit Go targets. They are consumed individually by
[go-embedded-ruby](https://github.com/go-embedded-ruby/ruby) (rbgo), the pure-Go
Ruby VM. This module ties them together.

## Two things in one module

**1. The catalog.** A programmatic index of every module in the ecosystem — used
by the [ecosystem portal](https://go-ruby-stdlib.github.io/), by CI conformance
checks, and by rbgo's binding tooling as the single source of truth:

```go
package main

import (
	"fmt"

	"github.com/go-ruby-stdlib/stdlib"
)

func main() {
	fmt.Println(stdlib.Count()) // 43

	for _, m := range stdlib.ByWave(4) {
		fmt.Printf("%-14s %-14s %s\n", m.Name, m.Ruby, m.ImportPath())
	}
	// base64        Base64         github.com/go-ruby-base64/base64
	// benchmark     Benchmark      github.com/go-ruby-benchmark/benchmark
	// ...

	m, _ := stdlib.Lookup("matrix")
	fmt.Println(m.Description) // Matrix and Vector linear algebra
	fmt.Println(m.DocsURL())   // https://go-ruby-matrix.github.io/docs/
}
```

**2. The aggregate.** The `stdlib/all` sub-package blank-imports every module in
the catalog, so a single dependency pulls the entire pure-Go Ruby standard
library into your build graph:

```go
import _ "github.com/go-ruby-stdlib/stdlib/all"
```

Building `stdlib/all` is also the family's cross-module integration proof: it
compiles every sibling, at its pinned pseudo-version, on every supported
architecture. When the siblings are in your build graph, `stdlib.BuildVersion`
reports the version each was pinned to at build time.

## Install

```sh
go get github.com/go-ruby-stdlib/stdlib
```

## The ecosystem

43 modules across four build waves. Each links to its repository, landing page,
and docs; all follow the uniform `github.com/go-ruby-<name>/<name>` convention.

### Wave 1–2 — core

| Module | Ruby | Description |
| --- | --- | --- |
| [regexp](https://github.com/go-ruby-regexp/regexp) | `Regexp` | Onigmo-compatible regular-expression engine |
| [erb](https://github.com/go-ruby-erb/erb) | `ERB` | ERB embedded-Ruby templating |
| [yaml](https://github.com/go-ruby-yaml/yaml) | `YAML` | YAML emitter and loader (Psych) |
| [format](https://github.com/go-ruby-format/format) | `format` | `sprintf` / `format` / `String#%` formatting |
| [strscan](https://github.com/go-ruby-strscan/strscan) | `StringScanner` | StringScanner lexical scanning |
| [optparse](https://github.com/go-ruby-optparse/optparse) | `OptionParser` | OptionParser command-line option parsing |
| [json](https://github.com/go-ruby-json/json) | `JSON` | JSON generation and parsing |
| [bigdecimal](https://github.com/go-ruby-bigdecimal/bigdecimal) | `BigDecimal` | arbitrary-precision decimal arithmetic |
| [date](https://github.com/go-ruby-date/date) | `Date` | Date and DateTime |
| [uri](https://github.com/go-ruby-uri/uri) | `URI` | URI parsing and manipulation |
| [csv](https://github.com/go-ruby-csv/csv) | `CSV` | CSV reading and writing |
| [shellwords](https://github.com/go-ruby-shellwords/shellwords) | `Shellwords` | POSIX shell-style word splitting and escaping |
| [digest](https://github.com/go-ruby-digest/digest) | `Digest` | message digests (MD5/SHA1/SHA2/…) |
| [marshal](https://github.com/go-ruby-marshal/marshal) | `Marshal` | Marshal binary object serialization |

### Wave 3

| Module | Ruby | Description |
| --- | --- | --- |
| [set](https://github.com/go-ruby-set/set) | `Set` | unordered unique collection with full set algebra |
| [time](https://github.com/go-ruby-time/time) | `Time` | Time extensions (parse/strptime/RFC formats) |
| [getoptlong](https://github.com/go-ruby-getoptlong/getoptlong) | `GetoptLong` | GNU-style command-line option parsing |
| [scanf](https://github.com/go-ruby-scanf/scanf) | `scanf` | formatted input scanning |
| [stringio](https://github.com/go-ruby-stringio/stringio) | `StringIO` | in-memory string-backed IO |
| [abbrev](https://github.com/go-ruby-abbrev/abbrev) | `Abbrev` | unambiguous abbreviation calculation |
| [tsort](https://github.com/go-ruby-tsort/tsort) | `TSort` | topological sorting |
| [prime](https://github.com/go-ruby-prime/prime) | `Prime` | prime generation and Baillie-PSW primality |
| [cgi](https://github.com/go-ruby-cgi/cgi) | `CGI` | CGI escaping, cookies, and helpers |
| [zlib](https://github.com/go-ruby-zlib/zlib) | `Zlib` | deflate/inflate/gzip plus CRC32/Adler32 |
| [did-you-mean](https://github.com/go-ruby-did-you-mean/did-you-mean) | `DidYouMean` | spelling suggestions for names |
| [ipaddr](https://github.com/go-ruby-ipaddr/ipaddr) | `IPAddr` | IPv4/IPv6 address manipulation |
| [pathname](https://github.com/go-ruby-pathname/pathname) | `Pathname` | filesystem-path objects |
| [rational](https://github.com/go-ruby-rational/rational) | `Rational` | exact rational numbers |
| [prettyprint](https://github.com/go-ruby-prettyprint/prettyprint) | `PrettyPrint` | Wadler pretty-printing engine |
| [unicode-normalize](https://github.com/go-ruby-unicode-normalize/unicode-normalize) | `unicode_normalize` | Unicode normalization (NFC/NFD/NFKC/NFKD) |
| [cmath](https://github.com/go-ruby-cmath/cmath) | `CMath` | complex-valued math functions |
| [matrix](https://github.com/go-ruby-matrix/matrix) | `Matrix` | Matrix and Vector linear algebra |
| [complex](https://github.com/go-ruby-complex/complex) | `Complex` | complex numbers |
| [resolv](https://github.com/go-ruby-resolv/resolv) | `Resolv` | DNS resolution primitives |
| [rexml](https://github.com/go-ruby-rexml/rexml) | `REXML` | XML DOM, parsing, serialization, XPath subset |

### Wave 4

| Module | Ruby | Description |
| --- | --- | --- |
| [logger](https://github.com/go-ruby-logger/logger) | `Logger` | leveled logging |
| [base64](https://github.com/go-ruby-base64/base64) | `Base64` | Base64 encoding and decoding (SIMD-accelerated) |
| [securerandom](https://github.com/go-ruby-securerandom/securerandom) | `SecureRandom` | cryptographically-secure random values |
| [ostruct](https://github.com/go-ruby-ostruct/ostruct) | `OpenStruct` | dynamic attribute objects |
| [benchmark](https://github.com/go-ruby-benchmark/benchmark) | `Benchmark` | timing measurement |
| [pstore](https://github.com/go-ruby-pstore/pstore) | `PStore` | transactional file-backed object store |
| [observer](https://github.com/go-ruby-observer/observer) | `Observable` | publish/subscribe mixin |
| [find](https://github.com/go-ruby-find/find) | `Find` | recursive directory traversal |

The [ecosystem portal](https://go-ruby-stdlib.github.io/) presents the same
catalog with live links to every module's landing and documentation sites.

## API

```go
type Module struct {
	Name        string // "set" — org suffix and module name
	Ruby        string // "Set" — the Ruby module/class/feature
	Wave        int    // build wave, 1–4
	Description string
}

func (m Module) Org() string        // "go-ruby-set"
func (m Module) ImportPath() string // "github.com/go-ruby-set/set"
func (m Module) RepoURL() string    // "https://github.com/go-ruby-set/set"
func (m Module) LandingURL() string // "https://go-ruby-set.github.io/"
func (m Module) DocsURL() string    // "https://go-ruby-set.github.io/docs/"

func Count() int
func Modules() []Module                    // all, sorted by Name
func Lookup(name string) (Module, bool)
func ByWave(wave int) []Module             // sorted by Name
func Waves() []int                         // distinct waves, ascending
func ImportPaths() []string                // all import paths, sorted

// Version of a sibling as pinned in the running binary's build graph.
func BuildVersion(importPath string) (version string, ok bool)
```

## Tests & coverage

The catalog and its accessors are tested to **100%** of statements, including
`BuildVersion`'s build-info paths. A dedicated test parses `all/all.go` and
asserts its blank-import set is exactly `ImportPaths()`, so the aggregate can
never drift out of sync with the catalog. The `all` sub-package is pure imports
(no statements); building it on every architecture is the ecosystem's
integration proof.

```sh
COVERPKG=$(go list ./... | paste -sd, -)
go test -race -coverpkg="$COVERPKG" -coverprofile=cover.out ./...
go tool cover -func=cover.out | tail -1   # 100.0%
```

CGO-free, `gofmt` + `go vet` clean, and green across the six 64-bit Go targets
(amd64, arm64, riscv64, loong64, ppc64le, s390x) and three OSes (Linux, macOS,
Windows).

## License

BSD-3-Clause — see [LICENSE](LICENSE). Copyright the go-ruby-stdlib/stdlib authors.
