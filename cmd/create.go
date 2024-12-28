/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/dhasumi/jr/internal/jira"
	"github.com/spf13/cobra"
)

var params = jira.CreateParams{}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use: "create [flags] summary",
	// 	Short: "A brief description of your command",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:
	//
	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	ArgAliases: []string{
		"summary",
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if debug {
			var debugLevel = new(slog.LevelVar)
			debugLevel.Set(slog.LevelDebug)

			logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: debugLevel}))
			slog.SetDefault(logger)

			slog.Info("Debug Mode ON")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		params.Summary = args[0]

		sprint_id := jira.GetSprintID(params.SprintData)
		slog.Debug("createCmd.Run", "sprint_id", sprint_id)

		ticket_id, ticket_url := jira.CreateTicket(params)
		slog.Debug("createCmd.Run", "ticket_id", sprint_id)

		if sprint_id != "backlog" {
			err := jira.MoveTicketToSprint(ticket_id, sprint_id)
			if err != nil {
				slog.Error("createCmd.Run", "error", err)
				os.Exit(1)
			}
		}
		slog.Info("A ticket was successfully created", "ticket_id", ticket_id, "sprint_id", sprint_id, "ticket_url", ticket_url)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	createCmd.Flags().StringVarP(&params.Type, "type", "t", "Task", "Specifies Ticket Type e.g. Task, Bug, Story, Subtask")
	createCmd.Flags().StringVarP(&params.Body, "body", "b", "", "Specifies description text")
	createCmd.Flags().StringVarP(&params.Priority, "priority", "p", "Medium", "Specifies ticket priority e.g. Medium, High, Low, Highest, Lowerst")
	createCmd.Flags().StringSliceVarP(&params.Labels, "label", "l", []string{}, "Specifies ticket label. it can be put multiple labels with comma, like 'label1,label2'")
	createCmd.Flags().StringVarP(&params.Epic, "epic", "e", "", "Specifies EpicID")
	createCmd.Flags().StringVarP(&params.Assign, "assign", "a", "", "Specifies ticket assigner")
	createCmd.Flags().Uint8Var(&params.StoryPoints, "sp", 0, "Specifies StoryPoints value")
	createCmd.Flags().Int32VarP(&params.SprintData.SprintNum, "sprint", "s", 0, "Put the number of sprint")
	createCmd.Flags().BoolVar(&params.SprintData.NextSprint, "next-sprint", false, "Specifies if ticket should be located on the next sprint")
	createCmd.Flags().BoolVar(&params.SprintData.BackLog, "backlog", false, "Specifies if ticket should be located to backlog (no Sprint)")
	createCmd.Flags().Int32Var(&params.SprintData.FutureSprint, "future-sprint", 0, "Specifies the number of sprint ahead the ticket will be located from current")
	createCmd.Flags().StringVar(&params.TemplatePath, "template", "", "Specifies template file path to fill the description field")
}
