// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
	"github.com/y2-intel/y2-cli/internal/apiquery"
	"github.com/y2-intel/y2-cli/internal/requestflag"
	"github.com/y2-intel/y2-go"
	"github.com/y2-intel/y2-go/option"
)

var profilesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new intelligence profile with the specified configuration. The profile\nwill be owned by the authenticated user and start with `active` status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "frequency",
			Usage:    "Report generation frequency",
			Required: true,
			BodyPath: "frequency",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Profile display name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-time-of-day",
			Usage:    "Time of day for report generation (HH:mm, UTC)",
			Required: true,
			BodyPath: "scheduleTimeOfDay",
		},
		&requestflag.Flag[string]{
			Name:     "topic",
			Usage:    "Topic description for research",
			Required: true,
			BodyPath: "topic",
		},
		&requestflag.Flag[string]{
			Name:     "bluf-structure",
			Usage:    "Custom BLUF report structure template",
			BodyPath: "blufStructure",
		},
		&requestflag.Flag[string]{
			Name:     "custom-prompt",
			Usage:    "Custom system prompt for the AI analyst",
			BodyPath: "customPrompt",
		},
		&requestflag.Flag[bool]{
			Name:     "is-community",
			Usage:    "Whether this is a community (public) profile",
			Default:  false,
			BodyPath: "isCommunity",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "recursion-config",
			BodyPath: "recursionConfig",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-day-of-month",
			Usage:    "Day of month for monthly profiles",
			BodyPath: "scheduleDayOfMonth",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-day-of-week",
			Usage:    "Day of week for weekly/biweekly profiles",
			BodyPath: "scheduleDayOfWeek",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags for categorization",
			BodyPath: "tags",
		},
	},
	Action:          handleProfilesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"recursion-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "recursion-config.enabled",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "recursion-config.max-depth",
			InnerField: "maxDepth",
		},
		&requestflag.InnerFlag[string]{
			Name:       "recursion-config.strategy",
			InnerField: "strategy",
		},
	},
})

var profilesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Replaces all mutable fields of an existing intelligence profile. Only profiles\nowned by the authenticated user can be updated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "profile-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "audio-config",
			BodyPath: "audioConfig",
		},
		&requestflag.Flag[string]{
			Name:     "bluf-structure",
			BodyPath: "blufStructure",
		},
		&requestflag.Flag[string]{
			Name:     "branding-template-id",
			Usage:    "Branding template ID (Pro feature)",
			BodyPath: "brandingTemplateId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "budget-config",
			BodyPath: "budgetConfig",
		},
		&requestflag.Flag[string]{
			Name:     "custom-prompt",
			BodyPath: "customPrompt",
		},
		&requestflag.Flag[string]{
			Name:     "frequency",
			Usage:    "Report generation frequency",
			BodyPath: "frequency",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "freshness-config",
			BodyPath: "freshnessConfig",
		},
		&requestflag.Flag[bool]{
			Name:     "is-community",
			BodyPath: "isCommunity",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "model-config",
			BodyPath: "modelConfig",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "recursion-config",
			BodyPath: "recursionConfig",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-day-of-month",
			BodyPath: "scheduleDayOfMonth",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-day-of-week",
			BodyPath: "scheduleDayOfWeek",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-time-of-day",
			BodyPath: "scheduleTimeOfDay",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "search-config",
			BodyPath: "searchConfig",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Profile status",
			BodyPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "topic",
			BodyPath: "topic",
		},
	},
	Action:          handleProfilesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"audio-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "audio-config.enabled",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "audio-config.speed",
			InnerField: "speed",
		},
		&requestflag.InnerFlag[string]{
			Name:       "audio-config.voice-id",
			InnerField: "voiceId",
		},
	},
	"budget-config": {
		&requestflag.InnerFlag[float64]{
			Name:       "budget-config.alert-threshold",
			InnerField: "alertThreshold",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "budget-config.max-cost-per-report",
			InnerField: "maxCostPerReport",
		},
	},
	"freshness-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "freshness-config.enabled",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "freshness-config.max-age-ms",
			InnerField: "maxAgeMs",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "freshness-config.prefer-recent-sources",
			InnerField: "preferRecentSources",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "freshness-config.recency-weight",
			InnerField: "recencyWeight",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "freshness-config.validate-links",
			InnerField: "validateLinks",
		},
	},
	"model-config": {
		&requestflag.InnerFlag[int64]{
			Name:       "model-config.max-output-tokens",
			InnerField: "maxOutputTokens",
		},
		&requestflag.InnerFlag[string]{
			Name:       "model-config.model-id",
			InnerField: "modelId",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "model-config.temperature",
			InnerField: "temperature",
		},
	},
	"recursion-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "recursion-config.enabled",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "recursion-config.max-depth",
			InnerField: "maxDepth",
		},
		&requestflag.InnerFlag[string]{
			Name:       "recursion-config.strategy",
			InnerField: "strategy",
		},
	},
	"search-config": {
		&requestflag.InnerFlag[[]string]{
			Name:       "search-config.exclude-domains",
			InnerField: "excludeDomains",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "search-config.include-domains",
			InnerField: "includeDomains",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "search-config.max-results",
			InnerField: "maxResults",
		},
		&requestflag.InnerFlag[string]{
			Name:       "search-config.search-depth",
			InnerField: "searchDepth",
		},
		&requestflag.InnerFlag[string]{
			Name:       "search-config.time-range",
			InnerField: "timeRange",
		},
		&requestflag.InnerFlag[string]{
			Name:       "search-config.topic",
			InnerField: "topic",
		},
	},
})

var profilesList = cli.Command{
	Name:            "list",
	Usage:           "Returns a list of intelligence profiles the user is subscribed to, including\nsubscription status and delivery preferences.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleProfilesList,
	HideHelpCommand: true,
}

var profilesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes an intelligence profile and all associated subscriptions.\nOnly profiles owned by the authenticated user can be deleted. This action cannot\nbe undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "profile-id",
			Required: true,
		},
	},
	Action:          handleProfilesDelete,
	HideHelpCommand: true,
}

var profilesPartialUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "partial-update",
	Usage:   "Partially updates an existing intelligence profile. Only the fields included in\nthe request body will be modified; all other fields remain unchanged. Only\nprofiles owned by the authenticated user can be updated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "profile-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "audio-config",
			BodyPath: "audioConfig",
		},
		&requestflag.Flag[string]{
			Name:     "bluf-structure",
			BodyPath: "blufStructure",
		},
		&requestflag.Flag[string]{
			Name:     "branding-template-id",
			Usage:    "Branding template ID (Pro feature)",
			BodyPath: "brandingTemplateId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "budget-config",
			BodyPath: "budgetConfig",
		},
		&requestflag.Flag[string]{
			Name:     "custom-prompt",
			BodyPath: "customPrompt",
		},
		&requestflag.Flag[string]{
			Name:     "frequency",
			Usage:    "Report generation frequency",
			BodyPath: "frequency",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "freshness-config",
			BodyPath: "freshnessConfig",
		},
		&requestflag.Flag[bool]{
			Name:     "is-community",
			BodyPath: "isCommunity",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "model-config",
			BodyPath: "modelConfig",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "recursion-config",
			BodyPath: "recursionConfig",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-day-of-month",
			BodyPath: "scheduleDayOfMonth",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-day-of-week",
			BodyPath: "scheduleDayOfWeek",
		},
		&requestflag.Flag[string]{
			Name:     "schedule-time-of-day",
			BodyPath: "scheduleTimeOfDay",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "search-config",
			BodyPath: "searchConfig",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Profile status",
			BodyPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "topic",
			BodyPath: "topic",
		},
	},
	Action:          handleProfilesPartialUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"audio-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "audio-config.enabled",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "audio-config.speed",
			InnerField: "speed",
		},
		&requestflag.InnerFlag[string]{
			Name:       "audio-config.voice-id",
			InnerField: "voiceId",
		},
	},
	"budget-config": {
		&requestflag.InnerFlag[float64]{
			Name:       "budget-config.alert-threshold",
			InnerField: "alertThreshold",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "budget-config.max-cost-per-report",
			InnerField: "maxCostPerReport",
		},
	},
	"freshness-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "freshness-config.enabled",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "freshness-config.max-age-ms",
			InnerField: "maxAgeMs",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "freshness-config.prefer-recent-sources",
			InnerField: "preferRecentSources",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "freshness-config.recency-weight",
			InnerField: "recencyWeight",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "freshness-config.validate-links",
			InnerField: "validateLinks",
		},
	},
	"model-config": {
		&requestflag.InnerFlag[int64]{
			Name:       "model-config.max-output-tokens",
			InnerField: "maxOutputTokens",
		},
		&requestflag.InnerFlag[string]{
			Name:       "model-config.model-id",
			InnerField: "modelId",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "model-config.temperature",
			InnerField: "temperature",
		},
	},
	"recursion-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "recursion-config.enabled",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "recursion-config.max-depth",
			InnerField: "maxDepth",
		},
		&requestflag.InnerFlag[string]{
			Name:       "recursion-config.strategy",
			InnerField: "strategy",
		},
	},
	"search-config": {
		&requestflag.InnerFlag[[]string]{
			Name:       "search-config.exclude-domains",
			InnerField: "excludeDomains",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "search-config.include-domains",
			InnerField: "includeDomains",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "search-config.max-results",
			InnerField: "maxResults",
		},
		&requestflag.InnerFlag[string]{
			Name:       "search-config.search-depth",
			InnerField: "searchDepth",
		},
		&requestflag.InnerFlag[string]{
			Name:       "search-config.time-range",
			InnerField: "timeRange",
		},
		&requestflag.InnerFlag[string]{
			Name:       "search-config.topic",
			InnerField: "topic",
		},
	},
})

func handleProfilesCreate(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.ProfileNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Profiles.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "profiles create", obj, format, transform)
}

func handleProfilesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("profile-id") && len(unusedArgs) > 0 {
		cmd.Set("profile-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.ProfileUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Profiles.Update(
		ctx,
		cmd.Value("profile-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "profiles update", obj, format, transform)
}

func handleProfilesList(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Profiles.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "profiles list", obj, format, transform)
}

func handleProfilesDelete(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("profile-id") && len(unusedArgs) > 0 {
		cmd.Set("profile-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Profiles.Delete(ctx, cmd.Value("profile-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "profiles delete", obj, format, transform)
}

func handleProfilesPartialUpdate(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("profile-id") && len(unusedArgs) > 0 {
		cmd.Set("profile-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.ProfilePartialUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Profiles.PartialUpdate(
		ctx,
		cmd.Value("profile-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "profiles partial-update", obj, format, transform)
}
