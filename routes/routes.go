package routes

import (
	"main/health/healthController"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	hc := new(healthController.HealthController)
	hc.Init(r)
	return r
}
