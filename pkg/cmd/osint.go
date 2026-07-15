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
	Usage:   "Lists Conflict Indicators Index (CII) values with 0–100 scores and recent-change\ndeltas. Supports region and category filters.",
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
	Usage:   "Lists GPS interference zones inferred from ADS-B navigation-accuracy degradation\nand aggregated into H3 cells.",
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
	Usage:   "Lists theater posture assessments based on Wingbits ADS-B military aircraft\nactivity. Each includes a `normal`, `elevated`, or `critical` posture and\naircraft counts by type.",
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
	Usage:   "Lists Wingbits ADS-B military aircraft positions, classified by type such as\ntanker, AWACS, or fighter.",
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
			Usage:     `Filter by theater ID (e.g. "iran", "taiwan", "blacksea", "scs")`,
			QueryPath: "theater",
		},
	},
	Action:          handleOsintListAircraft,
	HideHelpCommand: true,
}

var osintListEvents = cli.Command{
	Name:    "list-events",
	Usage:   "Lists Situation Room threat events. Supports category and severity filters.",
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
	Usage:   "Lists USNI fleet-tracker positions for carrier strike groups and warships.",
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
	Usage:   "Lists geolocated OSINT events for map display. Excludes events without\ncoordinates.",
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

	params := y2.OsintGetConflictIndicatorsParams{}

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
		RawOutput:      cmd.Root().Bool("raw-output"),
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

	params := y2.OsintGetGpsJammingZonesParams{}

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
		RawOutput:      cmd.Root().Bool("raw-output"),
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

	params := y2.OsintGetMilitaryPostureParams{}

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
		RawOutput:      cmd.Root().Bool("raw-output"),
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

	params := y2.OsintListAircraftParams{}

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
		RawOutput:      cmd.Root().Bool("raw-output"),
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

	params := y2.OsintListEventsParams{}

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
		RawOutput:      cmd.Root().Bool("raw-output"),
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

	params := y2.OsintListVesselsParams{}

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
		RawOutput:      cmd.Root().Bool("raw-output"),
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

	params := y2.OsintMapEventsParams{}

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
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "osint map-events",
		Transform:      transform,
	})
}
