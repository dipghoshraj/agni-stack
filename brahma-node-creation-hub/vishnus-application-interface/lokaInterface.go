package vishnusapplicationinterface

import (
	cosmicmodel "brahma-node-creation-hub/cosmic-model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type NodeRegistrationRequest struct {
	Capacity int64 `json:"capacity"`
	Nodes    int64 `json:"nodes"`
}

func (v *VAI) RegisterLoka(w http.ResponseWriter, r *http.Request) {
	var request NodeRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Error:   fmt.Sprintf("failed to decode request: %v", err),
		})
		return
	}

	volumnName := uuid.New().String()

	loka := &cosmicmodel.LokaNode{
		ID:            uuid.New().String(),
		VolumeName:    volumnName,
		Status:        cosmicmodel.LokaSthitiInactive,
		Capacity:      request.Capacity,
		UsedSpace:     0,
		LastHeartbeat: time.Now(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := v.VAI.BengingLok(loka); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Error:   fmt.Sprintf("failed to register node: %v", err),
		})
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    loka,
	})
}
