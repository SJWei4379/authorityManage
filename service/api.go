package service

import (
	"authoritymanage/global"
	"authoritymanage/model"
	"authoritymanage/utils"
	"strconv"
)

func CreateApi(api model.Api) utils.Response {
	if err := global.DB.Model(&api).Create(&api).Error; err != nil {
		return utils.ErrorMess("插入api失败", err)
	}
	return utils.SuccessMess("插入api成功", api.ID)
}

func DeleteApi(id int) utils.Response {
	var DBApi model.Api
	//在角色表中删除此权限
	if err := global.DB.Model(&model.Api{}).Where("id = ?", id).Delete(&DBApi).Error; err != nil {
		return utils.ErrorMess("删除api失败", err)
	}
	return utils.SuccessMess("删除api成功", id)

}

func UpdateApi(api model.Api) utils.Response {
	if err := global.DB.Model(&model.Api{}).Where("id = ?", api.ID).Updates(&api).Error; err != nil {
		return utils.ErrorMess("更新失败", err)
	}
	return utils.SuccessMess("更新成功", api.ID)
}

func GetApi(method, name, pageSize, currPage string) utils.Response {
	size, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return utils.ErrorMess("行数字段过长", err)
	}
	curr, err := strconv.ParseInt(currPage, 10, 64)
	if err != nil {
		return utils.ErrorMess("指定页面字段过长", err)
	}
	skip := (curr - 1) * size

	var DBApi []model.Api
	if err := global.DB.Model(&model.Api{}).Where("name LIKE ? OR method LIKE ?", "%"+name+"%", "%"+method+"%").Limit(int(size)).Offset(int(skip)).Find(&DBApi).Error; err != nil {
		return utils.ErrorMess("查询分页失败", err)
	}
	var count int64
	global.DB.Table("apis").Count(&count)

	type APIData struct {
		API   interface{}
		Count int64
	}
	data := APIData{
		API:   DBApi,
		Count: count,
	}

	return utils.SuccessMess("查询成功", data)
}
