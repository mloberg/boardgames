package cmd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/meilisearch/meilisearch-go"
	"github.com/mloberg/boardgames/internal/bgg"
	"github.com/spf13/cobra"
)

type document map[string]interface{}

var refresh bool
var clear bool
var indexCmd = &cobra.Command{
	Use:   "index <username>",
	Args:  cobra.ExactArgs(1),
	Short: "Load board games into MeiliSearch",
	Long:  `Grab a user's collection from BoardGameGeek and load it into MeiliSearch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := bgg.New()
		collection, err := client.GetCollection(args[0])
		if err != nil {
			return err
		}

		var docs []document
		for _, i := range collection.Items {
			fmt.Print(".")
			var doc document
			if !refresh {
				err = index.GetDocument(strconv.Itoa(i.ID), &doc)
				if err != nil && err.(*meilisearch.Error).StatusCode != 404 {
					return err
				}
			}

			if doc == nil {
				thing, err := client.GetThing(i.ID)
				if err != nil {
					return err
				}

				rating, err := strconv.ParseFloat(thing.Item.Statistics.Rating.Value, 64)
				if err != nil {
					return err
				}
				weight, err := strconv.ParseFloat(thing.Item.Statistics.Weight.Value, 64)
				if err != nil {
					return err
				}

				doc = map[string]interface{}{
					"id":          thing.Item.ID,
					"type":        thing.Item.Type,
					"name":        thing.Item.Name(),
					"description": thing.Item.Description,
					"image":       thing.Item.Image,
					"thumbnail":   thing.Item.Thumbnail,
					"minplayers":  thing.Item.MinPlayers.Value,
					"maxplayers":  thing.Item.MaxPlayers.Value,
					"minplaytime": thing.Item.MinPlayTime.Value,
					"maxplaytime": thing.Item.MaxPlayTime.Value,
					"rating":      math.Round(rating*10) / 10,
					"weight":      math.Round(weight*100) / 100,
					"categories":  thing.Item.Categories(),
					"mechanics":   thing.Item.Mechanics(),
				}
			}

			docs = append(docs, doc)
		}
		fmt.Println()

		if clear {
			delete, err := index.DeleteAllDocuments()
			if err != nil {
				return err
			}
			if err := waitForUpdate(delete.UID); err != nil {
				return err
			}
		}

		update, err := index.AddDocuments(docs)
		if err != nil {
			return err
		}

		return waitForUpdate(update.UID)
	},
}

func init() {
	indexCmd.Flags().BoolVarP(&refresh, "refresh", "r", false, "Replace documents with fresh data")
	indexCmd.Flags().BoolVarP(&clear, "clear", "c", false, "Delete all documents")

	rootCmd.AddCommand(indexCmd)
}
