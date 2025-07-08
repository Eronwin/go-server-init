package test

import (
	"github.com/gin-gonic/gin"
	v1 "go-server-init/internal/api/v1"
	"go-server-init/internal/repository"
	"go-server-init/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
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
	if !contains(w.Body.String(), `"msg":"pong"`) {
		t.Errorf("Body = %s; want contains msg pong", w.Body.String())
	}

}

func contains(s, sub string) bool {
	return strings.Contains(s, sub)
}
