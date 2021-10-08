/*
 * @Description:main
 * @Version: 1.0
 * @Date: 2021-10-08 21:18:07
 * @LastEditTime: 2021-10-08 23:26:34
 */
package main

import (
	"net/http"

	"gitee.com/findnew/dockapp/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	r.Run(pkg.HTTPPort)

}
