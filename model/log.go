package model

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	//User      interface{} `gorm:"column:user" json:"user"`
	User   interface{} `gorm:"type:json" json:"user"`
	Path   string      `gorm:"column:path" json:"path"`
	Method string      `gorm:"column:method" json:"method"`
	Status int         `gorm:"column:status" json:"status"`
	Query  string      `gorm:"column:query" json:"query"`
	//Body      interface{} `gorm:"column:body" json:"body"`
	Body      interface{} `gorm:"-" json:"body"`
	Ip        string      `gorm:"column:ip" json:"ip"`
	UserAgent string      `gorm:"column:userAgent" json:"userAgent"`
	Errors    string      `gorm:"column:errors" json:"errors"`
	Cost      string      `gorm:"column:cost" json:"cost"`
}
