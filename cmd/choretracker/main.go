package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Galdoba/choretracker/cmd/choretracker/app"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/constants"
)

func main() {
	actx, err := appcontext.InitCli(constants.AppName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "initiation: %v", err)
		os.Exit(1)
	}
	program := app.NewApp(actx)

	if err := program.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v run error: %v", constants.AppName, err)
		os.Exit(1)
	}
}
