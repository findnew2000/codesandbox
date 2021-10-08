/*
 * @Description:main
 * @Version: 1.0
 * @Date: 2021-10-08 21:18:07
 * @LastEditTime: 2021-10-09 00:29:17
 */
package main

import (
	"fmt"
	"net/http"

	"gitee.com/findnew/dockapp/pkg"
	"gitee.com/findnew/dockapp/routers"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", pkg.HTTPPort),
		Handler:        router,
		ReadTimeout:    pkg.ReadTimeout,
		WriteTimeout:   pkg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
