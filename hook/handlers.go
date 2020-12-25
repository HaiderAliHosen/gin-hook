package hook

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ihandlers interface {
	createHookHandler(c *gin.Context)
	getAllHookHandlers(c *gin.Context)
	getSingleHookHandler(c *gin.Context)
	deleteHandler(c *gin.Context)
	updateHandler(c *gin.Context)
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
	router.GET("hook", h.getAllHookHandlers)
	router.GET("hook/:id", h.getSingleHookHandler)
	router.DELETE("hook/:id", h.deleteHandler)
	router.PUT("hook/:id", h.updateHandler)
}

type hookRequest struct {
	hook Hook
}

type createHookResponse struct {
	Hook Hook   `json:"hook"`
	Err  string `json:"err"`
}

func (h1 *handler) createHookHandler(c *gin.Context) {
	var req hookRequest

	if err := c.ShouldBindJSON(&req.hook); err != nil {
		c.JSON(http.StatusInternalServerError, createHookResponse{
			Hook: Hook{},
			Err:  "data binding error",
		})
		return
	}
	postResults, err := h1.hookService.CreateHookService(req.hook)

	if err != nil {
		c.JSON(http.StatusInternalServerError, createHookResponse{
			Hook: Hook{},
			Err:  "data binding error",
		})
		return
	}
	c.JSON(http.StatusOK, createHookResponse{
		Hook: postResults,
		Err:  "",
	})
	return
}

type readAllHookResponse struct {
	Hook []Hook `json:"hooks"`
	Err  string `json:"err"`
}

func (h1 *handler) getAllHookHandlers(c *gin.Context) {
	getResults, err := h1.hookService.GetAllHookService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, readAllHookResponse{
			Hook: []Hook{},
			Err:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, readAllHookResponse{
		Hook: getResults,
		Err:  "",
	})
}

type getSingleHookHandlersResponse struct {
	Hook Hook   `json:"_id"`
	Err  string `json:"err"`
}

func (h1 *handler) getSingleHookHandler(c *gin.Context) {
	hookID := c.Param("id")
	getSingleResults, err := h1.hookService.GetSingleHookService(hookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, getSingleHookHandlersResponse{
			Hook: Hook{},
			Err:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, getSingleHookHandlersResponse{
		Hook: getSingleResults,
		Err:  "",
	})
}

type deleteHandlerResponse struct {
	Err string `json:"err"`
}

func (h1 *handler) deleteHandler(c *gin.Context) {
	hookID := c.Param("id")
	err := h1.hookService.deleteHookService(hookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, deleteHandlerResponse{
			Err: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, deleteHandlerResponse{
		Err: "",
	})
	return
}

type updateRequest struct {
	Hook Hook
}
type updateHookHandlerResponse struct {
	Hook Hook   `json:"hook"`
	Err  string `json:"err"`
}

func (h1 *handler) updateHandler(c *gin.Context) {
	var req updateRequest
	hookID := c.Param("id")
	if err := c.ShouldBindJSON(&req.Hook); err != nil {
		c.JSON(http.StatusInternalServerError, updateHookHandlerResponse{
			Hook: Hook{},
			Err:  err.Error(),
		})
		return
	}

	updatehanler, err := h1.hookService.updateHookService(hookID, req.Hook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, updateHookHandlerResponse{
			Hook: Hook{},
			Err:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updateHookHandlerResponse{
		Hook: updatehanler,
		Err:  "",
	})
	return
}
