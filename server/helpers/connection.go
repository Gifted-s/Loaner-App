package helpers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Collections struct {
	Articles      *mongo.Collection
	Audios        *mongo.Collection
	Users         *mongo.Collection
	Talks         *mongo.Collection
	Otps          *mongo.Collection
	Loan_requests *mongo.Collection
	Videos        *mongo.Collection
}

func ConnectDB() Collections {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://Sunkanmi:sunkanmi123@cluster0.bpei9.mongodb.net/helpful?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	users := client.Database("helpful_db").Collection("users")
	audios := client.Database("helpful_db").Collection("audios")
	videos := client.Database("helpful_db").Collection("videos")
	articles := client.Database("helpful_db").Collection("articles")
	talks := client.Database("helpful_db").Collection("talks")
	otps := client.Database("helpful_db").Collection("otps")
	loan_requests := client.Database("helpful_db").Collection("loan_requests")

	var collections Collections
	collections.Articles = articles
	collections.Videos = videos
	collections.Audios = audios
	collections.Users = users
	collections.Loan_requests = loan_requests
	collections.Talks = talks
	collections.Otps = otps
	return collections
}
