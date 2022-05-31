package controller

import (
	"encoding/json"
	"log"

	//"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	///	"video-chat-app/server/helpers"
	"helpful_server/server/helpers"
	"helpful_server/server/models"

	"github.com/dgrijalva/jwt-go"
	//"fmt"
	// "github.com/gorilla/websocket"

	// "io"
	"context"
	"net/http"
)

// HandleSignUp : this will handle user signup
func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	collections := helpers.ConnectDB()
	var creds Credentials
	var requestBody models.SigninStruct
	var userStruct models.SignupStruct
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	filter := bson.M{"phone": requestBody.Phone}
	e := collections.Users.FindOne(context.TODO(), filter).Decode(&userStruct)
	if e != nil {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "No user with this phone number was found, please register"}, Token: ""}
		json.NewEncoder(w).Encode(resp)
		return
	}
	if !helpers.CheckPasswordHash(requestBody.Password, userStruct.Password) {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "The password you entered was incorrect"}, Token: ""}
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(15 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
 
	resp := models.ResponseStruct{Status: "success", Body: userStruct.ID, Token: tokenString}
	Last_login := currentTime.Format("2006-01-02 3:4:5 pm")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Cound Not get for this user")
	}
	var Last_login_IP string
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				Last_login_IP = ipnet.IP.String()
			}
		}
	}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "last_login", Value: Last_login},
			primitive.E{Key: "last_login_ip", Value: Last_login_IP},
			primitive.E{Key: "token", Value: tokenString},
		}},
	}
	_, res := collections.Users.UpdateOne(context.TODO(), filter, update)

	log.Println(res)
	json.NewEncoder(w).Encode(resp)

	//helpers.SendMail(requestBody.FullName,requestBody.Email,"Welcome to helpful","<h1>Hello and welcom</h1><br/><p>We are always willing to help</p>" )

}
