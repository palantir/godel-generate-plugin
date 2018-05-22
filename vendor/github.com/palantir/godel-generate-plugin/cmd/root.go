// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/palantir/godel/framework/pluginapi"
	"github.com/palantir/pkg/cobracli"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "generate",
		Short: "Run generators specified in configuration",
	}

	projectDirFlagVal string
	cfgFlagVal        string
)

func Execute() int {
	return cobracli.ExecuteWithDefaultParams(rootCmd)
}

func init() {
	pluginapi.AddProjectDirPFlagPtr(rootCmd.PersistentFlags(), &projectDirFlagVal)
	rootCmd.PersistentFlags().StringVar(&cfgFlagVal, "config", "", "the YAML configuration file for the generate task")
}
