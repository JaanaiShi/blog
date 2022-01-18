package model

import (
	"Blog/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error    // 便于接受错误


func InitDb(){
	db, err = gorm.Open(utils.DB, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		utils.DBUser,
		utils.DBPassWord,
		utils.DBHost,
		utils.DBPost,
		utils.DBName,
		))
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数，",err)
	}

	// 禁止默认表名的复数形式
	db.SingularTable(true)

	// 自动迁移模型
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	// 不能大于框架启动的时间
	db.DB().SetConnMaxLifetime(10*time.Second)

}
