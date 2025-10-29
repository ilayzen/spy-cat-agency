package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ilayzen/spy-cat-agency/pkg/models"
)

func WriteNotFound(w http.ResponseWriter, resourceName string) {
	WriteJSON(w, http.StatusNotFound, &models.ErrorResponse{
		Code:    http.StatusNotFound,
		Message: resourceName + " not found",
	})
}

func WriteInternalError(w http.ResponseWriter, msg string) {
	WriteJSON(w, http.StatusInternalServerError, &models.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: msg,
	})
}

func WriteStatusBadRequest(w http.ResponseWriter, msg string) {
	WriteJSON(w, http.StatusBadRequest, &models.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: msg,
	})
}

func WriteStatusUnprocessableEntity(w http.ResponseWriter, msg string) {
	WriteJSON(w, http.StatusUnprocessableEntity, &models.ErrorResponse{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
	})
}

func WriteJSON(w http.ResponseWriter, httpCode int, msg interface{}) {
	b, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, fmt.Errorf("cannot return response, err: %s", err).Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to write response, err: %s", err).Error(), http.StatusInternalServerError)
		return
	}
}

func WriteUnauthorized(w http.ResponseWriter) {
	WriteJSON(w, http.StatusUnauthorized, &models.ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized request",
	})
}

func WriteForbidden(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, &models.ErrorResponse{
		Code:    http.StatusForbidden,
		Message: "Request is forbidden",
	})
}

func WriteStatusConflict(w http.ResponseWriter, msg string) {
	WriteJSON(w, http.StatusConflict, &models.ErrorResponse{
		Code:    http.StatusConflict,
		Message: msg,
	})
}

func WriteLimited(w http.ResponseWriter) {
	WriteJSON(w, http.StatusServiceUnavailable, &models.ErrorResponse{
		Code:    http.StatusServiceUnavailable,
		Message: "Rate limit exceeded",
	})
}
