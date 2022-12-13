package healthz

import "github.com/gin-gonic/gin"

type IHandler interface {
	Healthz(c *gin.Context)
}
