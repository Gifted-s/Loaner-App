package controller

import (
	"encoding/json"
	"net"
	"net/http"
	"time"

	"log"
	"helpful_server/server/helpers"
	"helpful_server/server/models"
    "context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleEditProfile(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	collections := helpers.ConnectDB()
	var editProf models.EditProfile;
	err := json.NewDecoder(r.Body).Decode(&editProf)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
    log.Println(editProf)
	filter := bson.M{"phone": editProf.Phone}
	Last_login := currentTime.Format("2006-01-02 3:4:5 pm")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Cound Not get for this user")
	}
    var Last_login_IP  string;
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				Last_login_IP = ipnet.IP.String()
			}
		}
	}
    update := bson.D{
		primitive.E{Key:"$set",Value :bson.D{
			primitive.E{Key:"email", Value:editProf.Email},
			primitive.E{Key:"fullName", Value:editProf.FullName},
			primitive.E{Key:"address", Value:editProf.Address},
			primitive.E{Key:"last_login_ip",Value: Last_login_IP},
			primitive.E{Key:"last_login",Value: Last_login},
		}},
	}
	_,err1 := collections.Users.UpdateOne(context.TODO(), filter, update)
	if err1 != nil {
		log.Println(err1)
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg":"Update Failed, Please ensure no field is empty and try again "}}
	   json.NewEncoder(w).Encode(resp)
	   return
	}
	
	resp := models.ResponseStruct{Status: "success", Body: map[string]string{"msg":"Account Update Successful"}, }
	json.NewEncoder(w).Encode(resp)
	// chekck it OTP is Vali
}
