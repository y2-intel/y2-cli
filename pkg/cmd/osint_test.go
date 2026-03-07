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
			t, "osint", "get-conflict-indicators",
			"--api-key", "string",
			"--category", "seismic",
			"--limit", "1",
			"--region", "mena",
		)
	})
}

func TestOsintGetGpsJammingZones(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "osint", "get-gps-jamming-zones",
			"--api-key", "string",
			"--limit", "1",
			"--severity", "low",
		)
	})
}

func TestOsintGetMilitaryPosture(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "osint", "get-military-posture",
			"--api-key", "string",
			"--limit", "1",
		)
	})
}

func TestOsintListAircraft(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "osint", "list-aircraft",
			"--api-key", "string",
			"--limit", "1",
			"--theater", "theater",
		)
	})
}

func TestOsintListEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "osint", "list-events",
			"--api-key", "string",
			"--category", "seismic",
			"--limit", "1",
			"--severity", "low",
		)
	})
}

func TestOsintListVessels(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "osint", "list-vessels",
			"--api-key", "string",
			"--limit", "1",
			"--region", "region",
		)
	})
}

func TestOsintMapEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "osint", "map-events",
			"--api-key", "string",
			"--limit", "1",
			"--region", "mena",
		)
	})
}
