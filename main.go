package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"helpful_server/server"
	"helpful_server/server/controller"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create", server.CreateRoomRequestHandler)
	r.HandleFunc("/join", server.JoinRoomRequestHandler)
	r.HandleFunc("/signup", controller.HandleSignUp).Methods("POST")
	r.HandleFunc("/get-user/{id}", controller.HandleGetUser).Methods("GET")
	r.HandleFunc("/signin", controller.HandleSignIn).Methods("POST")
	r.HandleFunc("/verify-otp", controller.HandleVerifyOTP).Methods("POST")
	r.HandleFunc("/loan-request/{id}", controller.HandleProcessLoan).Methods("POST")
	r.HandleFunc("/edit-account", controller.HandleEditProfile).Methods("POST")
	r.HandleFunc("/change-password", controller.HandleChangePassword).Methods("POST")
	r.HandleFunc("/handle-talk/{id}", controller.HandleTalk).Methods("POST")
	r.HandleFunc("/load-content", controller.HandleLoadContent).Methods("GET")
	r.HandleFunc("/create-blog/{id}", controller.HandleCreateBlog).Methods("POST")

	http.Handle("/", r)
	log.Println("Starting Server on Port 8080")
	err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))
	if err != nil {
		log.Fatal(err)
	}
}
