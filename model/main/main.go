package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"web/model"
)

func main() {

	dsn := "root:root@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"

	logger := logger2.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger2.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			Colorful:      true,        //禁用彩色打印
			LogLevel:      logger2.Info,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: logger,
	})

	if err != nil {
		panic(err)
	}

	//_ = db.Set(model.TableOptions, model.GetOptions("字典类型表")).AutoMigrate(&model.DictType{})
	//
	//_ = db.Set(model.TableOptions, model.GetOptions("字典表")).AutoMigrate(&model.Dict{})
	//
	//_ = db.Set(model.TableOptions, model.AutoIncrementOptions("用户表")).AutoMigrate(&model.User{})
	//
	//_ = db.Set(model.TableOptions, model.GetOptions("角色表")).AutoMigrate(&model.Role{})
	//
	//_ = db.Set(model.TableOptions, model.GetOptions("菜单表")).AutoMigrate(&model.Menu{})

	//_ = db.Set(model.TableOptions, model.GetOptions("角色菜单权限表")).AutoMigrate(&model.RoleMenu{})

	//_ = db.Set(model.TableOptions, model.GetOptions("用户表")).AutoMigrate(&model.User{})

	//_ = db.Set(model.TableOptions, model.GetOptions("医疗案例")).AutoMigrate(&model.MedicalCases{})

	_ = db.Set(model.TableOptions, model.GetOptions("典籍分类")).AutoMigrate(&model.ClassicsCategory{})
	_ = db.Set(model.TableOptions, model.GetOptions("典籍")).AutoMigrate(&model.Classics{})
	_ = db.Set(model.TableOptions, model.GetOptions("典籍内容")).AutoMigrate(&model.ClassicsContent{})

}
