package main

import (
	"context"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/service/v2"

	"github.com/starudream/notify-filterbox/server"
)

func init() {
	args := cobra.FlagArgs(rootCmd.PersistentFlags(), "config")
	if c := config.LoadedFile(); c != "" {
		args = append(args, "-c", c)
	}
	service.AddCommand(rootCmd, service.New("notify-filterbox", serviceRun, service.WithArguments(args...)))
}

func serviceRun(context.Context) {
	err := server.Start()
	if err != nil {
		slog.Error("server run error: %v", err)
	}
}
