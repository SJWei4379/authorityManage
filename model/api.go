package model

import "gorm.io/gorm"

type Api struct {
	gorm.Model
	Name   string `gorm:"column:name" json:"name"`     //api名称
	Url    string `gorm:"column:url" json:"url"`       //路由
	Method string `gorm:"column:method" json:"method"` //方法
	Desc   string `gorm:"column:desc" json:"desc"`     //描述
}
