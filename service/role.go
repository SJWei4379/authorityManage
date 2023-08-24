package service

import (
	"authoritymanage/global"
	"authoritymanage/model"
	"authoritymanage/utils"
	"strconv"
)

func CreateRole(role model.Role) utils.Response {
	var DBRole model.Role

	if err := global.DB.Model(&role).Where("name = ?", role.Name).Find(&DBRole).Error; err != nil || DBRole.Name == "" {
		if err = global.DB.Model(&role).Create(&role).Error; err != nil {
			return utils.ErrorMess("角色存入失败", err)
		}
	} else {
		return utils.ErrorMess("角色重复", err)
	}
	return utils.SuccessMess("角色存入成功", nil)
}

func UpdateRole(role model.Role) utils.Response {
	//判断更新的API是否存在
	var DBApi model.Api
	for _, api := range role.Apis {
		if err := global.DB.Model(&model.Api{}).Where("id = ?", api).Find(&DBApi).Error; err != nil {
			return utils.ErrorMess(err.Error()+"此api不存在", api)
		}
	}
	//返回更新后的数据
	if err := global.DB.Model(&model.Role{}).Where("id = ?", role.ID).Updates(&role); err != nil {
		return utils.ErrorMess("更新失败", err)
	}
	return utils.SuccessMess("更新成功", role)
}

func GetRole(name, pageSize, currPage string) utils.Response {
	size, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return utils.ErrorMess("行数字段过长", err)
	}
	curr, err := strconv.ParseInt(currPage, 10, 64)
	if err != nil {
		return utils.ErrorMess("指定页面字段过长", err)
	}
	skip := (curr - 1) * size

	var DBRole model.Role
	if err := global.DB.Model(&model.Role{}).Where("name LIKE ?", "%"+name+"%").Limit(int(size)).Offset(int(skip)).Find(&DBRole).Error; err != nil {
		return utils.ErrorMess("角色分页查询失败", err)
	}
	var count int64
	global.DB.Table("roles").Count(&count)

	type RoleData struct {
		Role  interface{}
		Count int64
	}
	data := RoleData{
		Role:  DBRole,
		Count: count,
	}
	return utils.SuccessMess("角色分页查询成功", data)
}
