// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
)

func TestSubscriptionsUpdateDelivery(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"subscriptions", "update-delivery",
		"--api-key", "string",
		"--subscription-id", "subscriptionId",
		"--delivery-method", "email",
		"--webhook-config-id", "webhookConfigId",
	)
}
