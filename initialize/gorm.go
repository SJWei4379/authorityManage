package initialize

import (
	"authoritymanage/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysqlInit() {
	if global.DB == nil {
		global.DB = getGormMysqlClient("root:root@tcp(localhost:3306)/authoritymanage?charset=utf8mb4&parseTime=True&loc=Local")
	}
}

func getGormMysqlClient(url string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println("加载服务器异常，异常信息：", err)
		return nil
	}

	//获取通用数据库对象 sql.DB，然后使用其提供的功能
	mysqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("获取通用数据库对象异常信息：", err)
	}
	//验证连接
	err = mysqlDB.Ping()
	if err != nil {
		fmt.Println("数据库连接异常信息：", err.Error())
		return nil
	}
	fmt.Println("数据库连接成功")
	return DB
}
