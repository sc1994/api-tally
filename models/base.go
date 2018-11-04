package models

import "github.com/jinzhu/gorm"

// GetDB 获取DB
func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:sun940622@tcp(118.24.27.231:3306)/tally?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	return db
}

// InitDB 初始化数据表
func InitDB() {
	db := GetDB()
	if db.HasTable(&Schedule{}) {
		db.AutoMigrate(&Schedule{})
	}
	if db.HasTable(&Rule{}) {
		db.AutoMigrate(&Rule{})
	}
}
