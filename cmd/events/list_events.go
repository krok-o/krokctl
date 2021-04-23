package events

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/krok-o/krok/pkg/models"

	"github.com/krok-o/krokctl/cmd"
	"github.com/krok-o/krokctl/pkg/formatter"
)

var (
	// ListEventsCmd lists all events.
	ListEventsCmd = &cobra.Command{
		Use:   "events",
		Short: "List events",
		Run:   runListEventCmd,
	}
	listEventsArgs struct {
		page, pageSize, repoID int
		startDate, endDate     string
	}
)

func init() {
	cmd.ListCmd.AddCommand(ListEventsCmd)

	f := ListEventsCmd.PersistentFlags()
	f.IntVar(&listEventsArgs.page, "page", 0, "List events on the current page.")
	f.IntVar(&listEventsArgs.pageSize, "page-size", 10, "Maximum items per page.")
	f.IntVar(&listEventsArgs.repoID, "repository-id", -1, "The ID of the repository to list events for.")
	f.StringVar(&listEventsArgs.startDate, "start-date", "", "Start date to look events for ( including )")
	f.StringVar(&listEventsArgs.endDate, "end-date", "", "End date to look events for ( excluding )")
}

func runListEventCmd(c *cobra.Command, args []string) {
	var (
		fromDate *time.Time
		toDate   *time.Time
	)
	if listEventsArgs.startDate != "" {
		fd, err := cmd.ParseTime(listEventsArgs.startDate)
		if err != nil {
			cmd.CLILog.Fatal().Err(err).Msg("Failed to parse from date.")
		}
		fromDate = &fd
	}
	if listEventsArgs.endDate != "" {
		ed, err := cmd.ParseTime(listEventsArgs.endDate)
		if err != nil {
			cmd.CLILog.Fatal().Err(err).Msg("Failed to parse end date.")
		}
		toDate = &ed
	}
	opts := &models.ListOptions{
		PageSize:     listEventsArgs.pageSize,
		Page:         listEventsArgs.page,
		StartingDate: fromDate,
		EndDate:      toDate,
	}
	repos, err := cmd.KC.EventClient.List(listEventsArgs.repoID, opts)
	if err != nil {
		cmd.CLILog.Fatal().Err(err).Msg("Failed to list event.")
	}
	fmt.Print(formatter.FormatEvents(repos, cmd.KrokArgs.Formatter))
}
