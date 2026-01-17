package routes
import "net/http"
import (
	"github.com/gorilla/mux"

	"your_project/handlers"
	"your_project/middleware"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.Logger)

	r.HandleFunc(
		"/api/entity-classification/analyze",
		handlers.AnalyzeHandler,
	).Methods("POST")

	r.HandleFunc(
		"/api/entity-classification/results/{analysis_id}",
		handlers.GetResultHandler,
	).Methods("GET")
	
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}).Methods("GET")

	return r
}
