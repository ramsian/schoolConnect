package healthController

import (
	"main/health/healthHelper"

	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (hc *HealthController) GetServerPing(c *gin.Context) {
	helper := new(healthHelper.HealthHelper)
	message := helper.GetServerPing()
	c.String(http.StatusOK, message)
}

func (hc *HealthController) CreateHealthData(c *gin.Context) {
	helper := new(healthHelper.HealthHelper)
	helper.CreateHealthData()
	c.String(http.StatusOK, "Inserted health data!")
}

func (hc *HealthController) Init(r *gin.Engine) {
	r.GET("/getServerPing", hc.GetServerPing)
	r.GET("/createHealthRecord", hc.CreateHealthData)
}
