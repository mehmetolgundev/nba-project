package infra

import (
	"context"
	"log"
	"time"

	"github.com/mehmetolgundev/nba-project/domain/match/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	runMigration(client)
	return client

}
func runMigration(client *mongo.Client) {
	col := client.Database("FixtureDB").Collection("Matches")
	documentCount, err := col.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if documentCount > 0 {
		return
	}
	matches := []entities.Match{
		{
			HomeTeam: "Chicago Bulls",
			AwayTeam: "LA Clippers",
			DateTime: time.Date(2024, time.October, 15, 14, 10, 30, 0, time.UTC),
			Events: []entities.Event{
				{
					EventId:  "1",
					Type:     "MatchStarted",
					DateTime: time.Date(2024, time.October, 15, 14, 10, 30, 0, time.UTC),
					Time:     0,
				},
				{
					EventId:     "2",
					Type:        "Score",
					Team:        "Home",
					Player:      "Michael Porter Jr.",
					AsistPlayer: "Jerami Grant",
					Point:       2,
					Time:        4,
				},
				{
					EventId:     "3",
					Type:        "Score",
					Team:        "Away",
					Player:      "Damian Lillard",
					AsistPlayer: "Klay Thompson",
					Point:       3,
					Time:        12,
				},
				{
					EventId:     "4",
					Type:        "Score",
					Team:        "Away",
					Player:      "Kevin Durant",
					AsistPlayer: "",
					Point:       3,
					Time:        20,
				},
				{
					EventId:  "5",
					Type:     "MatchFinished",
					DateTime: time.Date(2024, time.October, 15, 14, 50, 30, 0, time.UTC),
					Time:     240,
				},
			},
		},
		{
			HomeTeam: "Indiana Pacers",
			AwayTeam: "Memphis Grizzlies",
			DateTime: time.Date(2024, 10, 15, 14, 10, 30, 0, time.UTC),
			Events: []entities.Event{
				{
					EventId:  "1",
					Type:     "MatchStarted",
					DateTime: time.Date(2024, 10, 15, 14, 10, 30, 0, time.UTC),
					Time:     0,
				},
				{
					EventId:     "2",
					Type:        "Score",
					Team:        "Home",
					Player:      "Anthony Davis",
					AsistPlayer: "Jimmy Butler",
					Point:       3,
					Time:        16,
				},
				{
					EventId:     "3",
					Type:        "Score",
					Team:        "Away",
					Player:      "Jayson Tatum",
					AsistPlayer: "Stephen Curry",
					Point:       2,
					Time:        34,
				},
				{
					EventId:     "4",
					Type:        "Score",
					Team:        "Home",
					Player:      "Luka Dončić",
					AsistPlayer: "James Harden",
					Point:       2,
					Time:        173,
				},
				{
					EventId:  "5",
					Type:     "MatchFinished",
					DateTime: time.Date(2024, 10, 15, 14, 50, 30, 0, time.UTC),
					Time:     240,
				},
			},
		},
	}
	documents := make([]interface{}, len(matches))
	for i, v := range matches {
		documents[i] = v
	}

	_, err = col.InsertMany(context.TODO(), documents)
	if err != nil {
		log.Fatal("Error inserting matches: ", err)
	}

}
