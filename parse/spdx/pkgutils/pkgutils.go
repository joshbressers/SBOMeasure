package pkgutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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

func LoadPackageJSON() []OnePackage {
	// Load the json file named test-output.json
	jsonFile, err := os.Open("test-output.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	theJSON, _ := ioutil.ReadAll(jsonFile)

	var p []OnePackage
	if err := json.Unmarshal([]byte(theJSON), &p); err != nil {
		panic(err)
	}

	return p
}
