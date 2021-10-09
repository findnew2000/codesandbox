/*
 * @Description:main
 * @Version: 1.0
 * @Date: 2021-10-08 21:18:07
 * @LastEditTime: 2021-10-09 15:16:53
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/snmimi/dockapp/pkg"
	"github.com/snmimi/dockapp/routers"
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
