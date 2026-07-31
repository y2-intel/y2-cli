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

var newsList = cli.Command{
	Name:    "list",
	Usage:   "Lists cached Y2 News Terminal items with topic filters and pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country-code",
			Usage:     "Filter by canonical ISO 3166-1 alpha-2 country code. When supplied without\n`topics`, the query searches every News Terminal topic.\n",
			QueryPath: "countryCode",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque continuation token from the previous response. Bound to the original filters and ordering.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Use `ndjson` for row-oriented streaming output.",
			QueryPath: "format",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "topics",
			Usage:     "Comma-separated list of topics to filter by.\nUse `GET /news/feeds` to discover the current topic catalog.\nDefault: crypto, geopolitics, macro, equities, ai, energy\n",
			QueryPath: "topics",
		},
	},
	Action:          handleNewsList,
	HideHelpCommand: true,
}

var newsGetRecaps = cli.Command{
	Name:    "get-recaps",
	Usage:   "Lists AI-generated recaps for selected topics and timeframe.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "timeframe",
			Usage:     "Time period for recap data",
			Default:   "12h",
			QueryPath: "timeframe",
		},
		&requestflag.Flag[string]{
			Name:      "topics",
			Usage:     "Comma-separated list of topics.\nUse `GET /news/feeds` to discover the current topic catalog.\n",
			QueryPath: "topics",
		},
	},
	Action:          handleNewsGetRecaps,
	HideHelpCommand: true,
}

var newsListFeeds = cli.Command{
	Name:            "list-feeds",
	Usage:           "Lists news feed topics and descriptions.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleNewsListFeeds,
	HideHelpCommand: true,
}

func handleNewsList(ctx context.Context, cmd *cli.Command) error {
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

	params := y2.NewsListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.News.List(ctx, params, options...)
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
		Title:          "news list",
		Transform:      transform,
	})
}

func handleNewsGetRecaps(ctx context.Context, cmd *cli.Command) error {
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

	params := y2.NewsGetRecapsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.News.GetRecaps(ctx, params, options...)
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
		Title:          "news get-recaps",
		Transform:      transform,
	})
}

func handleNewsListFeeds(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.News.ListFeeds(ctx, options...)
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
		Title:          "news list-feeds",
		Transform:      transform,
	})
}
