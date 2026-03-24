// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
)

func TestOsintCountriesGetCountryInstabilityIndex(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint:countries", "get-country-instability-index",
			"--country-code", "UA",
		)
	})
}

func TestOsintCountriesGetCountryNews(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint:countries", "get-country-news",
			"--country-code", "US",
			"--limit", "1",
		)
	})
}

func TestOsintCountriesGetIntelligenceBrief(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint:countries", "get-intelligence-brief",
			"--country-code", "US",
		)
	})
}

func TestOsintCountriesGetPredictionMarkets(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint:countries", "get-prediction-markets",
			"--country-code", "US",
			"--limit", "1",
		)
	})
}

func TestOsintCountriesGetStockMarketIndex(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint:countries", "get-stock-market-index",
			"--country-code", "US",
		)
	})
}
