package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	mauth "github.com/zonradkuse/oauth-authenticator"
)

const (
	LOCAL_PATH_STATIC    = "./static/"
	LOCAL_PATH_TEMPLATES = "./templates/"
	WEB_PATH_STATIC      = "/static/"
	WEB_PATH_LOGIN       = "/oauth/authorize"
	WEB_PATH_TOKEN       = "/oauth/token"
	WEB_PATH_INFO        = "/api/v4/user"
)

func startServer(server *mauth.OAuthServer) {
	log.Println("Starting Webservice...")
	r := mux.NewRouter()
	r.PathPrefix(WEB_PATH_STATIC).Handler(http.StripPrefix(WEB_PATH_STATIC, http.FileServer(http.Dir(LOCAL_PATH_STATIC))))
	r.HandleFunc(WEB_PATH_LOGIN, handleLoginLanding).Methods("GET")
	r.HandleFunc(WEB_PATH_LOGIN, server.HandleAuthorizeRequest).Methods("POST")
	r.HandleFunc(WEB_PATH_TOKEN, server.HandleTokenRequest).Methods("POST")
	r.HandleFunc(WEB_PATH_INFO, server.HandleUserInfoRequest).Methods("GET")

	// Start http server
	log.Println("Listening on :3000")
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":3000", loggedRouter)
}

func handleLoginLanding(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html")
}
