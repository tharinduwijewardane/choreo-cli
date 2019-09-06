/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package main

import (
	"github.com/spf13/cobra"
	"github.com/wso2/choreo/components/cli/internal/pkg/cmd"
	cmdCommon "github.com/wso2/choreo/components/cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo/components/cli/internal/pkg/config"
)

func main() {

	cliConfig, err := config.InitConfig()
	if err != nil {
		cmdCommon.ExitWithError("Error loading configs", err)
	}

	command := cobra.Command{
		Use:   "choreo <command>",
		Short: "Manage integration applications with Choreo platform",
	}

	command.AddCommand(cmd.NewVersionCommand(cliConfig))

	if err := command.Execute(); err != nil {
		cmdCommon.ExitWithError("Error executing choreo command", err)
	}
}
