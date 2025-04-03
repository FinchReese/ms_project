package user

import (
	"errors"

	common "test.com/project-common"
)

type RegisterReq struct {
	Email     string `form:"email"`
	Name      string `form:"name"`
	Password  string `form:"password"`
	Password2 string `form:"password2"`
	Mobile    string `form:"mobile"`
	Captcha   string `form:"captcha"`
}

func (r RegisterReq) VerifyPassword() bool {
	return r.Password == r.Password2
}

func (r RegisterReq) Verify() error {
	if !common.VerifyEmailFormat(r.Email) {
		return errors.New("邮箱格式不正确")
	}
	if !common.VerifyMobile(r.Mobile) {
		return errors.New("手机号格式不正确")
	}
	if !r.VerifyPassword() {
		return errors.New("两次密码输入不一致")
	}
	return nil
}
