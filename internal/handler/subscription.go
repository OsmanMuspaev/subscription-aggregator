package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/subscription-service/internal/model"
	"github.com/yourusername/subscription-service/internal/repository"
)

type SubscriptionHandler struct {
	Repo *repository.SubscriptionRepo
}

func NewSubscriptionHandler(repo *repository.SubscriptionRepo) *SubscriptionHandler {
	return &SubscriptionHandler{Repo: repo}
}

// POST /subscriptions
func (h *SubscriptionHandler) Create(c *gin.Context) {
	var sub model.Subscription
	if err := c.BindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start, err := time.Parse("2006-01-02", sub.StartDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format, use YYYY-MM-DD"})
		return
	}
	sub.StartDate = start

	if sub.EndDateStr != nil && *sub.EndDateStr != "" {
		end, err := time.Parse("2006-01-02", *sub.EndDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format, use YYYY-MM-DD"})
			return
		}
		sub.EndDate = &end
	}

	if err := h.Repo.Create(&sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sub)
}

// GET /subscriptions
func (h *SubscriptionHandler) List(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")

	subs, err := h.Repo.List(userID, serviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subs)
}

// GET /subscriptions/summary
func (h *SubscriptionHandler) Summary(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")
	fromStr := c.Query("from")
	toStr := c.Query("to")

	from, _ := time.Parse("01-2006", fromStr)
	to, _ := time.Parse("01-2006", toStr)

	sum, err := h.Repo.SumCost(userID, serviceName, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": sum})
}