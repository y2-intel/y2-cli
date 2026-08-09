// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
	"github.com/y2-intel/y2-cli/internal/apiquery"
	"github.com/y2-intel/y2-cli/internal/requestflag"
	"github.com/y2-intel/y2-go"
	"github.com/y2-intel/y2-go/option"
)

var reportsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a compact report by default. Use bounded `include` values or\n`view=agent`; request `text/markdown` for the canonical Markdown representation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "report-id",
			Required:  true,
			PathParam: "reportId",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Explicit representation override.",
			QueryPath: "format",
		},
		&requestflag.Flag[string]{
			Name:      "include",
			Usage:     "Comma-separated `content,sources,signals,graph,audio` expansions.",
			QueryPath: "include",
		},
		&requestflag.Flag[string]{
			Name:      "view",
			Usage:     "Compact projection optimized for grounded agent context.",
			QueryPath: "view",
		},
	},
	Action:          handleReportsRetrieve,
	HideHelpCommand: true,
}

var reportsList = cli.Command{
	Name:    "list",
	Usage:   "Lists reports for the user's subscribed profiles by generation date, newest\nfirst.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque continuation token from the previous response. Bound to the original filters and ordering.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "`json` uses the resource envelope; `ndjson` streams one canonical row per line.",
			QueryPath: "format",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of reports to return for this page.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "profile-id",
			Usage:     "Filter by stable public profile ID (`prf_...`).",
			QueryPath: "profileId",
		},
	},
	Action:          handleReportsList,
	HideHelpCommand: true,
}

var reportsRetrieveAudio = cli.Command{
	Name:    "retrieve-audio",
	Usage:   "Returns audio file metadata or redirects to the CDN URL. Requires the\n`reports:audio` scope.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "report-id",
			Required:  true,
			PathParam: "reportId",
		},
		&requestflag.Flag[bool]{
			Name:      "redirect",
			Usage:     "When true, redirects with `302` to the audio CDN URL",
			Default:   false,
			QueryPath: "redirect",
		},
	},
	Action:          handleReportsRetrieveAudio,
	HideHelpCommand: true,
}

func handleReportsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("report-id") && len(unusedArgs) > 0 {
		cmd.Set("report-id", unusedArgs[0])
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

	params := y2.ReportGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Reports.Get(
		ctx,
		cmd.Value("report-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "reports retrieve",
		Transform:      transform,
	})
}

func handleReportsList(ctx context.Context, cmd *cli.Command) error {
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

	params := y2.ReportListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Reports.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "reports list",
		Transform:      transform,
	})
}

func handleReportsRetrieveAudio(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("report-id") && len(unusedArgs) > 0 {
		cmd.Set("report-id", unusedArgs[0])
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

	params := y2.ReportGetAudioParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Reports.GetAudio(
		ctx,
		cmd.Value("report-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "reports retrieve-audio",
		Transform:      transform,
	})
}
