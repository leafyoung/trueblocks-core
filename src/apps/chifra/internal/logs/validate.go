// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package logsPkg

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"

func (opts *LogsOptions) ValidateLogs() error {
	opts.TestLog()

	if opts.BadFlag != nil {
		return opts.BadFlag
	}

    if len(opts.Globals.File) > 0 {
        // Do nothing
    } else {
        if len(opts.Transactions) == 0 {
            return validate.Usage("Please supply one or more transaction identifiers.")
        }
    }

	return opts.Globals.ValidateGlobals()
}
