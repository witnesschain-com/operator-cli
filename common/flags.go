package wc_common

import "github.com/urfave/cli/v2"

var (
	KeyNameFlag = cli.StringFlag{
		Name:     "key-name",
		Aliases:  []string{"k"},
		Required: true,
		Usage:    "Name of the key to be created",
	}

	InsecureFlag = cli.BoolFlag{
		Name:    "insecure",
		Aliases: []string{"i"},
		Usage:   "Use this flag to skip strong password validation",
		EnvVars: []string{"INSECURE"},
	}

	ConfigPathFlag = cli.StringFlag{
		Name:    "config-file",
		Aliases: []string{"c"},
		Usage:   "Path of the config file",
		EnvVars: []string{"CONFIG_PATH"},
	}
)
