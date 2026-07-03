// Package all blank-imports every module in the go-ruby-* ecosystem, so that a
// single
//
//	import _ "github.com/go-ruby-stdlib/stdlib/all"
//
// pulls the entire pure-Go, CGO-free, MRI-compatible Ruby standard library into
// a consumer's build graph behind one dependency.
//
// Importing this package links all of the sibling modules into the binary. A
// consumer that only needs a few modules should import those directly instead;
// this package exists for the "give me everything" case and, in CI, as the
// integration proof that every module in the catalog compiles together at its
// pinned pseudo-version on every supported architecture.
//
// The set of imports below is kept exactly in sync with
// [github.com/go-ruby-stdlib/stdlib.ImportPaths] by a test in the parent
// package; do not edit it by hand without updating the catalog.
package all

import (
	// Wave 1-2.
	_ "github.com/go-ruby-bigdecimal/bigdecimal"
	_ "github.com/go-ruby-csv/csv"
	_ "github.com/go-ruby-date/date"
	_ "github.com/go-ruby-digest/digest"
	_ "github.com/go-ruby-erb/erb"
	_ "github.com/go-ruby-format/format"
	_ "github.com/go-ruby-json/json"
	_ "github.com/go-ruby-marshal/marshal"
	_ "github.com/go-ruby-optparse/optparse"
	_ "github.com/go-ruby-regexp/regexp"
	_ "github.com/go-ruby-shellwords/shellwords"
	_ "github.com/go-ruby-strscan/strscan"
	_ "github.com/go-ruby-uri/uri"
	_ "github.com/go-ruby-yaml/yaml"

	// Wave 3.
	_ "github.com/go-ruby-abbrev/abbrev"
	_ "github.com/go-ruby-cgi/cgi"
	_ "github.com/go-ruby-cmath/cmath"
	_ "github.com/go-ruby-complex/complex"
	_ "github.com/go-ruby-did-you-mean/did-you-mean"
	_ "github.com/go-ruby-getoptlong/getoptlong"
	_ "github.com/go-ruby-ipaddr/ipaddr"
	_ "github.com/go-ruby-matrix/matrix"
	_ "github.com/go-ruby-pathname/pathname"
	_ "github.com/go-ruby-prettyprint/prettyprint"
	_ "github.com/go-ruby-prime/prime"
	_ "github.com/go-ruby-rational/rational"
	_ "github.com/go-ruby-resolv/resolv"
	_ "github.com/go-ruby-rexml/rexml"
	_ "github.com/go-ruby-scanf/scanf"
	_ "github.com/go-ruby-set/set"
	_ "github.com/go-ruby-stringio/stringio"
	_ "github.com/go-ruby-time/time"
	_ "github.com/go-ruby-tsort/tsort"
	_ "github.com/go-ruby-unicode-normalize/unicode-normalize"
	_ "github.com/go-ruby-zlib/zlib"

	// Wave 4.
	_ "github.com/go-ruby-base64/base64"
	_ "github.com/go-ruby-benchmark/benchmark"
	_ "github.com/go-ruby-find/find"
	_ "github.com/go-ruby-logger/logger"
	_ "github.com/go-ruby-observer/observer"
	_ "github.com/go-ruby-ostruct/ostruct"
	_ "github.com/go-ruby-pstore/pstore"
	_ "github.com/go-ruby-securerandom/securerandom"
)
