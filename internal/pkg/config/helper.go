/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package config

func GetUserConfigReader(cliConfig Config, configDefinition map[int]KeyEntry) func(entry int) string {
	return func(entry int) string {
		return cliConfig.GetStringForKeyEntry(configDefinition[entry])
	}
}

func GetEnvironmentConfigReader(cliConfig Config, configDefinition map[int]KeyEntry) func(entry int) string {
	configManager := cliConfig.GetEnvironmentConfig()

	return func(entry int) string {
		return configManager.GetStringForKeyEntry(configDefinition[entry])
	}
}
