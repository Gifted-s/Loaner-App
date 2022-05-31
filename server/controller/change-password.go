package controller

import (
	"context"
	"encoding/json"
	//"log"
	"net/http"
	"helpful_server/server/helpers"
	"helpful_server/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleChangePassword(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var updatePass models.UpdatePassword

	err := json.NewDecoder(r.Body).Decode(&updatePass)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	filter := bson.M{"phone": updatePass.Phone}
	type fields struct {
		ID       int `bson:"_id"`
		Password int `bson:"password"`
	}
	projection := fields{ID: 0, Password: 1}
	type Password_Struct struct {
		Password string `bson:"password"`
	}
	password := Password_Struct{}
	collections.Users.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&password)
	if !helpers.CheckPasswordHash(updatePass.Password, password.Password) {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "The current password you entered is invalid"}, Token: ""}
		json.NewEncoder(w).Encode(resp)
		return
	}
	hash, _ := helpers.HashPassword(updatePass.NewPassword)
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "password", Value: hash},
		}},
	}
	_, err1 := collections.Users.UpdateOne(context.TODO(), filter, update)
	if err1 != nil {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "Update Failed, Please ensure no field is empty and try again "}}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := models.ResponseStruct{Status: "success", Body: map[string]string{"msg": "Account Update Successful"}}
	json.NewEncoder(w).Encode(resp)
}
