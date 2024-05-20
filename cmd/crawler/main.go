package main

import (
	"fmt"
	"os"

	"github.com/ecumenos-social/network-node/cmd/crawler/configurations"
	cli "github.com/urfave/cli/v2"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Println("exiting", "err", err)
		os.Exit(-1)
	}
}

func run(args []string) error {
	app := cli.App{
		Name:    configurations.ServiceName,
		Usage:   "crawling data from other nodes",
		Version: string(configurations.ServiceVersion),
		Flags:   []cli.Flag{},
		Commands: []*cli.Command{
			runAppCmd,
		},
	}

	return app.Run(args)
}
