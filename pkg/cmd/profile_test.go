// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/y2-intel/y2-cli/internal/mocktest"
	"github.com/y2-intel/y2-cli/internal/requestflag"
)

func TestProfilesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "profiles", "create",
			"--api-key", "string",
			"--frequency", "daily",
			"--name", "Cybersecurity Weekly",
			"--schedule-time-of-day", "09:00",
			"--topic", "Cybersecurity threats, vulnerabilities, and defense strategies",
			"--audio-config", "{enabled: true, speed: 0, voiceId: voiceId}",
			"--bluf-structure", "blufStructure",
			"--branding-template-id", "brandingTemplateId",
			"--budget-config", "{alertThreshold: 0, maxCostPerReport: 0}",
			"--custom-prompt", "customPrompt",
			"--freshness-config", "{enabled: true, maxAgeMs: 0, preferRecentSources: true, recencyWeight: 0, validateLinks: true}",
			"--is-community=true",
			"--model-config", "{maxOutputTokens: 0, modelId: modelId, temperature: 0}",
			"--recursion-config", "{enabled: true, maxDepth: 1, strategy: breadth-first}",
			"--schedule-day-of-month", "1",
			"--schedule-day-of-week", "monday",
			"--search-config", "{excludeDomains: [string], includeDomains: [string], maxResults: 0, searchDepth: basic, timeRange: timeRange, topic: topic}",
			"--tag", "string",
			"--tool-config", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(profilesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "profiles", "create",
			"--api-key", "string",
			"--frequency", "daily",
			"--name", "Cybersecurity Weekly",
			"--schedule-time-of-day", "09:00",
			"--topic", "Cybersecurity threats, vulnerabilities, and defense strategies",
			"--audio-config.enabled=true",
			"--audio-config.speed", "0",
			"--audio-config.voice-id", "voiceId",
			"--bluf-structure", "blufStructure",
			"--branding-template-id", "brandingTemplateId",
			"--budget-config.alert-threshold", "0",
			"--budget-config.max-cost-per-report", "0",
			"--custom-prompt", "customPrompt",
			"--freshness-config.enabled=true",
			"--freshness-config.max-age-ms", "0",
			"--freshness-config.prefer-recent-sources=true",
			"--freshness-config.recency-weight", "0",
			"--freshness-config.validate-links=true",
			"--is-community=true",
			"--model-config.max-output-tokens", "0",
			"--model-config.model-id", "modelId",
			"--model-config.temperature", "0",
			"--recursion-config.enabled=true",
			"--recursion-config.max-depth", "1",
			"--recursion-config.strategy", "breadth-first",
			"--schedule-day-of-month", "1",
			"--schedule-day-of-week", "monday",
			"--search-config.exclude-domains", "[string]",
			"--search-config.include-domains", "[string]",
			"--search-config.max-results", "0",
			"--search-config.search-depth", "basic",
			"--search-config.time-range", "timeRange",
			"--search-config.topic", "topic",
			"--tag", "string",
			"--tool-config", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"frequency: daily\n" +
			"name: Cybersecurity Weekly\n" +
			"scheduleTimeOfDay: '09:00'\n" +
			"topic: Cybersecurity threats, vulnerabilities, and defense strategies\n" +
			"audioConfig:\n" +
			"  enabled: true\n" +
			"  speed: 0\n" +
			"  voiceId: voiceId\n" +
			"blufStructure: blufStructure\n" +
			"brandingTemplateId: brandingTemplateId\n" +
			"budgetConfig:\n" +
			"  alertThreshold: 0\n" +
			"  maxCostPerReport: 0\n" +
			"customPrompt: customPrompt\n" +
			"freshnessConfig:\n" +
			"  enabled: true\n" +
			"  maxAgeMs: 0\n" +
			"  preferRecentSources: true\n" +
			"  recencyWeight: 0\n" +
			"  validateLinks: true\n" +
			"isCommunity: true\n" +
			"modelConfig:\n" +
			"  maxOutputTokens: 0\n" +
			"  modelId: modelId\n" +
			"  temperature: 0\n" +
			"recursionConfig:\n" +
			"  enabled: true\n" +
			"  maxDepth: 1\n" +
			"  strategy: breadth-first\n" +
			"scheduleDayOfMonth: '1'\n" +
			"scheduleDayOfWeek: monday\n" +
			"searchConfig:\n" +
			"  excludeDomains:\n" +
			"    - string\n" +
			"  includeDomains:\n" +
			"    - string\n" +
			"  maxResults: 0\n" +
			"  searchDepth: basic\n" +
			"  timeRange: timeRange\n" +
			"  topic: topic\n" +
			"tags:\n" +
			"  - string\n" +
			"toolConfig: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "profiles", "create",
			"--api-key", "string",
		)
	})
}

func TestProfilesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "profiles", "update",
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(profilesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "profiles", "update",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"audioConfig:\n" +
			"  enabled: true\n" +
			"  speed: 0\n" +
			"  voiceId: voiceId\n" +
			"blufStructure: blufStructure\n" +
			"brandingTemplateId: brandingTemplateId\n" +
			"budgetConfig:\n" +
			"  alertThreshold: 0\n" +
			"  maxCostPerReport: 0\n" +
			"customPrompt: customPrompt\n" +
			"frequency: daily\n" +
			"freshnessConfig:\n" +
			"  enabled: true\n" +
			"  maxAgeMs: 0\n" +
			"  preferRecentSources: true\n" +
			"  recencyWeight: 0\n" +
			"  validateLinks: true\n" +
			"isCommunity: true\n" +
			"modelConfig:\n" +
			"  maxOutputTokens: 0\n" +
			"  modelId: modelId\n" +
			"  temperature: 0\n" +
			"name: name\n" +
			"recursionConfig:\n" +
			"  enabled: true\n" +
			"  maxDepth: 1\n" +
			"  strategy: breadth-first\n" +
			"scheduleDayOfMonth: scheduleDayOfMonth\n" +
			"scheduleDayOfWeek: scheduleDayOfWeek\n" +
			"scheduleTimeOfDay: '73:16'\n" +
			"searchConfig:\n" +
			"  excludeDomains:\n" +
			"    - string\n" +
			"  includeDomains:\n" +
			"    - string\n" +
			"  maxResults: 0\n" +
			"  searchDepth: basic\n" +
			"  timeRange: timeRange\n" +
			"  topic: topic\n" +
			"status: active\n" +
			"tags:\n" +
			"  - string\n" +
			"topic: topic\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "profiles", "update",
			"--api-key", "string",
			"--profile-id", "k57abc123def456",
		)
	})
}

func TestProfilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "profiles", "list",
			"--api-key", "string",
		)
	})
}

func TestProfilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "profiles", "delete",
			"--api-key", "string",
			"--profile-id", "k57abc123def456",
		)
	})
}

func TestProfilesPartialUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "profiles", "partial-update",
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(profilesPartialUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "profiles", "partial-update",
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"audioConfig:\n" +
			"  enabled: true\n" +
			"  speed: 0\n" +
			"  voiceId: voiceId\n" +
			"blufStructure: blufStructure\n" +
			"brandingTemplateId: brandingTemplateId\n" +
			"budgetConfig:\n" +
			"  alertThreshold: 0\n" +
			"  maxCostPerReport: 0\n" +
			"customPrompt: customPrompt\n" +
			"frequency: daily\n" +
			"freshnessConfig:\n" +
			"  enabled: true\n" +
			"  maxAgeMs: 0\n" +
			"  preferRecentSources: true\n" +
			"  recencyWeight: 0\n" +
			"  validateLinks: true\n" +
			"isCommunity: true\n" +
			"modelConfig:\n" +
			"  maxOutputTokens: 0\n" +
			"  modelId: modelId\n" +
			"  temperature: 0\n" +
			"name: name\n" +
			"recursionConfig:\n" +
			"  enabled: true\n" +
			"  maxDepth: 1\n" +
			"  strategy: breadth-first\n" +
			"scheduleDayOfMonth: scheduleDayOfMonth\n" +
			"scheduleDayOfWeek: scheduleDayOfWeek\n" +
			"scheduleTimeOfDay: '73:16'\n" +
			"searchConfig:\n" +
			"  excludeDomains:\n" +
			"    - string\n" +
			"  includeDomains:\n" +
			"    - string\n" +
			"  maxResults: 0\n" +
			"  searchDepth: basic\n" +
			"  timeRange: timeRange\n" +
			"  topic: topic\n" +
			"status: active\n" +
			"tags:\n" +
			"  - string\n" +
			"topic: topic\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "profiles", "partial-update",
			"--api-key", "string",
			"--profile-id", "k57abc123def456",
		)
	})
}
