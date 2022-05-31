package controller

import (
	"encoding/json"
	//"log"
	"net/http"

	//"video-chat-app/server/helpers"
	//	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"helpful_server/server/helpers"
	"helpful_server/server/models"

	//"github.com/gorilla/mux"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Otp = map[string]string{"userId": "sunkanmi", "otp": "4322"}

func HandleVerifyOTP(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	type OTP struct {
		UserId string `json:"userId,omitempty" bson:"userId,omitempty"`
		Otp    string `json:"otp,omitempty" bson:"otp,omitempty"`
	}
	var otp OTP
	var otp_marshal_store OTP
	err := json.NewDecoder(r.Body).Decode(&otp)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	id, _ := primitive.ObjectIDFromHex(otp.UserId)
	filter := bson.M{"userId": id}

	e := collections.Otps.FindOne(context.TODO(), filter).Decode(&otp_marshal_store)
	if e != nil {
		resp := models.ResponseStruct{Status: "failed", Body: "could not verify otp"}
		json.NewEncoder(w).Encode(resp)
	}
	if otp.Otp == otp_marshal_store.Otp {
		resp := models.ResponseStruct{Status: "success", Body: "OTP Verifucation successful"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := models.ResponseStruct{Status: "failed", Body: "OTP invalid"}
	json.NewEncoder(w).Encode(resp)

}
