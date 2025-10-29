package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilayzen/spy-cat-agency/internal/service"
	"github.com/ilayzen/spy-cat-agency/pkg/models"
	"github.com/ilayzen/spy-cat-agency/pkg/rest"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}

func (h *Handler) FetchCats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cats, err := h.Services.Cats.Fetch(r.Context())
		if err != nil {
			rest.WriteStatusBadRequest(w, err.Error())
			return
		}

		rest.WriteJSON(w, http.StatusOK, cats)
	}
}

func (h *Handler) GetCatByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseUint(params["id"], models.Base, models.BitSize)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		cat, err := h.Services.Cats.GetByID(r.Context(), id)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		rest.WriteJSON(w, http.StatusOK, cat)
	}
}

func (h *Handler) DeleteCatByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseUint(params["id"], models.Base, models.BitSize)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		err = h.Services.Cats.DeleteByID(r.Context(), id)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)

	}
}

func (h *Handler) UpdateCatByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseUint(params["id"], models.Base, models.BitSize)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		salary := models.SalaryRequest{}

		err = json.NewDecoder(r.Body).Decode(&salary)
		if err != nil {
			rest.WriteStatusBadRequest(w, err.Error())
			return
		}

		err = h.Services.Cats.UpdateByID(r.Context(), id, salary.Salary)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) CreateCat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cat := models.Cat{}
		err := json.NewDecoder(r.Body).Decode(&cat)
		if err != nil {
			rest.WriteStatusBadRequest(w, err.Error())
			return
		}

		err = h.Services.Cats.Create(r.Context(), cat)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// missions

func (h *Handler) FetchMissions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := h.Services.Missions.FetchMissions(r.Context())
		if err != nil {
			rest.WriteStatusBadRequest(w, err.Error())
			return
		}

		rest.WriteJSON(w, http.StatusOK, m)
	}
}
func (h *Handler) CreateMission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := models.RequestMission{}
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			rest.WriteStatusBadRequest(w, err.Error())
			return
		}

		err = h.Services.Missions.CreateMission(r.Context(), m)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
func (h *Handler) GetMissionByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseUint(params["id"], models.Base, models.BitSize)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		cat, err := h.Services.Missions.GetMissionByID(r.Context(), id)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		rest.WriteJSON(w, http.StatusOK, cat)
	}
}

func (h *Handler) DeleteMissionByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseUint(params["id"], models.Base, models.BitSize)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		err = h.Services.Missions.DeleteMissionByID(r.Context(), id)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) UpdateMissionByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseUint(params["id"], models.Base, models.BitSize)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		m := models.Mission{}
		err = json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			rest.WriteStatusBadRequest(w, err.Error())
			return
		}

		err = h.Services.Missions.UpdateMissionByID(r.Context(), id, m)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) AddTargetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (h *Handler) UpdateTargetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseUint(params["id"], models.Base, models.BitSize)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		targetID, err := strconv.ParseUint(params["targetId"], models.Base, models.BitSize)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}

		t := models.Target{}
		err = json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			rest.WriteStatusBadRequest(w, err.Error())
			return
		}

		err = h.Services.Target.UpdateByMissionIDAndTargetID(r.Context(), id, targetID, t)
		if err != nil {
			rest.WriteInternalError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)

	}
}
func (h *Handler) DeleteTargetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) AssignCat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
