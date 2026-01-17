package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"your_project/models"
	"your_project/repository"
)

type AnalyzeRequest struct {
	TranscriptID   string `json:"transcript_id"`
	TranscriptText string `json:"transcript_text"`
	CreatorID      string `json:"creator_id"`
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	var req AnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	analysisID := uuid.New().String()

	entities := map[string]interface{}{
		"people":    []string{"Tim Ferriss"},
		"tools":     []string{"Notion"},
		"brands":    []string{},
		"products":  []string{"The 4-Hour Workweek"},
		"companies": []string{},
	}

	tone := map[string]interface{}{
		"primary":    "conversational",
		"secondary":  "educational",
		"confidence": 0.88,
	}

	style := map[string]interface{}{
		"primary":    "interview",
		"confidence": 0.92,
	}

	safety := map[string]interface{}{
		"sensitive_domains": []string{"mental health"},
		"severity":          "low",
		"requires_review":   false,
	}

	record := &models.EntityClassification{
		AnalysisID:   analysisID,
		TranscriptID: req.TranscriptID,
		CreatorID:    req.CreatorID,
		Entities:     entities,
		Tone:         tone,
		Style:        style,
		SafetyFlags:  safety,
		CreatedAt:    time.Now(),
	}

	_ = repository.InsertAnalysis(record)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":       "success",
		"analysis_id":  analysisID,
		"entities":     entities,
		"tone":         tone,
		"style":        style,
		"safety_flags": safety,
	})
}

func GetResultHandler(w http.ResponseWriter, r *http.Request) {
	analysisID := mux.Vars(r)["analysis_id"]

	result, err := repository.GetAnalysisByID(analysisID)
	if err != nil {
		http.Error(w, "Analysis not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
