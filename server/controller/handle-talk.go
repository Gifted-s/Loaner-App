package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"helpful_server/server/helpers"
	"helpful_server/server/models"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleTalk(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var talk models.TalkStruct
	err := json.NewDecoder(r.Body).Decode(&talk)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	currentTime := time.Now()
	Time_Created := currentTime.Format("2006-01-02 3:4:5 pm")
	talk.DateCreated = Time_Created
	var params = mux.Vars(r)
	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])
    talk.UserId= id
	result, err := collections.Talks.InsertOne(context.TODO(), talk)
	if err != nil {
		log.Println(err)
		// If there is an error in hasing password
		err := models.ResponseStruct{Status: "failed", Body: map[string]string{"error": "Cannot process this request now, please try again later"}}
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Println(result)
	resp := models.ResponseStruct{Status: "success", Body: nil}
	json.NewEncoder(w).Encode(resp)
   helpers.SendMail("Helpful Agent","sunkanmiadewumi1@gmail.com",`Someone wants to speak to you`,"<h4>Hello Helpful CEO someone wants to speak to you</h4><button style='border:none; padding:20px 10px;background-color:rgb(216, 69, 11);border-radius:6px;color:white'>View Full Details</button>" )

	// chekck it OTP is Valid

}