package stdlib

import "runtime/debug"

// readBuildInfo is indirected so tests can exercise the missing-build-info path.
var readBuildInfo = debug.ReadBuildInfo

// BuildVersion returns the module version that importPath was pinned to in the
// build graph of the running binary, if that module is present.
//
// The version is only available when the sibling module is actually part of the
// binary's build graph — for example because the binary imports the all
// sub-package, or imports the sibling directly. When it is not present, or when
// the binary was built without module information, ok is false.
func BuildVersion(importPath string) (version string, ok bool) {
	info, present := readBuildInfo()
	if !present {
		return "", false
	}
	return versionOf(info, importPath)
}

// versionOf finds importPath among the build info's dependencies, honouring a
// replace directive if one is in effect.
func versionOf(info *debug.BuildInfo, importPath string) (string, bool) {
	for _, dep := range info.Deps {
		if dep.Path != importPath {
			continue
		}
		if dep.Replace != nil {
			return dep.Replace.Version, true
		}
		return dep.Version, true
	}
	return "", false
}
