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

var osintCountriesGetCountryInstabilityIndex = cli.Command{
	Name:    "get-country-instability-index",
	Usage:   "Returns a country's Conflict Indicators Index (CII) score, baseline, delta, and\ncomponents.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country-code",
			Required:  true,
			PathParam: "countryCode",
		},
	},
	Action:          handleOsintCountriesGetCountryInstabilityIndex,
	HideHelpCommand: true,
}

var osintCountriesGetCountryNews = cli.Command{
	Name:    "get-country-news",
	Usage:   "Returns recent country news from the OSINT event pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country-code",
			Required:  true,
			PathParam: "countryCode",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque continuation token from the previous response. Bound to the original filters and ordering.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Select the JSON resource envelope, row-oriented NDJSON, or an RFC 7946 FeatureCollection.",
			QueryPath: "format",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of news items to return",
			Default:   8,
			QueryPath: "limit",
		},
	},
	Action:          handleOsintCountriesGetCountryNews,
	HideHelpCommand: true,
}

var osintCountriesGetIntelligenceBrief = cli.Command{
	Name:    "get-intelligence-brief",
	Usage:   "Returns a periodically generated, cached intelligence brief for a country.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country-code",
			Required:  true,
			PathParam: "countryCode",
		},
	},
	Action:          handleOsintCountriesGetIntelligenceBrief,
	HideHelpCommand: true,
}

var osintCountriesGetPredictionMarkets = cli.Command{
	Name:    "get-prediction-markets",
	Usage:   "Returns prediction-market probabilities and trading volumes for a country.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country-code",
			Required:  true,
			PathParam: "countryCode",
		},
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
			Usage:     "Maximum number of predictions to return",
			Default:   3,
			QueryPath: "limit",
		},
	},
	Action:          handleOsintCountriesGetPredictionMarkets,
	HideHelpCommand: true,
}

var osintCountriesGetStockMarketIndex = cli.Command{
	Name:    "get-stock-market-index",
	Usage:   "Returns a country's primary stock index, weekly change, and currency.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country-code",
			Required:  true,
			PathParam: "countryCode",
		},
	},
	Action:          handleOsintCountriesGetStockMarketIndex,
	HideHelpCommand: true,
}

func handleOsintCountriesGetCountryInstabilityIndex(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("country-code") && len(unusedArgs) > 0 {
		cmd.Set("country-code", unusedArgs[0])
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
	_, err = client.Osint.Countries.GetCountryInstabilityIndex(ctx, cmd.Value("country-code").(string), options...)
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
		Title:          "osint:countries get-country-instability-index",
		Transform:      transform,
	})
}

func handleOsintCountriesGetCountryNews(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("country-code") && len(unusedArgs) > 0 {
		cmd.Set("country-code", unusedArgs[0])
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

	params := y2.OsintCountryGetCountryNewsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Osint.Countries.GetCountryNews(
		ctx,
		cmd.Value("country-code").(string),
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
		Title:          "osint:countries get-country-news",
		Transform:      transform,
	})
}

func handleOsintCountriesGetIntelligenceBrief(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("country-code") && len(unusedArgs) > 0 {
		cmd.Set("country-code", unusedArgs[0])
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
	_, err = client.Osint.Countries.GetIntelligenceBrief(ctx, cmd.Value("country-code").(string), options...)
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
		Title:          "osint:countries get-intelligence-brief",
		Transform:      transform,
	})
}

func handleOsintCountriesGetPredictionMarkets(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("country-code") && len(unusedArgs) > 0 {
		cmd.Set("country-code", unusedArgs[0])
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

	params := y2.OsintCountryGetPredictionMarketsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Osint.Countries.GetPredictionMarkets(
		ctx,
		cmd.Value("country-code").(string),
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
		Title:          "osint:countries get-prediction-markets",
		Transform:      transform,
	})
}

func handleOsintCountriesGetStockMarketIndex(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("country-code") && len(unusedArgs) > 0 {
		cmd.Set("country-code", unusedArgs[0])
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
	_, err = client.Osint.Countries.GetStockMarketIndex(ctx, cmd.Value("country-code").(string), options...)
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
		Title:          "osint:countries get-stock-market-index",
		Transform:      transform,
	})
}
