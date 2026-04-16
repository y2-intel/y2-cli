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

var osintGetConflictIndicators = cli.Command{
	Name:    "get-conflict-indicators",
	Usage:   "Returns the Conflict Indicators Index (CII) values. Each item represents a\nconflict indicator with a score from 0-100 and a delta showing recent change.\nSupports filtering by region and category.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "category",
			Usage:     "Filter by event category",
			QueryPath: "category",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Filter by geographic region",
			QueryPath: "region",
		},
	},
	Action:          handleOsintGetConflictIndicators,
	HideHelpCommand: true,
}

var osintGetGpsJammingZones = cli.Command{
	Name:    "get-gps-jamming-zones",
	Usage:   "Returns GPS interference zones detected via ADS-B navigation accuracy\ndegradation analysis, aggregated into H3 hex cells.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of zones to return",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "severity",
			Usage:     "Filter by interference severity",
			QueryPath: "severity",
		},
	},
	Action:          handleOsintGetGpsJammingZones,
	HideHelpCommand: true,
}

var osintGetMilitaryPosture = cli.Command{
	Name:    "get-military-posture",
	Usage:   "Returns military posture assessments for monitored theaters, based on detected\nmilitary aircraft activity from the OpenSky Network. Each theater has a posture\nlevel (normal, elevated, critical) and aircraft breakdown by type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
	},
	Action:          handleOsintGetMilitaryPosture,
	HideHelpCommand: true,
}

var osintListAircraft = cli.Command{
	Name:    "list-aircraft",
	Usage:   "Returns tracked military aircraft positions from the OpenSky Network, filtered\nand classified by type (tanker, AWACS, fighter, etc.).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of aircraft to return",
			Default:   100,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "theater",
			Usage:     `Filter by theater ID (e.g. "iran", "taiwan", "baltic")`,
			QueryPath: "theater",
		},
	},
	Action:          handleOsintListAircraft,
	HideHelpCommand: true,
}

var osintListEvents = cli.Command{
	Name:    "list-events",
	Usage:   "Returns OSINT threat events from the Situation Room. Supports filtering by\ncategory, severity, region, and country.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "category",
			Usage:     "Filter by event category",
			QueryPath: "category",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of events to return",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "severity",
			Usage:     "Filter by severity level",
			QueryPath: "severity",
		},
	},
	Action:          handleOsintListEvents,
	HideHelpCommand: true,
}

var osintListVessels = cli.Command{
	Name:    "list-vessels",
	Usage:   "Returns naval vessel positions sourced from USNI fleet tracker data, including\ncarrier strike groups and individual warships.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of vessels to return",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Filter by region name",
			QueryPath: "region",
		},
	},
	Action:          handleOsintListVessels,
	HideHelpCommand: true,
}

var osintMapEvents = cli.Command{
	Name:    "map-events",
	Usage:   "Returns OSINT events with geographic coordinates for map display. Events without\ncoordinates are excluded.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of events to return",
			Default:   200,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Filter by geographic region",
			QueryPath: "region",
		},
	},
	Action:          handleOsintMapEvents,
	HideHelpCommand: true,
}

func handleOsintGetConflictIndicators(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.OsintGetConflictIndicatorsParams{}

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
	_, err = client.Osint.GetConflictIndicators(ctx, params, options...)
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
		Title:          "osint get-conflict-indicators",
		Transform:      transform,
	})
}

func handleOsintGetGpsJammingZones(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.OsintGetGpsJammingZonesParams{}

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
	_, err = client.Osint.GetGpsJammingZones(ctx, params, options...)
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
		Title:          "osint get-gps-jamming-zones",
		Transform:      transform,
	})
}

func handleOsintGetMilitaryPosture(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.OsintGetMilitaryPostureParams{}

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
	_, err = client.Osint.GetMilitaryPosture(ctx, params, options...)
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
		Title:          "osint get-military-posture",
		Transform:      transform,
	})
}

func handleOsintListAircraft(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.OsintListAircraftParams{}

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
	_, err = client.Osint.ListAircraft(ctx, params, options...)
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
		Title:          "osint list-aircraft",
		Transform:      transform,
	})
}

func handleOsintListEvents(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.OsintListEventsParams{}

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
	_, err = client.Osint.ListEvents(ctx, params, options...)
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
		Title:          "osint list-events",
		Transform:      transform,
	})
}

func handleOsintListVessels(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.OsintListVesselsParams{}

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
	_, err = client.Osint.ListVessels(ctx, params, options...)
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
		Title:          "osint list-vessels",
		Transform:      transform,
	})
}

func handleOsintMapEvents(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.OsintMapEventsParams{}

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
	_, err = client.Osint.MapEvents(ctx, params, options...)
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
		Title:          "osint map-events",
		Transform:      transform,
	})
}
