// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/y2-cli/internal/apiquery"
	"github.com/stainless-sdks/y2-cli/internal/requestflag"
	"github.com/stainless-sdks/y2-go"
	"github.com/stainless-sdks/y2-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var reportsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Returns the full content of a specific intelligence report, including HTML\ncontent, sources, and audio metadata.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "report-id",
		},
	},
	Action:          handleReportsRetrieve,
	HideHelpCommand: true,
}

var reportsList = cli.Command{
	Name:  "list",
	Usage: "Returns a list of reports for the user's subscribed profiles. Results are sorted\nby generation date (newest first).",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of reports to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "profile-id",
			Usage:     "Filter reports by profile ID",
			QueryPath: "profileId",
		},
	},
	Action:          handleReportsList,
	HideHelpCommand: true,
}

var reportsRetrieveAudio = cli.Command{
	Name:  "retrieve-audio",
	Usage: "Returns audio file metadata or redirects to the CDN URL. Requires the\n`reports:audio` scope.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "report-id",
		},
		&requestflag.Flag[bool]{
			Name:      "redirect",
			Usage:     "If true, returns 302 redirect to audio CDN URL",
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Reports.Get(ctx, cmd.Value("report-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "reports retrieve", obj, format, transform)
}

func handleReportsList(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.ReportListParams{}

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
	_, err = client.Reports.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "reports list", obj, format, transform)
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

	params := y2.ReportGetAudioParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "reports retrieve-audio", obj, format, transform)
}
