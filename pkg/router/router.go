package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewRouter() {
	//connection := db.Open()
	//connection.SyncSchema()

	router := mux.NewRouter()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	router.HandleFunc("/health", health).Methods(http.MethodGet)

	port := ":8080"
	log.Fatal(http.ListenAndServe(port, router))
}

func health(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("ok")
	if err != nil {
		panic(err)
	}
}
