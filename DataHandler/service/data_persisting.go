package service

import (
	"context"
	"fmt"
	"github.com/SanjanaKansal/data_handler/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// PersistData, takes the data and stores it in MongoDB.
func PersistData(scrapedData models.ScrapedData) models.SuccessMessage {
	var out models.SuccessMessage
	scrapedData.LastUpdatedTime = time.Now()

	// Make a client
	client, err := mongo.NewClient(options.Client().ApplyURI(models.MONGOURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// Connect to the client
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		out.Success = false
		out.Message = err.Error()
		return out
	}
	fmt.Println("Connected to MongoDB!")
	defer client.Disconnect(ctx)

	// Access a MongoDB collection through a database
	collection := client.Database(models.MONGODBNAME).Collection(models.MONGOCOLLECTIONNAME)
	

	// Update document in Mongo.
	filter := bson.M{"url": scrapedData.URL}
	update := bson.M{
		"$set" : bson.M{
			"product": scrapedData.Product,
			"lastupdatedtime": scrapedData.LastUpdatedTime,
		},
	}
	updated, _ := collection.UpdateOne(ctx, filter, update)
	if updated.ModifiedCount == 0 {
		_, err := collection.InsertOne(ctx, scrapedData)
		if err != nil {
			log.Println(err)
			out.Success = false
			out.Message = err.Error()
			return out
		}
	}
	log.Println("Successfully pushed document to Mongo")
	out.Success = true
	out.Message = "Successfully pushed document to Mongo"
	return out
}
