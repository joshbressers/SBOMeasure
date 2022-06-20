package pkgutils

// Let's make this not very smart to start. We'll just use an array of
// package-version as the output

type OnePackage struct {
	PackageName    string `json:"PackageName"`
	PackageVersion string `json:"PackageVersion"`
}

// Figure out if a package is in an array of packages
func Contains(packages []OnePackage, p OnePackage) bool {
	for _, v := range packages {
		if v == p {
			return true
		}
	}

	return false
}
