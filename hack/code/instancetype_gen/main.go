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
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"github.com/Azure/karpenter-provider-azure/pkg/providers/instancetype"
	"github.com/alecthomas/kong"
	"github.com/samber/lo"
)

var args struct {
	Format         string `kong:"required,enum='testfakes,nameonly',help='Output format: testfakes (full SKU details) or nameonly (SKU names only).'"`
	Path           string `kong:"required,help='Output file path.'"`
	Location       string `kong:"optional,help='Azure region/location (required for testfakes, not allowed for nameonly).'"`
	Sizes          string `kong:"optional,help='Comma-separated list of VM sizes to fetch (testfakes only; if omitted, all sizes are included).'"`
	IgnoreFamilies string `kong:"optional,name='ignore-families',help='Comma-separated family:date pairs to ignore (nameonly only). SKUs in the given family are excluded until the date. Example: standardDasv7Family:2026-06-01,standardEasv6Family:2026-07-01'"`
}

func main() {
	kong.Parse(&args,
		kong.Name("instancetype-testdata-gen"),
		kong.Description("Generate instance type test data from Azure SKUs."),
	)

	if args.Format == "nameonly" && args.Sizes != "" {
		panic("--sizes cannot be specified with nameonly format")
	}
	if args.Format == "nameonly" && args.Location != "" {
		panic("--location cannot be specified with nameonly format (all regions are queried)")
	}
	if args.Format == "testfakes" && args.Location == "" {
		panic("--location is required for testfakes format")
	}
	if args.Format == "testfakes" && args.IgnoreFamilies != "" {
		panic("--ignore-families cannot be specified with testfakes format")
	}

	ignored := parseIgnoreFamilies(args.IgnoreFamilies)

	fmt.Println("starting generation of sku data...")
	sub := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if sub == "" {
		panic("AZURE_SUBSCRIPTION_ID env var is required")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(fmt.Sprintf("failed to create credential: %v", err))
	}

	client, err := armcompute.NewResourceSKUsClient(sub, cred, nil)
	if err != nil {
		panic(fmt.Sprintf("failed to create client: %v", err))
	}

	ctx := context.Background()

	switch args.Format {
	case "testfakes":
		pager := client.NewListPager(&armcompute.ResourceSKUsClientListOptions{
			Filter: lo.ToPtr(fmt.Sprintf("location eq '%s'", args.Location)),
		})
		var targetSkus map[string]struct{}
		if args.Sizes != "" {
			targetSkus = map[string]struct{}{}
			for _, s := range strings.Split(args.Sizes, ",") {
				targetSkus[s] = struct{}{}
			}
		}
		skuData := []*armcompute.ResourceSKU{}
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				panic(fmt.Sprintf("failed to get next page: %v", err))
			}
			for _, sku := range page.Value {
				if targetSkus != nil {
					if _, ok := targetSkus[*sku.Name]; !ok {
						continue
					}
				}
				skuData = append(skuData, sku)
			}
		}
		fmt.Println("Successfully Fetched all the SKUs", len(skuData))
		writeTestFakes(skuData, args.Location, args.Path)

	case "nameonly":
		pager := client.NewListPager(&armcompute.ResourceSKUsClientListOptions{})
		skuMap := map[string]*armcompute.ResourceSKU{}
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				panic(fmt.Sprintf("failed to get next page: %v", err))
			}
			for _, sku := range page.Value {
				if sku.ResourceType != nil && *sku.ResourceType == "virtualMachines" && sku.Name != nil {
					name := *sku.Name
					if _, exists := skuMap[name]; !exists {
						skuMap[name] = sku
					}
				}
			}
		}
		// Sort by name
		names := lo.Keys(skuMap)
		sort.Strings(names)

		sortedSKUs := lo.Map(names, func(name string, _ int) *armcompute.ResourceSKU {
			return skuMap[name]
		})
		fmt.Println("Successfully Fetched VM SKU names:", len(sortedSKUs))
		writeNameOnly(sortedSKUs, args.Path, ignored)
	}
	fmt.Println("Successfully Generated output at", args.Path)
}

func readExistingSKUs(path string) map[string]instancetype.SKUEntry {
	_ = "STUB: not implemented"
	return nil
}

// File doesn't exist yet; that's fine

// parseIgnoreFamilies parses a comma-separated string of "family:date" pairs.
// SKUs in the given family are excluded until the specified date (inclusive).
func parseIgnoreFamilies(raw string) map[string]time.Time { _ = "STUB: not implemented"; return nil }

func writeNameOnly(skus []*armcompute.ResourceSKU, path string, ignoredFamilies map[string]time.Time) {
	_ = "STUB: not implemented"
	return
}

func writeTestFakes(ResourceSkus []*armcompute.ResourceSKU, location, path string) {
	_ = "STUB: not implemented"
	return
}
