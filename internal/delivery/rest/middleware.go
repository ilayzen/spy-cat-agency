package delivery

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/ilayzen/spy-cat-agency/pkg/models"
)

func BreedCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		breeds, err := GetAll(r.Context())
		if err != nil {
			http.Error(w, "failed to load cat breeds", http.StatusBadGateway)
			return
		}

		const maxBody = 2 << 20
		body, err := io.ReadAll(io.LimitReader(r.Body, maxBody))
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		_ = r.Body.Close()

		var cat models.Cat
		if len(body) == 0 {
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(body, &cat); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		normalized := strings.ToLower(strings.TrimSpace(cat.Breed))
		if normalized == "" {
			http.Error(w, "breed is required", http.StatusBadRequest)
			return
		}
		if _, ok := breeds[normalized]; !ok {
			http.Error(w, "unknown cat breed: "+cat.Breed, http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewReader(body))

		ctx := context.WithValue(r.Context(), cat.Breed, breeds)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
