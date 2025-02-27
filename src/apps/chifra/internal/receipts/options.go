// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * This file was auto generated with makeClass --gocmds. DO NOT EDIT.
 */

package receiptsPkg

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/internal/globals"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/blockRange"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
)

type ReceiptsOptions struct {
	Transactions   []string
	TransactionIds []blockRange.Identifier
	Articulate     bool
	Globals        globals.GlobalOptions
	BadFlag        error
}

var receiptsCmdLineOptions ReceiptsOptions

func (opts *ReceiptsOptions) TestLog() {
	logger.TestLog(len(opts.Transactions) > 0, "Transactions: ", opts.Transactions)
	logger.TestLog(opts.Articulate, "Articulate: ", opts.Articulate)
	opts.Globals.TestLog()
}

func (opts *ReceiptsOptions) ToCmdLine() string {
	options := ""
	if opts.Articulate {
		options += " --articulate"
	}
	options += " " + strings.Join(opts.Transactions, " ")
	options += fmt.Sprintf("%s", "") // silence go compiler for auto gen
	return options
}

func ReceiptsFinishParseApi(w http.ResponseWriter, r *http.Request) *ReceiptsOptions {
	opts := &ReceiptsOptions{}
	for key, value := range r.URL.Query() {
		switch key {
		case "transactions":
			for _, val := range value {
				s := strings.Split(val, " ") // may contain space separated items
				opts.Transactions = append(opts.Transactions, s...)
			}
		case "articulate":
			opts.Articulate = true
		default:
			if !globals.IsGlobalOption(key) {
				opts.BadFlag = validate.Usage("Invalid key ({0}) in {1} route.", key, "receipts")
				return opts
			}
		}
	}
	opts.Globals = *globals.GlobalsFinishParseApi(w, r)
	// EXISTING_CODE
	// EXISTING_CODE

	return opts
}

func ReceiptsFinishParse(args []string) *ReceiptsOptions {
	opts := GetOptions()
	opts.Globals.FinishParse(args)
	defFmt := "txt"
	// EXISTING_CODE
	opts.Transactions = args
	// EXISTING_CODE
	if len(opts.Globals.Format) == 0 || opts.Globals.Format == "none" {
		opts.Globals.Format = defFmt
	}
	return opts
}

func GetOptions() *ReceiptsOptions {
	// EXISTING_CODE
	// EXISTING_CODE
	return &receiptsCmdLineOptions
}
