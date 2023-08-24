package service

import (
	"authoritymanage/global"
	"authoritymanage/model"
	"authoritymanage/utils"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

func Register(user model.User) utils.Response {
	var DBUser model.User
	var DBRole model.Role
	if err := global.DB.Model(&user).Where("account = ?", user.Account).Find(&DBUser).Error; err != nil || DBUser.Account == "" {
		//判断角色是否存在
		if err = global.DB.Model(&model.Role{}).Where("id = ?", user.RoleId).Find(&DBRole).Error; err != nil || DBRole.Name == "" {
			return utils.ErrorMess("角色不存在", err)
		}
		rand.Seed(time.Now().Unix()) //根据时间戳生成种子
		//生成盐
		salt := strconv.FormatInt(rand.Int63(), 10)
		encryptedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password+salt), bcrypt.DefaultCost)
		if err != nil {
			return utils.ErrorMess("密码加密失败", err)
		}
		user.Password, user.Salt = string(encryptedPass), salt
		user.CreatedAt = time.Now()
		if err := global.DB.Model(&user).Create(&user).Error; err != nil {
			return utils.ErrorMess("存入错误", err)
		}
		return utils.SuccessMess("注册成功", user)
	} else {
		return utils.ErrorMess("账号重复", err)
	}
}

func Login(user model.User) utils.Response {
	var DBUser model.User
	//校验账号
	if err := global.DB.Model(&model.User{}).Where("account = ?", user.Account).Find(&DBUser).Error; err != nil || DBUser.Account == "" {
		return utils.ErrorMess("账号错误", err)
	}
	//校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password+DBUser.Salt)); err != nil {
		return utils.ErrorMess("密码错误", err)
	}
	//查询角色信息
	var role model.Role
	if err := global.DB.Model(&model.Role{}).Where("id = ?", DBUser.RoleId).Find(&role).Error; err != nil || role.Name == "" {
		return utils.ErrorMess("此用户角色不存在", err)
	}
	//生成token
	//token, err := midd

	return utils.SuccessMess("登录成功", role)

}
