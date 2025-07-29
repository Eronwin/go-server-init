package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "go-server-init/internal/api/v1"
	"go-server-init/internal/repository"
	"go-server-init/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	repo := repository.New()
	svc := service.New(repo)
	r := gin.New()
	r.GET("/ping", v1.NewPingHandler(svc).Ping)

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ping: expected status 200 got %d", w.Code)
	}

	type PingResp struct {
		Data struct {
			Message string `json:"message"`
			UUID    string `json:"uuid"`
		} `json:"data"`
	}

	var resp PingResp
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	if resp.Data.Message != "pong" {
		t.Errorf("expected message=pong, got %s", resp.Data.Message)
	}
}
