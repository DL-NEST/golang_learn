package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

type BaseModel struct {
	ID        uint           `gorm:"primaryKey"`     // 主键ID
	CreatedAt time.Time      `gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

// DeviceMsg User 用户表
type DeviceMsg struct {
	BaseModel
	DeviceSubject string
	DeviceMsg     string
}

func LinkMySql() *gorm.DB {
	// 打开数据库
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		"root",
		"example",
		"127.0.0.1:3306",
		"test",
		"utf8")), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		log.Fatalln("连接数据库失败")
	}
	return db
}

func main() {
	DB = LinkMySql()

	DB.AutoMigrate(&DeviceMsg{})

	DB.Create(&DeviceMsg{
		BaseModel:     BaseModel{},
		DeviceSubject: "deviceName",
		DeviceMsg:     "bool",
	})
}
