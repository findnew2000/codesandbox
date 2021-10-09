/*
 * @Description:author
 * @Version: 1.0
 * @Date: 2021-10-09 15:39:24
 * @LastEditTime: 2021-10-10 01:28:20
 */
package v2

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/snmimi/dockapp/models"
)

type LoginInfo struct {
	UserID     string    `json:"userId"`
	ClientIP   string    `json:"clientIP"`
	LoginState string    `json:"loginState"`
	LoginTime  time.Time `json:"loginTime"`
}

func LoginAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user models.User
	user.Username = username
	user.Password = password
	models.Db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("username:%s,password:%s", username, password),
	})
}

func Register(c *gin.Context) {
	var user LoginInfo
	user.ClientIP = "www.snmimi.ml"
	user.LoginState = "ok"
	user.UserID = uuid.NewString()
	user.LoginTime = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	models.Cli.InsertOne(ctx, &user)
}
