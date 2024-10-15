package match

import (
	"context"
	"log"

	"github.com/mehmetolgundev/nba-project/domain/match/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchRepository struct {
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

const CollectionName = "Matches"
const DatabaseName = "FixtureDB"

func NewRepository(mongoClient *mongo.Client) MatchRepository {
	collection := mongoClient.Database(DatabaseName).Collection(CollectionName)
	return MatchRepository{mongoClient: mongoClient, collection: collection}
}
func (r MatchRepository) GetMatches(ctx context.Context) []entities.Match {
	filter := bson.D{}
	var results []entities.Match
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var match entities.Match
		if err := cur.Decode(&match); err != nil {
			log.Fatal(err)
		}

		results = append(results, match)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results

}
