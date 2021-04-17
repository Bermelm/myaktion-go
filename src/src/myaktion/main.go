package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/Bermelm/myaktion-go/src/myaktion/handler"
	"github.com/gorilla/mux"
)

func init() { //init wird nur 1mal ausgeführt
	//defer db.init()
	// init logger
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Info("Log level not specified, set default to: INFO")
		log.SetLevel(log.InfoLevel)
		return
	}
	log.SetLevel(level)
}

func main() {
	log.Println("Starting My-Aktion API server")
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods("GET") //hier wird eine Funktion angemeldet. Hier ist es handler.health
	router.HandleFunc("/campaign", handler.CreateCampaign).Methods("POST")
	router.HandleFunc("/campaigns", handler.GetCampaigns).Methods("GET")
	router.HandleFunc("/campaigns/{id}", handler.GetSingleCampaign).Methods("GET")
	router.HandleFunc("/campaigns/{id}", handler.UpdateCampaign).Methods("PUT")
	router.HandleFunc("/campaigns/{id}", handler.DeleteCampaign).Methods("DELETE")
	router.HandleFunc("/campaigns/{id}/donation", handler.AddDonation).Methods("POST")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
