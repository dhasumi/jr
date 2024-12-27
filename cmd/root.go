/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jr",
	Short: "This is a custom Jira CLI wrapper command for personal use",
	Long:  `This is a custom Jira CLI wrapper command for personal use.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if debug {
		var debugLevel = new(slog.LevelVar)
		debugLevel.Set(slog.LevelDebug)

		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: debugLevel}))
		slog.SetDefault(logger)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jr.yaml)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
}
