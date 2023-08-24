package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Account   string `gorm:"column:account;size:100" json:"account"`
	Password  string `gorm:"column:password" json:"password"`
	Name      string `gorm:"column:name;size:100" json:"name"`
	AvatarUrl string `gorm:"column:avatarUrl" json:"avatarUrl"`
	Sex       string `gorm:"column:sex;size:100" json:"sex"`
	Phone     string `gorm:"column:phone;size:100" json:"phone"`
	Salt      string `gorm:"column:salt" json:"salt"`
	RoleId    uint   `gorm:"column:roleId" json:"roleId"`
	OpenId    string `gorm:"column:openId" json:"openId"`
}
