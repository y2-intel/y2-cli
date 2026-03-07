// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
)

func TestSubscriptionsUpdateDelivery(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "subscriptions", "update-delivery",
			"--api-key", "string",
			"--subscription-id", "subscriptionId",
			"--delivery-method", "email",
			"--webhook-config-id", "webhookConfigId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"deliveryMethod: email\n" +
			"webhookConfigId: webhookConfigId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "subscriptions", "update-delivery",
			"--api-key", "string",
			"--subscription-id", "subscriptionId",
		)
	})
}
