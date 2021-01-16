package model

import (
	"blog/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db,err = gorm.Open(utils.Db,fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		))
	if err != nil {
		fmt.Printf("init db failed,err:%#v",err)
	}

	// 禁用默认表名的复数形式 必须放在自动迁移的前面！！！
	db.SingularTable(true)

	// 自动迁移模式
	db.AutoMigrate(&User{},&Article{},&Category{})

	// SetMaxIdleConns 设置连接池中的最大闲置连接数
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置数据库的最大连接数量
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置连接的最大可复用时间
	db.DB().SetConnMaxLifetime(10*time.Second)
}
