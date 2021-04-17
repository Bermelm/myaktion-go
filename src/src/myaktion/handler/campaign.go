package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/Bermelm/myaktion-go/src/myaktion/model"
	"github.com/Bermelm/myaktion-go/src/myaktion/service"
)

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign *model.Campaign //hier wichtig, dass es ein pointer ist
	campaign, err := getCampaign(r)
	if err != nil {
		log.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateCampaign(campaign); err != nil { //Service layer, err ist hier nur in diesem Kontext. Hier wird Objekt persistiert
		log.Errorf("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, campaign)
	sendJson(w, result{Success: "OK"})
}

func GetCampaigns(w http.ResponseWriter, _ *http.Request) {
	campaigns, err := service.GetCampaigns()
	if err != nil {
		log.Errorf("Error calling service GetCampaigns: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, campaigns)
}

func GetSingleCampaign(w http.ResponseWriter, request *http.Request) {
	id, err := getId(request)
	if err != nil {
		log.Errorf("Can't parse request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err := service.GetCampaign(id)
	if campaign == nil {
		http.Error(w, "404 campaign not found", http.StatusNotFound)
		return
	}
	sendJson(w, campaign)
}

func UpdateCampaign(w http.ResponseWriter, request *http.Request) {
	id, err := getId(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err := getCampaign(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service.UpdateCampaign(id, campaign)
	//update := service.UpdateCampaign(id, campaign)
	//sendJson(w, update)
}

func DeleteCampaign(w http.ResponseWriter, request *http.Request) {
	id, err := getId(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err := service.DeleteCampaign(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if campaign == nil {
		http.Error(w, "404 campaign not found", http.StatusBadRequest)
		return
	}
	sendJson(w, result{Success: "ok"})
}

// go setzt private func ans Ende!
func getCampaign(r *http.Request) (*model.Campaign, error) {
	var campaign model.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Errorf("Can't serialize request body to campaign struct: %v", err)
		return nil, err
	}
	return &campaign, nil
}
