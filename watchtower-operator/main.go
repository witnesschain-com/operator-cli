package main

import (
	"fmt"
	"os"

	wc_common "github.com/witnesschain-com/operator-cli/common"
	operator_commands "github.com/witnesschain-com/operator-cli/watchtower-operator/commands"

	"github.com/urfave/cli/v2"
)

var VERSION string = "UNKNOWN"

func main() {
	cli.AppHelpTemplate = fmt.Sprintf(`
	__        ___ _                        ____ _           _       
	\ \      / (_) |_ _ __   ___  ___ ___ / ___| |__   __ _(_)_ __  
	 \ \ /\ / /| | __|  _ \ / _ \/ __/ __| |   |  _ \ / _  | |  _ \ 
	  \ V  V / | | |_| | | |  __/\__ \__ \ |___| | | | (_| | | | | |
	   \_/\_/  |_|\__|_| |_|\___||___/___/\____|_| |_|\__,_|_|_| |_|

	%s`, cli.AppHelpTemplate)
	app := cli.NewApp()

	app.Name = "WitnessChain"
	app.Usage = "WitnessChain Operator CLI"
	app.Version = VERSION
	app.Copyright = "(c) 2024 WitnessChain"

	app.Commands = []*cli.Command{
		wc_common.KeysCmd(),
		operator_commands.RegisterWatchtowerCmd(),
		operator_commands.DeRegisterWatchtowerCmd(),
		operator_commands.RegisterOperatorToAVSCmd(),
		operator_commands.DeRegisterOperatorFromAVSCmd(),
	}

	if err := app.Run(os.Args); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		wc_common.Unmount()
		os.Exit(1)
	}

	wc_common.Unmount()
}
