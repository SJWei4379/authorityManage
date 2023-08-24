package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name       string `gorm:"column:name" json:"name"`
	Code       string `gorm:"column:code" json:"code"` //标识
	Apis       []uint `gorm:"type:json" json:"apis"`
	RoleRouter string `gorm:"column:roleRouter" json:"roleRouter"` //角色所拥有的路由
	FirstPage  string `gorm:"column:firstPage" json:"firstPage"`   //角色首页
	Desc       string `gorm:"column:desc" json:"desc"`
}
