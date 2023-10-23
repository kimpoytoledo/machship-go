package api

import (
	"machship-go/core/port"
	"machship-go/util/errorhandler"
	"machship-go/util/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase port.GithubUsecase
	log     *logger.Logger
}

func NewHandler(u port.GithubUsecase, log *logger.Logger) *Handler {
	return &Handler{
		Usecase: u,
		log:     log,
	}
}

func (h *Handler) GetGithubUsers(c *gin.Context) {
	var usernamesRequest struct {
		Usernames []string `json:"usernames"`
	}
	if err := c.ShouldBindJSON(&usernamesRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	results, err := h.Usecase.FetchAndCacheGithubUsers(usernamesRequest.Usernames)
	if err != nil {
		responseError := errorhandler.HandleError(h.log, err)
		c.JSON(responseError.Code, gin.H{"error": responseError.Message})
		return
	}

	c.JSON(http.StatusOK, results)
}
