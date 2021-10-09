/*
 * @Description:router
 * @Version: 1.0
 * @Date: 2021-10-09 00:27:21
 * @LastEditTime: 2021-10-09 15:17:22
 */
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/snmimi/dockapp/pkg"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(pkg.RunMode)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	return r
}
