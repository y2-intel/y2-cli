// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
)

func TestWebhooksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "create",
			"--api-key", "string",
			"--name", "My Webhook",
			"--url", "https://example.com/webhook",
			"--headers", "{foo: string}",
			"--secret", "secret",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My Webhook\n" +
			"url: https://example.com/webhook\n" +
			"headers:\n" +
			"  foo: string\n" +
			"secret: secret\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "webhooks", "create",
			"--api-key", "string",
		)
	})
}

func TestWebhooksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "update",
			"--api-key", "string",
			"--webhook-id", "webhookId",
			"--headers", "{foo: string}",
			"--is-active=true",
			"--name", "name",
			"--secret", "secret",
			"--url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"headers:\n" +
			"  foo: string\n" +
			"isActive: true\n" +
			"name: name\n" +
			"secret: secret\n" +
			"url: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "webhooks", "update",
			"--api-key", "string",
			"--webhook-id", "webhookId",
		)
	})
}

func TestWebhooksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "list",
			"--api-key", "string",
		)
	})
}

func TestWebhooksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "delete",
			"--api-key", "string",
			"--webhook-id", "webhookId",
		)
	})
}

func TestWebhooksTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "test",
			"--api-key", "string",
			"--webhook-id", "webhookId",
		)
	})
}
