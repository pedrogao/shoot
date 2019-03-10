package controller

import (
	Err "github.com/PedroGao/shoot/libs/err"
	"github.com/PedroGao/shoot/model"
	"github.com/go-xorm/builder"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

func GetBook(c echo.Context) error {
	var (
		err error
		i   int
	)
	id := c.Param("id")
	i, err = strconv.Atoi(id)
	if err != nil {
		return Err.ParamsErr.Set("参数错误，id必须为正整数")

	}
	book, err := model.GetBookById(i)
	if err != nil {
		return Err.BookNotFound

	}
	return c.JSON(http.StatusOK, book)
}

func SearchBook(c echo.Context) error {
	var (
		err   error
		books []model.Book
	)
	keyword := c.QueryParam("keyword")
	if strings.TrimSpace(keyword) == "" {
		return Err.ParamsErr.Set("参数错误，关键字必须有效词")
	}
	// limit 5 for test
	err = model.Db.Where(builder.Like{"title", keyword}).Limit(5).Find(&books)
	if err != nil {
		return Err.ParamsErr.Set(err.Error())
	}
	if len(books) < 1 {
		return Err.BookNotFound
	}
	return c.JSON(http.StatusOK, books)
}
