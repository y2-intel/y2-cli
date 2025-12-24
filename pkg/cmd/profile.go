// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/y2-cli/internal/apiquery"
	"github.com/stainless-sdks/y2-go"
	"github.com/stainless-sdks/y2-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var profilesList = cli.Command{
	Name:            "list",
	Usage:           "Returns a list of intelligence profiles the user is subscribed to, including\nsubscription status and delivery preferences.",
	Flags:           []cli.Flag{},
	Action:          handleProfilesList,
	HideHelpCommand: true,
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
