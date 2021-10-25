package cmd

import (
	"fmt"
	"time"

	"github.com/meilisearch/meilisearch-go"
	"github.com/spf13/cobra"
)

var resetSettings bool
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup MeiliSearch index",
	RunE: func(cmd *cobra.Command, args []string) error {
		if resetSettings {
			reset, err := index.ResetSettings()
			if err != nil {
				return err
			}

			if err := waitForUpdate(reset.UpdateID); err != nil {
				return err
			}
		}

		filters, err := index.UpdateFilterableAttributes(&[]string{"type", "minplayers", "maxplayers", "maxplaytime", "rating", "weight"})
		if err != nil {
			return err
		}

		search, err := index.UpdateSearchableAttributes(&[]string{"name", "description", "categories", "mechanics"})
		if err != nil {
			return err
		}

		sort, err := index.UpdateSortableAttributes(&[]string{"name", "rating", "weight"})
		if err != nil {
			return err
		}

		if err := waitForUpdate(filters.UpdateID); err != nil {
			return err
		}
		if err := waitForUpdate(search.UpdateID); err != nil {
			return err
		}
		if err := waitForUpdate(sort.UpdateID); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	setupCmd.Flags().BoolVarP(&resetSettings, "reset", "r", false, "Reset index settings")

	rootCmd.AddCommand(setupCmd)
}

func waitForUpdate(updateID int64) error {
	for i := 0; i < 5; i++ {
		status, err := index.GetUpdateStatus(updateID)
		if err != nil {
			return err
		}
		if status.Status == meilisearch.UpdateStatusProcessed {
			return nil
		}
		if status.Status == meilisearch.UpdateStatusFailed {
			return fmt.Errorf("could not update index: %s", status.Error)
		}
		time.Sleep(time.Second)
	}

	return fmt.Errorf("could not verify update")
}
