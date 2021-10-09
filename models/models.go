/*
 * @Description:db models
 * @Version: 1.0
 * @Date: 2021-10-08 23:54:12
 * @LastEditTime: 2021-10-10 01:37:03
 */
package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qiniu/qmgo"
	"github.com/snmimi/dockapp/pkg"
)

var Db *gorm.DB
var Cli *qmgo.QmgoClient

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := pkg.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	} else {
		println("db opend")
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	Db.SingularTable(true)
	Db.LogMode(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	Cli, err = qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://n1.lan:27017", Database: "blog", Coll: "user"})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("mongoDB connected")
	}
}
func CloseDB() {
	ctx := context.Background()
	defer Db.Close()
	defer func() {
		err := Cli.Close(ctx)
		if err != nil {
			panic(err)
		}
	}()
}
