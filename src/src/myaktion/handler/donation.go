package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Bermelm/myaktion-go/src/myaktion/model"
	"github.com/Bermelm/myaktion-go/src/myaktion/service"
	log "github.com/sirupsen/logrus"
)

func AddDonation(w http.ResponseWriter, request *http.Request) {
	id, err := getId(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	donation, err := getDonation(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.AddDonation(id, donation)
	//update, err := service.AddDonation(id, donation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//sendJson(w, update)
}

func getDonation(r *http.Request) (*model.Donation, error) {
	var donation model.Donation
	err := json.NewDecoder(r.Body).Decode(&donation)
	if err != nil {
		log.Errorf("Can't serialize request body to donation struct: %v", err)
		return nil, err
	}
	return &donation, nil
}
