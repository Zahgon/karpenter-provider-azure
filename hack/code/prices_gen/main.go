/*
Portions Copyright (c) Microsoft Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"flag"
	"log"
	"os"
)

var regions = []string{
	"australiacentral",
	"australiacentral2",
	"australiaeast",
	"australiasoutheast",
	"austriaeast",
	"brazilsouth",
	"brazilsoutheast",
	"canadacentral",
	"canadaeast",
	"centralindia",
	"centralus",
	"chilecentral",
	"eastasia",
	"eastus",
	"eastus2",
	"francecentral",
	"francesouth",
	"germanynorth",
	"germanywestcentral",
	"indonesiacentral",
	"israelcentral",
	"italynorth",
	"japaneast",
	"japanwest",
	"jioindiacentral",
	"jioindiawest",
	"koreacentral",
	"koreasouth",
	"malaysiawest",
	"mexicocentral",
	"newzealandnorth",
	"northcentralus",
	"northeurope",
	"norwayeast",
	"norwaywest",
	"polandcentral",
	"qatarcentral",
	"southafricanorth",
	"southafricawest",
	"southcentralus",
	"southeastasia",
	"southindia",
	"spaincentral",
	"swedencentral",
	"swedensouth",
	"switzerlandnorth",
	"switzerlandwest",
	"uaecentral",
	"uaenorth",
	"uksouth",
	"ukwest",
	"westcentralus",
	"westeurope",
	"westindia",
	"westus",
	"westus2",
	"westus3",
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("Usage: %s pkg/providers/pricing/zz_generated.pricing.go", os.Args[0])
	}

	generatePricing(flag.Arg(0))
}

func generatePricing(filePath string) { _ = "STUB: not implemented"; return }

// error handling omitted for example

// Pin cloud to public for now

// Fetch prices for each region concurrently

// Collect results

// Write output in deterministic order

func writePricing(src *bytes.Buffer, instanceNames []string, region string, getPrice func(instanceType string) (float64, bool)) {
	_ = "STUB: not implemented"
	return
}

// TODO: look at grouping by families to make the generated output nicer
