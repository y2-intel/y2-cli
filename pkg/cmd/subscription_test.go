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
			t,
			"--api-key", "string",
			"subscriptions", "update-delivery",
			"--subscription-id", "sub_210b9798eb53baa4e69d31c1",
			"--delivery-method", "email",
			"--email-audience", "individual",
			"--webhook-config-id", "whk_210b9798eb53baa4e69d31c1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"deliveryMethod: email\n" +
			"emailAudience: individual\n" +
			"webhookConfigId: whk_210b9798eb53baa4e69d31c1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"subscriptions", "update-delivery",
			"--subscription-id", "sub_210b9798eb53baa4e69d31c1",
		)
	})
}
