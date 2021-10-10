/*
 * @Description:author
 * @Version: 1.0
 * @Date: 2021-10-09 15:39:24
 * @LastEditTime: 2021-10-10 14:52:58
 */
package v2

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snmimi/dockapp/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := models.Cli.Find(ctx, bson.M{"username": username}).One(&user)
	if err != nil {
		fmt.Println(err)
	}
	if user.Uname != username {
		user.Uname = username
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}
		user.Pass = string(hash)
		models.Cli.InsertOne(ctx, &user)
	}
}

func LoginAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := models.Cli.Find(ctx, bson.M{"username": username}).One(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "未注册",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "密码错",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
