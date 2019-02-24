package form

import "errors"

type Login struct {
	NickName string `json:"nickname" validate:"required,min=3" msg:"required:请输入昵称;min:昵称的最小长度为3"`
	Password string `json:"password" validate:"required,min=5" msg:"required:请输入密码;min:密码长度必须大于5"`
	//Email    string `json:"email" form:"email" query:"email" validate:"required,email" msg:"邮箱格式不正确"`
}

func (l Login) ValidateNameAndPassword() error {
	if l.NickName != "pedro" || l.Password != "123456" {
		return errors.New("用户名或密码不正确")
	}
	return nil
}
