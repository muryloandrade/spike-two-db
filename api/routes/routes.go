package routes

import (
	"log"
	"net/http"

	"connect-db/api/controllers"
	"connect-db/api/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleResquests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.HelloWorld).Methods("GET")
	r.HandleFunc("/data1", controllers.GetDatabaseOne).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
