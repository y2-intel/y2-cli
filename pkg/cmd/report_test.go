// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
)

func TestReportsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"reports", "retrieve",
			"--report-id", "rpt_0123456789abcdef01234567",
			"--format", "markdown",
			"--include", "include",
			"--view", "agent",
		)
	})
}

func TestReportsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"reports", "list",
			"--cursor", "cursor",
			"--format", "json",
			"--limit", "1",
			"--profile-id", "prf_0123456789abcdef01234567",
		)
	})
}

func TestReportsRetrieveAudio(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"reports", "retrieve-audio",
			"--report-id", "rpt_210b9798eb53baa4e69d31c1",
			"--redirect=true",
		)
	})
}
