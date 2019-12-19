/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package runtime

import (
	"io"
)

type Reader interface {
	GetString(key string) string
	GetStringOrDefault(key string, defaultValue string) string
}

type Writer interface {
	SetString(key string, value string)
}

type UserConfig interface {
	Reader
	Writer
}

type EnvConfig interface {
	Reader
}

type EnvConfigHolder interface {
	EnvConfig() EnvConfig
}

type UserConfigHolder interface {
	UserConfig() UserConfig
}

type ConsoleOutHolder interface {
	Out() io.Writer
}

type ConsoleDebugOutHolder interface {
	DebugOut() io.Writer
}

type ConsoleWriterHolder interface {
	ConsoleOutHolder
	ConsoleDebugOutHolder
}

type Application struct {
	Id          string `json:"id" header:"Id"`
	Name        string `json:"name" header:"Application Name"`
	Description string `json:"description" header:"Description"`
}

type ApplicationRequest struct {
	Name        string `json:"name" header:"Application Name"`
	Description string `json:"description" header:"Description"`
}

type DeploymentInput struct {
	Url           string `json:"repo_url"`
	ApplicationId string `json:"app_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

type DeploymentOut struct {
	DeploymentUrl string `json:"deployment_url"`
	ApplicationId string `json:"app_id"`
}

type ApplicationApiClient interface {
	CreateNewApp(name string, desc string) error
	ListApps() ([]Application, error)
	CreateAndDeployApp(deploymentRequest DeploymentInput) (DeploymentOut, error)
	FetchLogs(appId string, linesCount uint) (string, error)
	DeleteApp(appId string) error
	GetApplicationStatus(appId string) (string, error)
}

type AuthApiClient interface {
	CreateOauthStateString() (string, error)
}

type Client interface {
	ApplicationApiClient
	AuthApiClient
}

type ClientHolder interface {
	Client() Client
}

type CliContext interface {
	ConsoleWriterHolder
	UserConfigHolder
	EnvConfigHolder
	ClientHolder
}
