/*
 * @Description:db models
 * @Version: 1.0
 * @Date: 2021-10-08 23:54:12
 * @LastEditTime: 2021-10-10 14:56:19
 */
package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/qiniu/qmgo"
	"github.com/snmimi/dockapp/pkg"
)

var Cli *qmgo.QmgoClient

func init() {
	var (
		err                  error
		dbName, host, tbName string
	)
	sec, err := pkg.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbName = sec.Key("NAME").String()
	host = sec.Key("HOST").String()
	tbName = sec.Key("TABLE_PREFIX").String()
	// mongodb connection
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	Cli, err = qmgo.Open(ctx, &qmgo.Config{Uri: host, Database: dbName, Coll: tbName})
	if err != nil {
		panic(err)
	}
	if Cli.Ping(2) != nil {
		fmt.Println("ping mongodb failed")
	}
}

func CloseDB() {
	defer func() {
		ctx := context.Background()
		if err := Cli.Close(ctx); err != nil {
			panic(err)
		}
	}()
}
