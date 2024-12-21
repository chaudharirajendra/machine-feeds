package handlers

import (
	"encoding/json"
	"machine-feeds/services"
	"net/http"
)

func MachineFeedHandler(w http.ResponseWriter, r *http.Request) {
	machineID := r.URL.Query().Get("machine_id")
	if machineID == "" {
		http.Error(w, "machine_id query parameter is required", http.StatusBadRequest)
		return
	}

	feed, err := services.PrepareFeed(machineID)
	if err != nil {
		http.Error(w, "Error preparing feed", http.StatusInternalServerError)
		return
	}

	formattedFeed := services.FormatFeed(feed)

	response := map[string]interface{}{
		"machineId": machineID,
		"feeds":     formattedFeed,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
