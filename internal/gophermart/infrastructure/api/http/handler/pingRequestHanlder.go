package handler

import (
	"database/sql"
	"net/http"
)

type PingRequestHandler struct {
	db *sql.DB
}

func NewPingRequestHandler(db *sql.DB) *PingRequestHandler {
	return &PingRequestHandler{db: db}
}

func (h *PingRequestHandler) HandleRequest(response http.ResponseWriter, request *http.Request) {
	if err := h.db.Ping(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusOK)
}
