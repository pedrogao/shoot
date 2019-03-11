package controller

import (
	"fmt"
	"github.com/PedroGao/shoot/form"
	"github.com/PedroGao/shoot/libs/context"
	Err "github.com/PedroGao/shoot/libs/err"
	"github.com/PedroGao/shoot/libs/token"
	"github.com/PedroGao/shoot/model"
	"github.com/PedroGao/shoot/service"
	"github.com/labstack/echo"
	"net/http"
)

func Login(c echo.Context) error {
	var (
		login                     form.Login
		accessToken, refreshToken string
		err                       error
	)
	cc := c.(context.ExtendedContext)

	if err := cc.BindAndValidate(&login); err != nil {
		// 返回统一的错误
		return err
	}

	if err = login.ValidateNameAndPassword(); err != nil {
		// 返回统一的错误
		return err
	}

	accessToken, refreshToken, err = token.JwtInstance.GenerateTokens(login.NickName)
	// err为默认的错误，故将其转为统一的错误
	if err != nil {
		return Err.ParamsErr.Set(err.Error())
	}
	return cc.JSON(http.StatusOK, echo.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func Register(c echo.Context) error {
	var (
		register form.Register
		err      error
	)
	cc := c.(context.ExtendedContext)

	if err = cc.BindAndValidate(&register); err != nil {
		return err
	}
	model.CreateUser(register.NickName, register.Password, register.Email)
	//if err != nil {
	//}
	return Err.OK
}

func GetUsers(c echo.Context) error {
	infos, e := service.ListUser()
	value := c.Get("user")
	if value != nil {
		fmt.Println(value)
	}
	if e != nil {
		return Err.UserNotFound
	}
	return c.JSON(http.StatusOK, infos)
}
