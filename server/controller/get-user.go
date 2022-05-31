package controller

import (
	"encoding/json"
	"log"

	"helpful_server/server/helpers"
	"helpful_server/server/models"

	"go.mongodb.org/mongo-driver/bson"

	//"fmt"
	// "github.com/gorilla/websocket"

	// "io"
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var userBody models.SignupStruct
	var loan_requests []models.RequestBody
	var advices []models.TalkStruct
	var talks []models.TalkStruct
	var params = mux.Vars(r)
	log.Println(params)
	// string to primitive.ObjectID
	userid, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": userid}
	loan_filter := bson.M{}
	log.Println((filter))
	collections.Users.FindOne(context.TODO(), filter).Decode(&userBody)

	log.Println(userBody)
	if userBody.Phone == "" {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "user not found"}}
		json.NewEncoder(w).Encode(resp)
		return
	}

	cur, err := collections.Loan_requests.Find(context.TODO(), loan_filter)

	if err != nil {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "error occured while using the software"}}
		json.NewEncoder(w).Encode(resp)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var loan_request models.RequestBody
		// & character returns the memory address of the following variable.
		err := cur.Decode(&loan_request) // decode similar to deserialize process.
		log.Println(loan_request)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array

		loan_requests = append(loan_requests, loan_request)
	}
	advice_body := bson.M{"talkType": "advice"}
	curs, err := collections.Talks.Find(context.TODO(), advice_body)
	if err != nil {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "error occured while using the software"}}
		json.NewEncoder(w).Encode(resp)
		return
	}
	defer curs.Close(context.TODO())

	for curs.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var advice models.TalkStruct
		// & character returns the memory address of the following variable.
		err := curs.Decode(&advice) // decode similar to deserialize process.
		log.Println(advice)
		if err != nil {
			log.Fatal(err)
		}
		// add item our array
		advices = append(advices, advice)
	}

	talk_body := bson.M{"talkType": "talk"}
	cursor, err := collections.Talks.Find(context.TODO(), talk_body)
	if err != nil {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "error occured while using the software"}}
		json.NewEncoder(w).Encode(resp)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var talk models.TalkStruct
		// & character returns the memory address of the following variable.
		err := curs.Decode(&talk) // decode similar to deserialize process.
		log.Println(talk)
		if err != nil {
			log.Fatal(err)
		}
		// add item our array
		talks = append(talks, talk)
	}
	final_response := map[string]interface{}{"user": userBody, "requests": loan_requests, "advices": advices, "talks": talks}
	log.Println(final_response)
	resp := models.ResponseStruct{Status: "success", Body: final_response, Token: ""}
	json.NewEncoder(w).Encode(resp)

	// save loan request
	//helpers.SendMail("Helpful Platform","sunkanmiadewumi1@gmail.com","New Loan Request","<h1>New loan request</h1><br/><p>We are always willing to help</p> <a type='btn' href='https://hello.com'>View Rewuest<a/>" )
}
