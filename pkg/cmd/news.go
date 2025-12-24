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

var newsList = cli.Command{
	Name:  "list",
	Usage: "Returns news items from the GloriaAI terminal cache. Supports filtering by\ntopics and pagination.",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "topics",
			Usage:     "Comma-separated list of topics to filter by.\nValid topics: ai, ai_agents, aptos, base, bitcoin, crypto, dats, defi, ethereum, hyperliquid, machine_learning, macro, ondo, perps, ripple, rwa, solana, tech, virtuals.\nDefault: crypto, ai_agents, macro, bitcoin, ethereum, tech\n",
			QueryPath: "topics",
		},
	},
	Action:          handleNewsList,
	HideHelpCommand: true,
}

var newsGetRecaps = cli.Command{
	Name:  "get-recaps",
	Usage: "Returns AI-generated recap summaries for specified topics within a given\ntimeframe.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "timeframe",
			Usage:     "Time period for recap data",
			Default:   "12h",
			QueryPath: "timeframe",
		},
		&requestflag.Flag[string]{
			Name:      "topics",
			Usage:     "Comma-separated list of topics.\nValid topics: ai, ai_agents, aptos, base, bitcoin, crypto, dats, defi, ethereum, hyperliquid, machine_learning, macro, ondo, perps, ripple, rwa, solana, tech, virtuals\n",
			QueryPath: "topics",
		},
	},
	Action:          handleNewsGetRecaps,
	HideHelpCommand: true,
}

var newsListFeeds = cli.Command{
	Name:            "list-feeds",
	Usage:           "Returns all available news feed topics with descriptions.",
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

	params := y2.NewsListParams{}

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
	_, err = client.News.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "news list", obj, format, transform)
}

func handleNewsGetRecaps(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.NewsGetRecapsParams{}

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
	_, err = client.News.GetRecaps(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "news get-recaps", obj, format, transform)
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "news list-feeds", obj, format, transform)
}
