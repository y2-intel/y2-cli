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

var webhooksCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a webhook configuration. Requires a paid workspace plan with webhook\naccess. The URL must use HTTPS and pass SSRF validation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Webhook display name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "Webhook endpoint URL (must be HTTPS)",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "headers",
			Usage:    "Custom headers to include in webhook deliveries",
			BodyPath: "headers",
		},
		&requestflag.Flag[string]{
			Name:     "secret",
			Usage:    "Shared secret for signature verification",
			BodyPath: "secret",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleWebhooksCreate,
	HideHelpCommand: true,
}

var webhooksUpdate = cli.Command{
	Name:    "update",
	Usage:   "Replaces every mutable webhook configuration field. `name` and `url` are\nrequired; omitted optional fields are reset to their defaults.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "webhook-id",
			Required:  true,
			PathParam: "webhookId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Webhook display name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "Webhook endpoint URL (must be HTTPS)",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "headers",
			Usage:    "Custom headers to include in webhook deliveries",
			BodyPath: "headers",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			BodyPath: "isActive",
		},
		&requestflag.Flag[string]{
			Name:     "secret",
			Usage:    "Shared secret for signature verification",
			BodyPath: "secret",
		},
		&requestflag.Flag[string]{
			Name:       "if-match",
			HeaderPath: "If-Match",
		},
	},
	Action:          handleWebhooksUpdate,
	HideHelpCommand: true,
}

var webhooksList = cli.Command{
	Name:            "list",
	Usage:           "Lists the authenticated user's webhook configurations. Masks secrets.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWebhooksList,
	HideHelpCommand: true,
}

var webhooksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a webhook configuration. Returns `409` if any subscription uses it.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "webhook-id",
			Required:  true,
			PathParam: "webhookId",
		},
		&requestflag.Flag[string]{
			Name:       "if-match",
			HeaderPath: "If-Match",
		},
	},
	Action:          handleWebhooksDelete,
	HideHelpCommand: true,
}

var webhooksTest = cli.Command{
	Name:    "test",
	Usage:   "Sends a test payload to the webhook URL. Returns `422` if the endpoint responds\nwith an error.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "webhook-id",
			Required:  true,
			PathParam: "webhookId",
		},
	},
	Action:          handleWebhooksTest,
	HideHelpCommand: true,
}

func handleWebhooksCreate(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := y2.WebhookNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Webhooks.New(ctx, params, options...)
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
		Title:          "webhooks create",
		Transform:      transform,
	})
}

func handleWebhooksUpdate(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := y2.WebhookUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Webhooks.Update(
		ctx,
		cmd.Value("webhook-id").(string),
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
		Title:          "webhooks update",
		Transform:      transform,
	})
}

func handleWebhooksList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Webhooks.List(ctx, options...)
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
		Title:          "webhooks list",
		Transform:      transform,
	})
}

func handleWebhooksDelete(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
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

	params := y2.WebhookDeleteParams{}

	return client.Webhooks.Delete(
		ctx,
		cmd.Value("webhook-id").(string),
		params,
		options...,
	)
}

func handleWebhooksTest(ctx context.Context, cmd *cli.Command) error {
	client := y2.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
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
	_, err = client.Webhooks.Test(ctx, cmd.Value("webhook-id").(string), options...)
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
		Title:          "webhooks test",
		Transform:      transform,
	})
}
