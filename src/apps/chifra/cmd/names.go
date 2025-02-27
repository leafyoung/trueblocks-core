// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * This file was auto generated with makeClass --gocmds. DO NOT EDIT.
 */

package cmd

// EXISTING_CODE
import (
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/internal/globals"
	namesPkg "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/internal/names"
	"github.com/spf13/cobra"
)

// EXISTING_CODE

// namesCmd represents the names command
var namesCmd = &cobra.Command{
	Use:     usageNames,
	Short:   shortNames,
	Long:    longNames,
	Version: versionText,
	RunE:    namesPkg.RunNames,
}

var usageNames = `names [flags] <term> [term...]

Arguments:
  terms - a space separated list of one or more search terms (required)`

var shortNames = "query addresses or names of well known accounts"

var longNames = `Purpose:
  Query addresses or names of well known accounts.`

var notesNames = `
Notes:
  - The tool will accept up to three terms, each of which must match against any field in the database.
  - The --match_case option enables case sensitive matching.`

func init() {
	namesCmd.Flags().SortFlags = false

	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Expand, "expand", "e", false, "expand search to include all fields (search name, address, and symbol otherwise)")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().MatchCase, "match_case", "m", false, "do case-sensitive search")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().All, "all", "l", false, "include all accounts in the search")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Custom, "custom", "c", false, "include your custom named accounts")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Prefund, "prefund", "p", false, "include prefund accounts")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Named, "named", "n", false, "include well know token and airdrop addresses in the search")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Addr, "addr", "a", false, "display only addresses in the results (useful for scripting)")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Collections, "collections", "s", false, "display collections data")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Tags, "tags", "g", false, "export the list of tags and subtags only")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().ToCustom, "to_custom", "u", false, "for editCmd only, is the edited name a custom name or not (hidden)")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Clean, "clean", "C", false, "clean the data (addrs to lower case, sort by addr) (hidden)")
	namesCmd.Flags().StringVarP(&namesPkg.GetOptions().Autoname, "autoname", "A", "", "an address assumed to be a token, added automatically to names database if true (hidden)")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Create, "create", "", false, "create a new name record (hidden)")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Update, "update", "", false, "edit an existing name (hidden)")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Delete, "delete", "", false, "delete a name, but do not remove it (hidden)")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Undelete, "undelete", "", false, "undelete a previously deleted name (hidden)")
	namesCmd.Flags().BoolVarP(&namesPkg.GetOptions().Remove, "remove", "", false, "remove a previously deleted name (hidden)")
	if os.Getenv("TEST_MODE") != "true" {
		namesCmd.Flags().MarkHidden("to_custom")
		namesCmd.Flags().MarkHidden("clean")
		namesCmd.Flags().MarkHidden("autoname")
		namesCmd.Flags().MarkHidden("create")
		namesCmd.Flags().MarkHidden("update")
		namesCmd.Flags().MarkHidden("delete")
		namesCmd.Flags().MarkHidden("undelete")
		namesCmd.Flags().MarkHidden("remove")
	}
	globals.InitGlobals(namesCmd, &namesPkg.GetOptions().Globals)

	namesCmd.SetUsageTemplate(UsageWithNotes(notesNames))
	namesCmd.SetOut(os.Stderr)

	// EXISTING_CODE
	// EXISTING_CODE

	chifraCmd.AddCommand(namesCmd)
}
