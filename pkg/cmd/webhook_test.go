// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
)

func TestWebhooksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "create",
		"--api-key", "string",
		"--name", "My Webhook",
		"--url", "https://example.com/webhook",
		"--headers", "{foo: string}",
		"--secret", "secret",
	)
}

func TestWebhooksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "update",
		"--api-key", "string",
		"--webhook-id", "webhookId",
		"--headers", "{foo: string}",
		"--is-active=true",
		"--name", "name",
		"--secret", "secret",
		"--url", "https://example.com",
	)
}

func TestWebhooksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "list",
		"--api-key", "string",
	)
}

func TestWebhooksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "delete",
		"--api-key", "string",
		"--webhook-id", "webhookId",
	)
}

func TestWebhooksTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "test",
		"--api-key", "string",
		"--webhook-id", "webhookId",
	)
}
