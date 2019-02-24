package context

import (
	err2 "github.com/PedroGao/shoot/libs/err"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"strings"
)

type ExtendedContext struct {
	echo.Context
}

func (ctx *ExtendedContext) BindAndValidate(i interface{}) error {
	var (
		err error
	)
	if err = ctx.Bind(i); err != nil {
		return err2.ParamsErr.Set("参数无法获取，请检查参数的内容和格式是否正确")
	}
	if err = ctx.Validate(i); err != nil {
		errs := err.(validator.ValidationErrors)
		if len(errs) > 1 {
			mp := map[string]string{}
			for _, er := range errs {
				prop := er.StructField()
				tag := er.Tag()
				field, ok := reflect.TypeOf(i).Elem().FieldByName(prop)
				if ok {
					// required:请输入昵称;min:昵称的最小长度为3
					totalMsg := field.Tag.Get("msg")
					mp[prop] = getSingleMsg(totalMsg, tag)
				}
			}
			return err2.ParamsErr.Set(mp)
		} else {
			prop := errs[0].StructField()
			tag := errs[0].Tag()
			field, ok := reflect.TypeOf(i).Elem().FieldByName(prop)
			msg := ""
			if ok {
				totalMsg := field.Tag.Get("msg")
				msg = getSingleMsg(totalMsg, tag)
			}
			return err2.ParamsErr.Set(msg)
		}
	}
	return nil
}

func getSingleMsg(totalMsg, tag string) string {
	strs := strings.Split(totalMsg, ";")
	for _, str := range strs {
		tmps := strings.Split(str, ":")
		if tmps[0] == tag {
			return tmps[1]
		}
	}
	return totalMsg
}
