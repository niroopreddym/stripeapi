package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/niroopreddym/StripeAPI/clients"
	"github.com/niroopreddym/StripeAPI/models"
	"github.com/niroopreddym/StripeAPI/services"
)

// SessionHandler ...
type SessionHandler struct {
	SessionService services.ISessionService
}

// NewSesssionsHandler is the ctor
func NewSesssionsHandler() SessionHandler {
	clientAPIInstance := clients.NewStripeClient()
	return SessionHandler{
		SessionService: services.NewStripeSessionService(clientAPIInstance),
	}
}

// CreateSession start a session for a customer
func (h *SessionHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	sessionPayload := models.StripeSession{}
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		responseController(w, http.StatusInternalServerError, readErr)
		return
	}

	strBufferValue := string(bodyBytes)
	err := json.Unmarshal([]byte(strBufferValue), &sessionPayload)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	responsePayload, err := h.SessionService.CreateNewSession(sessionPayload)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, map[string]string{
		"sessionID": *&responsePayload.ID,
	})
}
