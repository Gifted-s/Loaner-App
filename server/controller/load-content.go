package controller
import (
	"encoding/json"
		///	"video-chat-app/server/helpers"
	//"video-chat-app/server/helpers"
	"helpful_server/server/models"
	//"fmt"
	// "github.com/gorilla/websocket"

	// "io"
	"net/http"
)


func HandleLoadContent (w http.ResponseWriter, r *http.Request) {
	
	// var contents models.ContentsStruct
	// err := json.NewDecoder(r.Body).Decode(&requestBody)
	// if err != nil {
	// 	json.NewEncoder(w).Encode(err)
	// }
    resp := models.ResponseStruct{Status: "success", Body: nil, Token:""}
	json.NewEncoder(w).Encode(resp)
	// save loan request
	//helpers.SendMail("Helpful Platform","sunkanmiadewumi1@gmail.com","New Loan Request","<h1>New loan request</h1><br/><p>We are always willing to help</p> <a type='btn' href='https://hello.com'>View Rewuest<a/>" )
}