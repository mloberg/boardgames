package cmd

import (
	"os"

	"github.com/meilisearch/meilisearch-go"
	"github.com/spf13/cobra"
)

var index *meilisearch.Index
var rootCmd = &cobra.Command{
	Use:   "bg",
	Short: "Load board games into MeiliSearch",
	Long:  `Grab a user's collection from BoardGameGeek and load it into MeiliSearch`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		host, err := cmd.Flags().GetString("url")
		if err != nil {
			return err
		}

		key, err := cmd.Flags().GetString("key")
		if err != nil {
			return err
		}

		indexName, err := cmd.Flags().GetString("index")
		if err != nil {
			return err
		}

		meili := meilisearch.NewClient(meilisearch.ClientConfig{
			Host:   host,
			APIKey: key,
		})
		index = meili.Index(indexName)

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("url", "u", one(os.Getenv("MEILISEARCH_HOST"), "http://localhost:7700"), "The MeiliSearch URL")
	rootCmd.PersistentFlags().StringP("key", "k", os.Getenv("MEILISEARCH_API_KEY"), "The MeiliSearch API Key")
	rootCmd.PersistentFlags().StringP("index", "i", one(os.Getenv("MEILISEARCH_INDEX"), "boardgames"), "The MeiliSearch index")
}

func one(value ...string) string {
	for _, v := range value {
		if v != "" {
			return v
		}
	}
	return ""
}
