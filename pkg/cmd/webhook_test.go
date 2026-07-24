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
			"--idempotency-key", "Idempotency-Key",
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
			"--idempotency-key", "Idempotency-Key",
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
			"--webhook-id", "whk_210b9798eb53baa4e69d31c1",
			"--name", "My Webhook",
			"--url", "https://example.com/webhook",
			"--headers", "{foo: string}",
			"--is-active=true",
			"--secret", "secret",
			"--if-match", "If-Match",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My Webhook\n" +
			"url: https://example.com/webhook\n" +
			"headers:\n" +
			"  foo: string\n" +
			"isActive: true\n" +
			"secret: secret\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhooks", "update",
			"--webhook-id", "whk_210b9798eb53baa4e69d31c1",
			"--if-match", "If-Match",
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
			"--webhook-id", "whk_210b9798eb53baa4e69d31c1",
			"--if-match", "If-Match",
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
			"--webhook-id", "whk_210b9798eb53baa4e69d31c1",
		)
	})
}
