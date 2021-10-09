/*
 * @Description:router
 * @Version: 1.0
 * @Date: 2021-10-09 00:27:21
 * @LastEditTime: 2021-10-09 23:11:56
 */
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/snmimi/dockapp/pkg"
	v2 "github.com/snmimi/dockapp/routers/v2"
)

func InitRouter() *gin.Engine {
	gin.SetMode(pkg.RunMode)
	r := gin.Default()

	apiv2 := r.Group("/v2")
	{
		auth := apiv2.Group("/auth")
		{
			auth.POST("/login", v2.LoginAuth)
			auth.POST("/register", v2.Register)
		}

	}

	return r
}
