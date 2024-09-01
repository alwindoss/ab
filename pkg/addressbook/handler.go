package addressbook

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler interface {
	CreateContacts(w http.ResponseWriter, r *http.Request)
}

func NewHandler(svc Service) Handler {
	return &abHandler{
		svc: svc,
	}
}

type abHandler struct {
	svc Service
}

type createContactsRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address,omitempty"`
}

// CreateContacts implements Handler.
func (a *abHandler) CreateContacts(w http.ResponseWriter, r *http.Request) {
	ccrs := []createContactsRequest{}
	err := json.NewDecoder(r.Body).Decode(&ccrs)
	if err != nil {
		http.Error(w, "unable to decode request", http.StatusBadRequest)
		return
	}
	for _, c := range ccrs {
		fmt.Println("First Name:", c.FirstName)
	}
}
