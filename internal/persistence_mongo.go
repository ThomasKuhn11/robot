package internal

import (
	"context"
	"fmt"
	"log"

	"time"

	"go-gin-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"github.com/sclevine/agouti"
)

// CRUD functions
func setupDB(ctx context.Context) *mongo.Database {
	client := createMongoConnection()
	database := client.Database("Database0809")
	//temperatureCollection := database.Collection("Temp")

	return database
}

func insertIntoDatabase(ctx context.Context, temp model.Temperature) {

	database := setupDB(ctx)
	temperatureCollection := database.Collection("Temp")
	//insert into database

	temperatureResult, err := temperatureCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"Period of the day", "Morning"},
			{"Temperature", temp.MorningT},
			{"Date", time.Now()},
		},
		bson.D{
			{"Period of the day", "Afternoon"},
			{"Temperature", temp.AfternoonT},
			{"Date", time.Now()},
		},
		bson.D{
			{"Period of the day", "Evening"},
			{"Temperature", temp.EveningT},
			{"Date", time.Now()},
		},
		bson.D{
			{"Period of the day", "Night"},
			{"Temperature", temp.NightT},
			{"Date", time.Now()},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(temperatureResult.InsertedIDs)

}

func readFromDatabase(ctx context.Context) {
	database := setupDB(ctx)
	temperatureCollection := database.Collection("Temp")

	// read from database

	cursor, err := temperatureCollection.Find(ctx, bson.M{"Period of the day": "Morning"})
	if err != nil {
		log.Fatal(err)
	}

	var weather []bson.M
	if err = cursor.All(ctx, &weather); err != nil {
		log.Fatal(err)
	}
	fmt.Println(weather)

}

func updateDatabase(ctx context.Context) {
	database := setupDB(ctx)
	temperatureCollection := database.Collection("Temp")

	//update database

	result, err := temperatureCollection.UpdateOne(
		ctx,
		bson.M{"Period of the day": "Night"},
		bson.D{
			{"$set", bson.D{{"Temperature", "10000°"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v documents \n ", result.ModifiedCount)

}

func deleteFromDatabase(ctx context.Context) {
	database := setupDB(ctx)
	temperatureCollection := database.Collection("Temp")

	// delete from database

	result, err := temperatureCollection.DeleteOne(ctx, bson.M{"Temperature": "10000°"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

}

//Setting up enviorment

func createMongoConnection() *mongo.Client {

	// create connection

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://thomasKuhn:Th0mas123@cluster0.bhvlt.mongodb.net/Cluster0?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)

	return client
}
