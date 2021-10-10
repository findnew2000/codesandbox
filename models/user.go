/*
 * @Description:Table user
 * @Version: 1.0
 * @Date: 2021-10-09 16:07:39
 * @LastEditTime: 2021-10-10 14:13:46
 */
package models

type User struct {
	Uname string `bson:"username" json:"username"`
	Pass  string `bson:"password" json:"password"`
}
