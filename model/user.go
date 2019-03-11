package model

import (
	"github.com/PedroGao/shoot/libs/enum"
	"github.com/PedroGao/shoot/libs/password"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
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
	// Admin默认为0
	return u.Admin == enum.ADMIN
}

func (u *User) SetPassword(pw string) {
	bts := password.CreatePassword(pw, viper.GetInt("strength"))
	u.Password = string(bts)
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) Verify(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw)) == nil
}

// 数据库操作
func GetUserByName(name string) (*User, error) {
	user := new(User)
	ok, err := Db.Where("nickname = ?", name).Get(user)
	if !ok {
		return nil, err
	}
	return user, nil
}

func CreateUser(nickname, password, email string) error {
	var user = &User{
		Nickname: nickname,
		Email:    email,
	}
	// 默认的email为 ""，故在创建的时候会把其设置为 ""
	// 默认的admin为int，故有默认值为0，会把mysql的默认值1给覆盖
	// todo:
	user.SetPassword(password)
	_, e := Db.InsertOne(user)
	return e
}
