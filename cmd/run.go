// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/palantir/go-generate/commoncmd"
)

var (
	runCmd = commoncmd.NewRunCmd(
		"run [flags]",
		&projectDirFlagVal,
		&cfgFlagVal,
		&verifyFlagVal,
	)

	verifyFlagVal bool
)

func init() {
	runCmd.Flags().BoolVar(&verifyFlagVal, "verify", false, "verify that running generators does not change the current output")
	rootCmd.AddCommand(runCmd)
}
