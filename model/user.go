package model

import (
	"github.com/PedroGao/shoot/libs/enum"
	"sync"
	"time"
)

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[int]*User
}

type User struct {
	Id       int    `xorm:"not null pk autoincr INT(11)" json:"id"`
	Nickname string `xorm:"not null VARCHAR(50)" json:"nickname"`
	Email    string `xorm:"VARCHAR(30)" json:"email"`
	Password string `xorm:"not null VARCHAR(1000)" json:"-"`
	Admin    int    `xorm:"TINYINT default 1" json:"admin"` // admin 1 表示普通用户 2 表示超级管理员

	CreateTime time.Time `xorm:"created" json:"create_time"`
	UpdateTime time.Time `xorm:"updated" json:"update_time"`
	DeleteTime time.Time `xorm:"deleted" json:"delete_time"`
}

func (u *User) IsAdmin() bool {
	return u.Admin == enum.ADMIN
}

func SetPassword(password string) {

}

func GetPassword() string {
	return ""
}

func Verify() {

}
