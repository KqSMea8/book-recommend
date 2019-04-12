package dal

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Info string `gorm:"column:info"`
}


//设置用户的某个属性
func SetUser(username string,key string, value string){
	var user User
	db_conn.First(&user, "username = ?", username)
	db_conn.Model(&user).Update(key, value)
}

//创建新用户
func CreateUser(username string, password string,info string){
	db_conn.Create(&User {Username: username, Password: password, Info: info})
}

//检查用户是否存在
func CheckUserExist(username string) bool {
	var user User
	isNotFound := db_conn.First(&user, "username = ?", username).RecordNotFound()
	if isNotFound {
		return false
	}
	return true
}

//检查用户输入密码是否正确
func CheckUserPassword(username string, password string)(bool, string){
	var user User
	isNotFound := db_conn.First(&user, "username = ?", username).RecordNotFound()
	if isNotFound {
		return false, ""
	}
	return true, user.Info

}


//修改用户信息
func ModifyUser(username string, password string, info string)bool{
	var user User
	isNotFound := db_conn.First(&user, "username = ?", username).RecordNotFound()
	if isNotFound {
		return false
	}
	db_conn.Model(&user).UpdateColumns(User{Password:password,Info:info})
	return true
}

//获取用户信息
func GetUserInfo(username string) (bool,string) {
	var user User
	isNotFound := db_conn.First(&user, "username = ?", username).RecordNotFound()
	if isNotFound{
		return false, ""
	}
	return true, user.Info
}












