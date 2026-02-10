// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
	"github.com/y2-intel/y2-cli/internal/requestflag"
)

func TestProfilesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "create",
		"--frequency", "daily",
		"--name", "Cybersecurity Weekly",
		"--schedule-time-of-day", "09:00",
		"--topic", "Cybersecurity threats, vulnerabilities, and defense strategies",
		"--bluf-structure", "blufStructure",
		"--custom-prompt", "customPrompt",
		"--is-community=true",
		"--recursion-config", "{enabled: true, maxDepth: 1, strategy: breadth-first}",
		"--schedule-day-of-month", "1",
		"--schedule-day-of-week", "monday",
		"--tag", "string",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(profilesCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "create",
		"--frequency", "daily",
		"--name", "Cybersecurity Weekly",
		"--schedule-time-of-day", "09:00",
		"--topic", "Cybersecurity threats, vulnerabilities, and defense strategies",
		"--bluf-structure", "blufStructure",
		"--custom-prompt", "customPrompt",
		"--is-community=true",
		"--recursion-config.enabled=true",
		"--recursion-config.max-depth", "1",
		"--recursion-config.strategy", "breadth-first",
		"--schedule-day-of-month", "1",
		"--schedule-day-of-week", "monday",
		"--tag", "string",
	)
}

func TestProfilesUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "update",
		"--profile-id", "k57abc123def456",
		"--audio-config", "{enabled: true, speed: 0, voiceId: voiceId}",
		"--bluf-structure", "blufStructure",
		"--branding-template-id", "brandingTemplateId",
		"--budget-config", "{alertThreshold: 0, maxCostPerReport: 0}",
		"--custom-prompt", "customPrompt",
		"--frequency", "daily",
		"--freshness-config", "{enabled: true, maxAgeMs: 0, preferRecentSources: true, recencyWeight: 0, validateLinks: true}",
		"--is-community=true",
		"--model-config", "{maxOutputTokens: 0, modelId: modelId, temperature: 0}",
		"--name", "name",
		"--recursion-config", "{enabled: true, maxDepth: 1, strategy: breadth-first}",
		"--schedule-day-of-month", "scheduleDayOfMonth",
		"--schedule-day-of-week", "scheduleDayOfWeek",
		"--schedule-time-of-day", "73:16",
		"--search-config", "{excludeDomains: [string], includeDomains: [string], maxResults: 0, searchDepth: basic, timeRange: timeRange, topic: topic}",
		"--status", "active",
		"--tag", "string",
		"--topic", "topic",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(profilesUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "update",
		"--profile-id", "k57abc123def456",
		"--audio-config.enabled=true",
		"--audio-config.speed", "0",
		"--audio-config.voice-id", "voiceId",
		"--bluf-structure", "blufStructure",
		"--branding-template-id", "brandingTemplateId",
		"--budget-config.alert-threshold", "0",
		"--budget-config.max-cost-per-report", "0",
		"--custom-prompt", "customPrompt",
		"--frequency", "daily",
		"--freshness-config.enabled=true",
		"--freshness-config.max-age-ms", "0",
		"--freshness-config.prefer-recent-sources=true",
		"--freshness-config.recency-weight", "0",
		"--freshness-config.validate-links=true",
		"--is-community=true",
		"--model-config.max-output-tokens", "0",
		"--model-config.model-id", "modelId",
		"--model-config.temperature", "0",
		"--name", "name",
		"--recursion-config.enabled=true",
		"--recursion-config.max-depth", "1",
		"--recursion-config.strategy", "breadth-first",
		"--schedule-day-of-month", "scheduleDayOfMonth",
		"--schedule-day-of-week", "scheduleDayOfWeek",
		"--schedule-time-of-day", "73:16",
		"--search-config.exclude-domains", "[string]",
		"--search-config.include-domains", "[string]",
		"--search-config.max-results", "0",
		"--search-config.search-depth", "basic",
		"--search-config.time-range", "timeRange",
		"--search-config.topic", "topic",
		"--status", "active",
		"--tag", "string",
		"--topic", "topic",
	)
}

func TestProfilesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "list",
	)
}

func TestProfilesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "delete",
		"--profile-id", "k57abc123def456",
	)
}

func TestProfilesPartialUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "partial-update",
		"--profile-id", "k57abc123def456",
		"--audio-config", "{enabled: true, speed: 0, voiceId: voiceId}",
		"--bluf-structure", "blufStructure",
		"--branding-template-id", "brandingTemplateId",
		"--budget-config", "{alertThreshold: 0, maxCostPerReport: 0}",
		"--custom-prompt", "customPrompt",
		"--frequency", "daily",
		"--freshness-config", "{enabled: true, maxAgeMs: 0, preferRecentSources: true, recencyWeight: 0, validateLinks: true}",
		"--is-community=true",
		"--model-config", "{maxOutputTokens: 0, modelId: modelId, temperature: 0}",
		"--name", "name",
		"--recursion-config", "{enabled: true, maxDepth: 1, strategy: breadth-first}",
		"--schedule-day-of-month", "scheduleDayOfMonth",
		"--schedule-day-of-week", "scheduleDayOfWeek",
		"--schedule-time-of-day", "73:16",
		"--search-config", "{excludeDomains: [string], includeDomains: [string], maxResults: 0, searchDepth: basic, timeRange: timeRange, topic: topic}",
		"--status", "active",
		"--tag", "string",
		"--topic", "topic",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(profilesPartialUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "partial-update",
		"--profile-id", "k57abc123def456",
		"--audio-config.enabled=true",
		"--audio-config.speed", "0",
		"--audio-config.voice-id", "voiceId",
		"--bluf-structure", "blufStructure",
		"--branding-template-id", "brandingTemplateId",
		"--budget-config.alert-threshold", "0",
		"--budget-config.max-cost-per-report", "0",
		"--custom-prompt", "customPrompt",
		"--frequency", "daily",
		"--freshness-config.enabled=true",
		"--freshness-config.max-age-ms", "0",
		"--freshness-config.prefer-recent-sources=true",
		"--freshness-config.recency-weight", "0",
		"--freshness-config.validate-links=true",
		"--is-community=true",
		"--model-config.max-output-tokens", "0",
		"--model-config.model-id", "modelId",
		"--model-config.temperature", "0",
		"--name", "name",
		"--recursion-config.enabled=true",
		"--recursion-config.max-depth", "1",
		"--recursion-config.strategy", "breadth-first",
		"--schedule-day-of-month", "scheduleDayOfMonth",
		"--schedule-day-of-week", "scheduleDayOfWeek",
		"--schedule-time-of-day", "73:16",
		"--search-config.exclude-domains", "[string]",
		"--search-config.include-domains", "[string]",
		"--search-config.max-results", "0",
		"--search-config.search-depth", "basic",
		"--search-config.time-range", "timeRange",
		"--search-config.topic", "topic",
		"--status", "active",
		"--tag", "string",
		"--topic", "topic",
	)
}
