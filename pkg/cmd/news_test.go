// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
)

func TestNewsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"news", "list",
		"--limit", "1",
		"--topics", "crypto,ai_agents,bitcoin",
	)
}

func TestNewsGetRecaps(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"news", "get-recaps",
		"--timeframe", "12h",
		"--topics", "topics",
	)
}

func TestNewsListFeeds(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"news", "list-feeds",
	)
}
