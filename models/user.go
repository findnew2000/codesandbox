/*
 * @Description:Table user
 * @Version: 1.0
 * @Date: 2021-10-09 16:07:39
 * @LastEditTime: 2021-10-09 23:45:19
 */
package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func init() {

	Db.AutoMigrate(&User{})
}
func (user *User) BeforeCreate(scope *gorm.Scope) (err error) {
	// UUID version 4
	// user.ID, _ = uuid.NewUUID()
	UUID, _ := uuid.NewUUID()
	scope.SetColumn("ID", UUID)
	return
}
