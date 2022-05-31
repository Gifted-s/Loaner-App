package controller

import (
	"encoding/json"
	"log"
	"helpful_server/server/helpers"
	"helpful_server/server/models"

	//"fmt"
	// "github.com/gorilla/websocket"
	"time"
	// "io"
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleProcessLoan(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var requestBody models.RequestBody
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
    requestBody.UserId= id
	result, err := collections.Loan_requests.InsertOne(context.TODO(), requestBody)
	if err != nil {
		// If there is an error in hasing password
		err := models.ResponseStruct{Status: "failed", Body: map[string]string{"error": "Counld Not Insert  Request To DB please try again"}}
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Println(result)
	resp := models.ResponseStruct{Status: "success", Body: requestBody, Token: ""}
	json.NewEncoder(w).Encode(resp)
	helpers.SendMail("Helpful Agent","sunkanmiadewumi1@gmail.com",`Someone needs some cash help`,"<h4>Hello Helpful Agent someone needs some cash help</h4><button style='border:none; padding:20px 10px;background-color:rgb(216, 69, 11);border-radius:6px;color:white;'>View Full Details</button>" )
	// save loan request
	//helpers.SendMail("Helpful Platform","sunkanmiadewumi1@gmail.com","New Loan Request","<h1>New loan request</h1><br/><p>We are always willing to help</p> <a type='btn' href='https://hello.com'>View Rewuest<a/>" )
}
