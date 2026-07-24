// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
)

func TestOsintGetConflictIndicators(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint", "get-conflict-indicators",
			"--category", "seismic",
			"--cursor", "cursor",
			"--format", "json",
			"--limit", "1",
			"--region", "mena",
		)
	})
}

func TestOsintGetGpsJammingZones(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint", "get-gps-jamming-zones",
			"--cursor", "cursor",
			"--format", "json",
			"--limit", "1",
			"--severity", "low",
		)
	})
}

func TestOsintGetMilitaryPosture(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint", "get-military-posture",
			"--cursor", "cursor",
			"--format", "json",
			"--limit", "1",
		)
	})
}

func TestOsintListAircraft(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint", "list-aircraft",
			"--cursor", "cursor",
			"--format", "json",
			"--limit", "1",
			"--theater", "theater",
		)
	})
}

func TestOsintListEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint", "list-events",
			"--category", "seismic",
			"--cursor", "cursor",
			"--format", "json",
			"--limit", "1",
			"--severity", "low",
		)
	})
}

func TestOsintListVessels(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint", "list-vessels",
			"--cursor", "cursor",
			"--format", "json",
			"--limit", "1",
			"--region", "region",
		)
	})
}

func TestOsintMapEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"osint", "map-events",
			"--cursor", "cursor",
			"--format", "json",
			"--limit", "1",
			"--region", "mena",
		)
	})
}
