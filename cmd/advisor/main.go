package advisor

import (
	"context"
	"customer-data-migration/migration/advisor"
	"customer-data-migration/model"
	"log"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var AdvisorCmd = &cobra.Command{
	Use:   "advisor",
	Short: "advisor",
	Long:  `import advisor tier`,
	Run: func(cmd *cobra.Command, args []string) {
		uri, _ := cmd.Flags().GetString("uri")
		db, _ := cmd.Flags().GetString("db")
		collection, _ := cmd.Flags().GetString("collection")
		path, _ := cmd.Flags().GetString("path")
		arguments := model.Argument{
			Uri:        uri,
			Db:         db,
			Collection: collection,
			Path:       path,
		}
		main(arguments)
	},
}

func init() {
	AdvisorCmd.Flags().String("uri", "", "mongo uri")
	AdvisorCmd.Flags().String("db", "", "mongo db")
	AdvisorCmd.Flags().String("collection", "", "mongo collection")
	AdvisorCmd.Flags().String("path", "", "path to csv")
}

func main(args model.Argument) {
	clientOptions := options.Client().ApplyURI(args.Uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(args.Db).Collection(args.Collection)
	config := model.MigrationConfig{
		Argument:   args,
		Collection: collection,
	}
	advisor.Migrate(config)
}
