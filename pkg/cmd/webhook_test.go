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
			t,
			"--api-key", "string",
			"webhooks", "create",
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
			t, pipeData,
			"--api-key", "string",
			"webhooks", "create",
		)
	})
}

func TestWebhooksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "update",
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
			t, pipeData,
			"--api-key", "string",
			"webhooks", "update",
			"--webhook-id", "webhookId",
		)
	})
}

func TestWebhooksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "list",
		)
	})
}

func TestWebhooksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "delete",
			"--webhook-id", "webhookId",
		)
	})
}

func TestWebhooksTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhooks", "test",
			"--webhook-id", "webhookId",
		)
	})
}
