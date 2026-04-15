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

var subscriptionsUpdateDelivery = cli.Command{
	Name:    "update-delivery",
	Usage:   "Changes the delivery method for a subscription. When setting to `webhook`, a\nvalid `webhookConfigId` must be provided. The webhook must be active.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "subscription-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "delivery-method",
			Usage:    "Subscription delivery method",
			Required: true,
			BodyPath: "deliveryMethod",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-config-id",
			Usage:    `Required when deliveryMethod is "webhook"`,
			BodyPath: "webhookConfigId",
		},
	},
	Action:          handleSubscriptionsUpdateDelivery,
	HideHelpCommand: true,
}

func handleSubscriptionsUpdateDelivery(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := y2.SubscriptionUpdateDeliveryParams{}

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
	_, err = client.Subscriptions.UpdateDelivery(
		ctx,
		cmd.Value("subscription-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "subscriptions update-delivery", obj, format, explicitFormat, transform)
}
