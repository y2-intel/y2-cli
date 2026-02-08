// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
)

func TestReportsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports", "retrieve",
		"--report-id", "k57abc123def456",
	)
}

func TestReportsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports", "list",
		"--limit", "1",
		"--profile-id", "k57abc123def456",
	)
}

func TestReportsRetrieveAudio(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"reports", "retrieve-audio",
		"--report-id", "reportId",
		"--redirect=true",
	)
}
