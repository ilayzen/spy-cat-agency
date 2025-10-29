package delivery

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(h *Handler) *mux.Router {
	router := mux.NewRouter()

	apiV1 := router.PathPrefix("/api/v1").Subrouter()

	secured := router.PathPrefix("/api/v1").Subrouter()
	secured.Use(BreedCacheMiddleware)

	apiV1.HandleFunc("/cats", h.FetchCats()).Methods(http.MethodGet)
	secured.HandleFunc("/cats", h.CreateCat()).Methods(http.MethodPost)
	apiV1.HandleFunc("/cats/{id}", h.GetCatByID()).Methods(http.MethodGet)
	apiV1.HandleFunc("/cats/{id}", h.DeleteCatByID()).Methods(http.MethodDelete)
	apiV1.HandleFunc("/cats/{id}", h.UpdateCatByID()).Methods(http.MethodPut)

	apiV1.HandleFunc("/missions", h.FetchMissions()).Methods(http.MethodGet)
	apiV1.HandleFunc("/missions", h.CreateMission()).Methods(http.MethodPost)
	apiV1.HandleFunc("/missions/{id}", h.GetMissionByID()).Methods(http.MethodGet)
	apiV1.HandleFunc("/missions/{id}", h.DeleteMissionByID()).Methods(http.MethodDelete)
	apiV1.HandleFunc("/missions/{id}", h.UpdateMissionByID()).Methods(http.MethodPut)

	apiV1.HandleFunc("/missions/{id}/targets", h.AddTargetByID()).Methods(http.MethodPost)
	apiV1.HandleFunc("/missions/{id}/targets/{targetId}", h.UpdateTargetByID()).Methods(http.MethodPut)
	apiV1.HandleFunc("/missions/{id}/targets/{targetId}", h.DeleteTargetByID()).Methods(http.MethodDelete)

	apiV1.HandleFunc("/missions/{id}/assign", h.AssignCat()).Methods(http.MethodPost)

	return router
}
