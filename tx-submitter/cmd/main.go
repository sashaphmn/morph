package main

import (
	"fmt"
	"os"

	"github.com/morph-l2/go-ethereum/log"
	"github.com/morph-l2/go-ethereum/params"
	"github.com/urfave/cli"

	submitter "morph-l2/tx-submitter"
	"morph-l2/tx-submitter/flags"
)

var (
	GitVersion = ""
	GitCommit  = ""
	GitDate    = ""
)

func main() {
	// Set up logger with a default INFO level in case we fail to parse flags.
	// Otherwise the final critical log won't show what the parsing error was.
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlInfo,
			log.StreamHandler(os.Stdout, log.TerminalFormat(false)),
		),
	)

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s", GitVersion, params.VersionWithCommit(GitCommit, GitDate))
	app.Name = "tx-submitter"
	app.Usage = "Batch Submitter Service"
	app.Description = "Service for generating and submitting batched transactions " +
		"that synchronize L2 state to L1 contracts"
	app.Action = submitter.Main()
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
