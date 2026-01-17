package repository

import (
	"errors"
	"sync"

	"your_project/models"
)

var (
	store = make(map[string]*models.EntityClassification)
	mu    sync.RWMutex
)

func InsertAnalysis(ec *models.EntityClassification) error {
	mu.Lock()
	defer mu.Unlock()

	store[ec.AnalysisID] = ec
	return nil
}

func GetAnalysisByID(analysisID string) (*models.EntityClassification, error) {
	mu.RLock()
	defer mu.RUnlock()

	if val, ok := store[analysisID]; ok {
		return val, nil
	}
	return nil, errors.New("analysis not found")
}
