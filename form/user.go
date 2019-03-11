package form

import (
	err2 "github.com/PedroGao/shoot/libs/err"
	"github.com/PedroGao/shoot/model"
)

type Login struct {
	NickName string `json:"nickname" validate:"required,min=3" msg:"required:请输入昵称;min:昵称的最小长度为3"`
	Password string `json:"password" validate:"required,min=5" msg:"required:请输入密码;min:密码长度必须大于5"`
	//Email    string `json:"email" form:"email" query:"email" validate:"required,email" msg:"邮箱格式不正确"`
}

func (l Login) ValidateNameAndPassword() error {
	user, _ := model.GetUserByName(l.NickName)
	if user == nil {
		return err2.NotFound.Set("用户不存在")
	}
	if !user.Verify(l.Password) {
		return err2.ParamsErr.Set("输入密码错误")
	}
	return nil
}

type Register struct {
	NickName        string `validate:"required,min=3" msg:"required:请输入昵称;min:昵称的最小长度为3" json:"nickname"`
	Password        string `validate:"required,min=5" msg:"required:请输入密码;min:密码长度必须大于5" json:"password"`
	ConfirmPassword string `validate:"required,eqfield=Password" msg:"required:请输入确认密码;eqfield=Password:两次输入密码不一致" json:"confirm_password"`
	Email           string `validate:"omitempty,email" msg:"email:输入邮箱不符合规范" json:"email"`
}
