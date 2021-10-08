/*
 * @Description:main
 * @Version: 1.0
 * @Date: 2021-10-08 21:18:07
 * @LastEditTime: 2021-10-09 03:21:36
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/findnew2000/dockapp/pkg"
	"github.com/findnew2000/dockapp/routers"
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
