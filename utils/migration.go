package utils

import (
	"context"
	"fmt"
	"log"
	"math"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Migrate(collection *mongo.Collection, bulkWrites []mongo.WriteModel) {
	chunkSize := 20000
	totalRows := len(bulkWrites)
	n := math.Ceil(float64(totalRows) / float64(chunkSize))
	updatedDocuments := 0

	for i := 0; float64(i) < n; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > totalRows {
			end = totalRows
		}
		updated, err := collection.BulkWrite(context.TODO(), bulkWrites[start:end], options.BulkWrite())
		if err != nil {
			log.Fatal(err)
		}
		updatedDocuments += int(updated.MatchedCount)

		fmt.Println("Updated", updatedDocuments, "records")
	}

	fmt.Println("Done")
}
