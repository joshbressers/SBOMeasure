// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

// Example for: *jsonparser2v2*

// This example demonstrates loading an SPDX JSON document from disk into memory,
// and then logging some of the attributes to the console.
// Run project: go run example_json_loader.go ../sample-docs/json/SPDXJSONExample-v2.2.spdx.json
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spdx/tools-golang/jsonloader"
)

// Let's make this not very smart to start. We'll just use an array of
// package-version as the output

type OnePackage struct {
	PackageName    string `json:"PackageName"`
	PackageVersion string `json:"PackageVersion"`
}

// Figure out if a package is in an array of packages
func contains(packages []OnePackage, p OnePackage) bool {
	for _, v := range packages {
		if v == p {
			return true
		}
	}

	return false
}

func load_test_json() []OnePackage {
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

func main() {

	// check that we've received the right number of arguments
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %v <json-file-in>\n", args[0])
		fmt.Printf("  Load SPDX 2.2 JSON file <spdx-file-in>, and\n")
		fmt.Printf("  print portions of its creation info data.\n")
		return
	}

	// open the SPDX file
	fileIn := args[1]
	r, err := os.Open(fileIn)
	if err != nil {
		fmt.Printf("Error while opening %v for reading: %v", fileIn, err)
		return
	}
	defer r.Close()

	// try to load the SPDX file's contents as a json file, version 2.2
	doc, err := jsonloader.Load2_2(r)
	if err != nil {
		fmt.Printf("Error while parsing %v: %v", args[1], err)
		return
	}

	// if we got here, the file is now loaded into memory.
	fmt.Printf("Successfully loaded %s\n", args[1])

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Some Attributes of the Document:")
	fmt.Printf("Document Name:         %s\n", doc.CreationInfo.DocumentName)
	fmt.Printf("DataLicense:           %s\n", doc.CreationInfo.DataLicense)
	fmt.Printf("Document Namespace:    %s\n", doc.CreationInfo.DocumentNamespace)
	fmt.Printf("SPDX Version:          %s\n", doc.CreationInfo.SPDXVersion)
	fmt.Println(strings.Repeat("=", 80))

	spdxPackages := make([]OnePackage, 0)
	for _, i := range doc.Packages {
		onePackage := OnePackage{i.PackageName, i.PackageVersion}
		spdxPackages = append(spdxPackages, onePackage)
	}

	testPackages := load_test_json()

	fmt.Println("\nSPDX Packages")
	for _, i := range spdxPackages {
		fmt.Printf("%s-%s\n", i.PackageName, i.PackageVersion)
	}

	fmt.Println("\nJSON Packages")

	for _, i := range testPackages {
		fmt.Printf("%s-%s\n", i.PackageName, i.PackageVersion)
		if contains(spdxPackages, i) {
			fmt.Println("-- Found")
		} else {
			fmt.Println("-- NOT Found")
		}
	}

}
