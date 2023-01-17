package handlers

import (
	"net/http"
)

// ISessionHandler ...
type ISessionHandler interface {
	CreateSession(w http.ResponseWriter, r *http.Request)
}
