package v1

import (
	"gf-server/app/api/request"
	"gf-server/app/model/user"
	"gf-server/app/service"
	"gf-server/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func Register(r *ghttp.Request) {
	var (
		err error
		userReturn *user.Entity
		R request.RegisterRequest
	)
	_ = r.Parse(&R)
	if e := gvalid.CheckStruct(R, nil); e != nil {
		g.Dump(e.Maps())
	}
	u := &user.Entity{Username: R.Username, Password: R.Password,Nickname: R.NickName, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	if userReturn, err = service.Register(u); err != nil{
		response.FailWithMessage(r, err.Error())
		r.ExitAll()
	}
	response.OkDetailed(r, userReturn, "注册成功!")
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(r *ghttp.Request) {
	var L request.LoginRequest
	_ = r.Parse(&L)
	if e := gvalid.CheckStruct(L, nil); e != nil {
		g.Dump(e.Maps())
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		u := &user.Entity{Username: L.Username, Password: L.Password}
		userReturn, err := service.Login(u)
		if err != nil {
			response.FailWithMessage(r, err.Error())
		}
		//tokenNext(userReturn)
		response.OkDetailed(r, userReturn, "登录成功!")
	}
}

func tokenNext(u *user.Entity)  {

}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户修改密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func ChangePassword(r *ghttp.Request) {
}

// @Tags SysUser
// @Summary 用户上传头像
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param username formData string true "用户上传头像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadHeaderImg [post]
func UploadHeaderImg(r *ghttp.Request) {
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(r *ghttp.Request) {
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "设置用户权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func SetUserAuthority(r *ghttp.Request) {
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
func DeleteUser(r *ghttp.Request) {
}
