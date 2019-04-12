package views

import "log"

var msgText  map[string]string


func init(){
	msgText = make(map[string]string)
	msgText["success"] = "success"
	msgText["failed"] = "failed"
	msgText["register_success"] = "注册成功"
	msgText["register_failed"] = "注册失败，用户名已存在或信息填写不规范"
	msgText["login_success"] = "登陆成功"
	msgText["login_failed"] = "登陆失败，请检查用户名或密码"
	msgText["modify_success"] = "修改信息成功"
	msgText["modify_failed"] = "修改个人信息失败，请检查用户是否存在"

	log.Println("text const init success")
}

