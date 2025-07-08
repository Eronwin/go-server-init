package v1

import (
	"github.com/gin-gonic/gin"
	"go-server-init/internal/service"
	"net/http"
)

type PingHandler struct {
	svc *service.Service
}

func NewPingHandler(svc *service.Service) *PingHandler {
	return &PingHandler{svc: svc}
}

//

func (h *PingHandler) Ping(c *gin.Context) {
	res, err := h.svc.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
