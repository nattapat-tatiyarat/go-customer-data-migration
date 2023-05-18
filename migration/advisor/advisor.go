package advisor

import (
	"customer-data-migration/model"
	"customer-data-migration/utils"
	"errors"
	"io"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Migrate(config model.MigrationConfig) {
	csv := utils.ReadCsv(config.Argument.Path)
	bulkWrites := []mongo.WriteModel{}

	for {
		line, err := csv.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				log.Fatal(err)
			}
		}

		userID, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		tier := line[1]

		user := model.AdvisorTier{
			UserID: userID,
			Tier:   tier,
		}

		updateModel := mongo.NewUpdateOneModel()
		updateModel.SetFilter(bson.M{"user_id": user.UserID})
		updateModel.SetUpdate(bson.M{"$set": bson.M{"advisor.tier": user.Tier, "updated_at": time.Now()}})
		updateModel.SetUpsert(true)

		bulkWrites = append(bulkWrites, updateModel)
	}

	utils.Migrate(config.Collection, bulkWrites)
}
