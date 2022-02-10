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

			if err := waitForUpdate(reset.UID); err != nil {
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

		if err := waitForUpdate(filters.UID); err != nil {
			return err
		}
		if err := waitForUpdate(search.UID); err != nil {
			return err
		}
		if err := waitForUpdate(sort.UID); err != nil {
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
		task, err := index.GetTask(updateID)
		if err != nil {
			return err
		}

		if task.Status == meilisearch.TaskStatusSucceeded {
			return nil
		}

		if task.Status == meilisearch.TaskStatusFailed {
			return fmt.Errorf("could not update index: %s", task.Error.Message)
		}

		time.Sleep(time.Second)
	}

	return fmt.Errorf("could not verify update")
}
