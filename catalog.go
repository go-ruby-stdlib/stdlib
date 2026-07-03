package stdlib

import "sort"

// Module describes one member of the go-ruby-* ecosystem: a pure-Go,
// MRI-faithful implementation of a single Ruby standard-library module.
//
// The organization, import path, and site URLs are not stored — they follow
// the family's uniform convention and are derived from Name by the methods
// below, so the catalog cannot drift out of sync with them.
type Module struct {
	// Name is the short module name, identical to the Ruby feature name where
	// it is a single word (for example "set", "json", "did-you-mean"). It is
	// both the org suffix and the repository/Go-module name.
	Name string
	// Ruby is the human-facing Ruby module, class, or feature the package
	// reimplements (for example "Set", "OpenStruct", "unicode_normalize").
	Ruby string
	// Wave is the ecosystem build wave the module shipped in (1 through 4).
	Wave int
	// Description is a one-line summary of what the module provides.
	Description string
}

// Org returns the GitHub organization that owns the module, for example
// "go-ruby-set".
func (m Module) Org() string { return "go-ruby-" + m.Name }

// ImportPath returns the Go module / import path, for example
// "github.com/go-ruby-set/set".
func (m Module) ImportPath() string { return "github.com/" + m.Org() + "/" + m.Name }

// RepoURL returns the canonical source-repository URL.
func (m Module) RepoURL() string { return "https://github.com/" + m.Org() + "/" + m.Name }

// LandingURL returns the module's landing-site URL.
func (m Module) LandingURL() string { return "https://" + m.Org() + ".github.io/" }

// DocsURL returns the module's documentation-site URL.
func (m Module) DocsURL() string { return "https://" + m.Org() + ".github.io/docs/" }

// catalog is the authoritative list of every module in the ecosystem, in the
// order the waves were built. Public accessors return sorted copies so callers
// can never mutate it.
var catalog = []Module{
	// Wave 1-2.
	{Name: "regexp", Ruby: "Regexp", Wave: 1, Description: "Onigmo-compatible regular-expression engine"},
	{Name: "erb", Ruby: "ERB", Wave: 1, Description: "ERB embedded-Ruby templating"},
	{Name: "yaml", Ruby: "YAML", Wave: 1, Description: "YAML emitter and loader (Psych)"},
	{Name: "format", Ruby: "format", Wave: 1, Description: "sprintf / format / String#% formatting"},
	{Name: "strscan", Ruby: "StringScanner", Wave: 1, Description: "StringScanner lexical scanning"},
	{Name: "optparse", Ruby: "OptionParser", Wave: 1, Description: "OptionParser command-line option parsing"},
	{Name: "json", Ruby: "JSON", Wave: 1, Description: "JSON generation and parsing"},
	{Name: "bigdecimal", Ruby: "BigDecimal", Wave: 1, Description: "arbitrary-precision decimal arithmetic"},
	{Name: "date", Ruby: "Date", Wave: 1, Description: "Date and DateTime"},
	{Name: "uri", Ruby: "URI", Wave: 1, Description: "URI parsing and manipulation"},
	{Name: "csv", Ruby: "CSV", Wave: 1, Description: "CSV reading and writing"},
	{Name: "shellwords", Ruby: "Shellwords", Wave: 1, Description: "POSIX shell-style word splitting and escaping"},
	{Name: "digest", Ruby: "Digest", Wave: 1, Description: "message digests (MD5/SHA1/SHA2/...)"},
	{Name: "marshal", Ruby: "Marshal", Wave: 1, Description: "Marshal binary object serialization"},

	// Wave 3.
	{Name: "set", Ruby: "Set", Wave: 3, Description: "unordered unique collection with full set algebra"},
	{Name: "time", Ruby: "Time", Wave: 3, Description: "Time extensions (parse/strptime/RFC formats)"},
	{Name: "getoptlong", Ruby: "GetoptLong", Wave: 3, Description: "GNU-style command-line option parsing"},
	{Name: "scanf", Ruby: "scanf", Wave: 3, Description: "formatted input scanning"},
	{Name: "stringio", Ruby: "StringIO", Wave: 3, Description: "in-memory string-backed IO"},
	{Name: "abbrev", Ruby: "Abbrev", Wave: 3, Description: "unambiguous abbreviation calculation"},
	{Name: "tsort", Ruby: "TSort", Wave: 3, Description: "topological sorting"},
	{Name: "prime", Ruby: "Prime", Wave: 3, Description: "prime generation and Baillie-PSW primality"},
	{Name: "cgi", Ruby: "CGI", Wave: 3, Description: "CGI escaping, cookies, and helpers"},
	{Name: "zlib", Ruby: "Zlib", Wave: 3, Description: "deflate/inflate/gzip plus CRC32/Adler32"},
	{Name: "did-you-mean", Ruby: "DidYouMean", Wave: 3, Description: "spelling suggestions for names"},
	{Name: "ipaddr", Ruby: "IPAddr", Wave: 3, Description: "IPv4/IPv6 address manipulation"},
	{Name: "pathname", Ruby: "Pathname", Wave: 3, Description: "filesystem-path objects"},
	{Name: "rational", Ruby: "Rational", Wave: 3, Description: "exact rational numbers"},
	{Name: "prettyprint", Ruby: "PrettyPrint", Wave: 3, Description: "Wadler pretty-printing engine"},
	{Name: "unicode-normalize", Ruby: "unicode_normalize", Wave: 3, Description: "Unicode normalization (NFC/NFD/NFKC/NFKD)"},
	{Name: "cmath", Ruby: "CMath", Wave: 3, Description: "complex-valued math functions"},
	{Name: "matrix", Ruby: "Matrix", Wave: 3, Description: "Matrix and Vector linear algebra"},
	{Name: "complex", Ruby: "Complex", Wave: 3, Description: "complex numbers"},
	{Name: "resolv", Ruby: "Resolv", Wave: 3, Description: "DNS resolution primitives"},
	{Name: "rexml", Ruby: "REXML", Wave: 3, Description: "XML DOM, parsing, serialization, XPath subset"},

	// Wave 4.
	{Name: "logger", Ruby: "Logger", Wave: 4, Description: "leveled logging"},
	{Name: "base64", Ruby: "Base64", Wave: 4, Description: "Base64 encoding and decoding (SIMD-accelerated)"},
	{Name: "securerandom", Ruby: "SecureRandom", Wave: 4, Description: "cryptographically-secure random values"},
	{Name: "ostruct", Ruby: "OpenStruct", Wave: 4, Description: "dynamic attribute objects"},
	{Name: "benchmark", Ruby: "Benchmark", Wave: 4, Description: "timing measurement"},
	{Name: "pstore", Ruby: "PStore", Wave: 4, Description: "transactional file-backed object store"},
	{Name: "observer", Ruby: "Observable", Wave: 4, Description: "publish/subscribe mixin"},
	{Name: "find", Ruby: "Find", Wave: 4, Description: "recursive directory traversal"},
}

// Count reports the number of modules in the ecosystem.
func Count() int { return len(catalog) }

// Modules returns every module in the ecosystem, sorted by Name.
func Modules() []Module {
	out := make([]Module, len(catalog))
	copy(out, catalog)
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// Lookup returns the module with the given Name.
func Lookup(name string) (Module, bool) {
	for _, m := range catalog {
		if m.Name == name {
			return m, true
		}
	}
	return Module{}, false
}

// ByWave returns the modules that shipped in the given wave, sorted by Name.
func ByWave(wave int) []Module {
	var out []Module
	for _, m := range catalog {
		if m.Wave == wave {
			out = append(out, m)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// Waves returns the distinct wave numbers present in the catalog, ascending.
func Waves() []int {
	seen := map[int]bool{}
	var out []int
	for _, m := range catalog {
		if !seen[m.Wave] {
			seen[m.Wave] = true
			out = append(out, m.Wave)
		}
	}
	sort.Ints(out)
	return out
}

// ImportPaths returns the Go import path of every module, sorted.
func ImportPaths() []string {
	out := make([]string, len(catalog))
	for i, m := range catalog {
		out[i] = m.ImportPath()
	}
	sort.Strings(out)
	return out
}
