package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type breed struct {
	Name string `json:"name"`
}

var (
	mu         sync.RWMutex
	cache      map[string]struct{}
	expiresAt  time.Time
	ttl        = 10 * time.Minute
	apiKey     string
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

func SetAPIKey(k string) {
	mu.Lock()
	apiKey = k
	mu.Unlock()
}

func GetAll(ctx context.Context) (map[string]struct{}, error) {
	mu.RLock()
	if time.Now().Before(expiresAt) && len(cache) > 0 {
		defer mu.RUnlock()
		copy := make(map[string]struct{}, len(cache))
		for k := range cache {
			copy[k] = struct{}{}
		}
		return copy, nil
	}
	mu.RUnlock()

	names, err := fetchAll(ctx)
	if err != nil {
		return nil, err
	}

	mu.Lock()
	cache = names
	expiresAt = time.Now().Add(ttl)
	mu.Unlock()

	return names, nil
}

func Validate(ctx context.Context, breedName string) (bool, error) {
	all, err := GetAll(ctx)
	if err != nil {
		return false, err
	}
	_, ok := all[strings.ToLower(strings.TrimSpace(breedName))]
	return ok, nil
}

func fetchAll(ctx context.Context) (map[string]struct{}, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.thecatapi.com/v1/breeds", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	mu.RLock()
	key := apiKey
	mu.RUnlock()
	if key != "" {
		req.Header.Set("x-api-key", key)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 256))
		return nil, fmt.Errorf("catapi status %d: %s", resp.StatusCode, string(b))
	}

	var breeds []breed
	if err := json.NewDecoder(io.LimitReader(resp.Body, 2<<20)).Decode(&breeds); err != nil {
		return nil, fmt.Errorf("decode breeds: %w", err)
	}

	out := make(map[string]struct{}, len(breeds))
	for _, b := range breeds {
		if b.Name != "" {
			out[strings.ToLower(b.Name)] = struct{}{}
		}
	}
	return out, nil
}
