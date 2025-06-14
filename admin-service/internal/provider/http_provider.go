package provider

import (
	"github.com/gin-gonic/gin"
)

func BootstrapHttp(router *gin.Engine) {
	BootstrapApp(router)
	BootstrapAdmin(router)
}
