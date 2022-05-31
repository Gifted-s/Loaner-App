package controller

import (
	"encoding/json"
	"log"
	//"bytes"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net"
	"os"
	//	"fmt"
	//	"io/ioutil"
	///	"video-chat-app/server/helpers"
	"helpful_server/server/helpers"
	"helpful_server/server/models"

	"github.com/dgrijalva/jwt-go"
	//"fmt"
	// "github.com/gorilla/websocket"
	"math/rand"
	"time"
	// "io"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtKey = []byte("my_secret_key")

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// HandleSignUp : this will handle user signup
func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	collections := helpers.ConnectDB()
	var creds Credentials
	var requestBody models.SignupStruct
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	filter := bson.M{"phone": requestBody.Phone}
	var exist interface{}
	collections.Users.FindOne(context.TODO(), filter).Decode(&exist)
	if exist != nil {
		resp := models.ResponseStruct{Status: "failed", Body: "This phone has been used to register on this platform already.Please go to forgot password or try to Signin"}
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
	requestBody.Password, err = helpers.HashPassword(requestBody.Password)
	if err != nil {
		// If there is an error in hasing password
		err := models.ResponseStruct{Status: "failed", Body: map[string]string{"error": "please enter a new password, the password you entered is not hashable"}, Token: tokenString}
		json.NewEncoder(w).Encode(err)
		return
	}
	requestBody.Token = tokenString
	requestBody.Last_login = currentTime.Format("2006-01-02 3:4:5 pm")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Error Getting IP: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				requestBody.Last_login_IP = ipnet.IP.String()
			}
		}
	}

	result, err := collections.Users.InsertOne(context.TODO(), requestBody)
	if err != nil {
		// If there is an error in hasing password
		err := models.ResponseStruct{Status: "failed", Body: map[string]string{"error": "Counld Not Insert To DB please try again"}}
		json.NewEncoder(w).Encode(err)
		return
	}
	log.Println(result.InsertedID.(primitive.ObjectID).String())
	id := result.InsertedID.(primitive.ObjectID).Hex()
	rand.Seed(time.Now().UnixNano())
	min := 15000
	max := 90000
	otp := rand.Intn(max-min+1) + min
	otpString := strconv.Itoa(otp)
	otpSample := models.OTP{UserId: result.InsertedID.(primitive.ObjectID), Otp: otpString}
	otpInsert, _ := collections.Otps.InsertOne(context.TODO(), otpSample)
	log.Println(otpInsert)
	resp := models.ResponseStruct{Status: "success", Body: requestBody, Token: id}

	json.NewEncoder(w).Encode(resp)
	//if requestBody.Email != "" {
	//helpers.SendMail(requestBody.FullName,requestBody.Email,"Welcome to helpful","<h1>Hello and welcom</h1><br/><p>We are always willing to help</p>" )
	//}
}
