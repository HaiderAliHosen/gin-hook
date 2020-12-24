package hook

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ihandlers interface {
	createHookHandler(c *gin.Context)
}

type handler struct {
	hookService *Service
}

//MakeHTTPHandlers ---
func MakeHTTPHandlers(router *gin.RouterGroup, hookService *Service) {
	h := handler{
		hookService: hookService,
	}
	router.POST("hook", h.createHookHandler)
}

type hookRequest struct {
	hook Hook
}

type hookResponse struct {
	Hook Hook   `json:"hook"`
	Err  string `json:"err"`
}

func (h1 *handler) createHookHandler(c *gin.Context) {
	var req hookRequest

	if err := c.ShouldBindJSON(&req.hook); err != nil {
		c.JSON(http.StatusInternalServerError, hookResponse{
			Hook: Hook{},
			Err:  "data binding error",
		})
		return
	}
	postResults, err := h1.hookService.CreateHookService(req.hook)

	if err != nil {
		c.JSON(http.StatusInternalServerError, hookResponse{
			Hook: Hook{},
			Err:  "data binding error",
		})
		return
	}
	c.JSON(http.StatusOK, hookResponse{
		Hook: postResults,
		Err:  "",
	})
	return
}
