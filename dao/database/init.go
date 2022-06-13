package database

import (
	"BAT-douyin/model"
	"BAT-douyin/model/tcomment"
	"BAT-douyin/model/tuser"
	"BAT-douyin/model/tvideo"
	"BAT-douyin/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(conf *setting.MySQLConfig) (err error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DB,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return err
	}

	//如果不存在就创建表
	err = DB.AutoMigrate(&tuser.User{}, &tvideo.Video{}, &model.FollowUser{}, &model.FavoriteVideo{}, &tcomment.Comment{})
	if err != nil {
		return err
	}
	return nil

}
