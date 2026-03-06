// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
)

func TestOsintCountriesGetCountryInstabilityIndex(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"osint:countries", "get-country-instability-index",
		"--api-key", "string",
		"--country-code", "UA",
	)
}

func TestOsintCountriesGetCountryNews(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"osint:countries", "get-country-news",
		"--api-key", "string",
		"--country-code", "US",
		"--limit", "1",
	)
}

func TestOsintCountriesGetIntelligenceBrief(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"osint:countries", "get-intelligence-brief",
		"--api-key", "string",
		"--country-code", "US",
	)
}

func TestOsintCountriesGetPredictionMarkets(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"osint:countries", "get-prediction-markets",
		"--api-key", "string",
		"--country-code", "US",
		"--limit", "1",
	)
}

func TestOsintCountriesGetStockMarketIndex(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"osint:countries", "get-stock-market-index",
		"--api-key", "string",
		"--country-code", "US",
	)
}
