package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckRoute(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status":"ok"})
}