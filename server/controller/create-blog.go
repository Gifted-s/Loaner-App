package controller

import (
	"encoding/json"
	"helpful_server/server/helpers"
	"helpful_server/server/models"
	"log"

	//"fmt"
	// "github.com/gorilla/websocket"
	"time"
	// "io"
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HandleCreateBlog this will help us create blog items
func HandleCreateBlog(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var requestBody models.BlogItem
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	log.Println(requestBody)
	currentTime := time.Now()
	Time_Created := currentTime.Format("2006-01-02 3:4:5 pm")
	requestBody.DateCreated = Time_Created
	var params = mux.Vars(r)
	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])
	requestBody.UserId = id
	log.Println(requestBody.Category)
	if requestBody.Category == "audio" {
		result, err := collections.Audios.InsertOne(context.TODO(), requestBody)
		if err != nil {
			// If there is an error in hasing password
			err := models.ResponseStruct{Status: "failed", Body: map[string]string{"error": "Counld Not Insert  Request To DB please try again"}}
			json.NewEncoder(w).Encode(err)
			return
		}
		log.Println(result)
		resp := models.ResponseStruct{Status: "success", Body: requestBody, Token: ""}
		json.NewEncoder(w).Encode(resp)
		return
	}
	if requestBody.Category == "video" {
		result, err := collections.Videos.InsertOne(context.TODO(), requestBody)
		if err != nil {
			// If there is an error in hasing password
			err := models.ResponseStruct{Status: "failed", Body: map[string]string{"error": "Counld Not Insert  Request To DB please try again"}}
			json.NewEncoder(w).Encode(err)
			return
		}
		log.Println(result)
		resp := models.ResponseStruct{Status: "success", Body: requestBody, Token: ""}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "item was not inserted"}, Token: ""}
	json.NewEncoder(w).Encode(resp)

	// save loan request
	//helpers.SendMail("Helpful Platform","sunkanmiadewumi1@gmail.com","New Loan Request","<h1>New loan request</h1><br/><p>We are always willing to help</p> <a type='btn' href='https://hello.com'>View Rewuest<a/>" )
}
