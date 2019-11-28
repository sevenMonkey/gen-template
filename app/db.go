package app

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	mysqlArgs := Config.DB.User + ":" + Config.DB.Password + "@tcp(" + Config.DB.Host + ":" + Config.DB.Port + ")/" + Config.DB.Name + "?charset=utf8&parseTime=True&loc=Local"
	//Db, err = gorm.Open("mysql", "root:123456789@/chat?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("mysql", mysqlArgs)

	DB.Callback().Create().Replace("gorm:update_time_stamp", createCallback)
	DB.Callback().Update().Replace("gorm:update_time_stamp", updateCallback)

	// 全局禁用表名复数
	DB.SingularTable(true)
	DB.DB().SetConnMaxLifetime(60 * time.Second)

	if err != nil {
		panic(err)
	}

	//defer DB.Close()
	fmt.Println("DB connect success!!!", mysqlArgs)
	DB.LogMode(true)
}

func createCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("created_at"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("updated_at"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("updated_at", time.Now().Unix())
	}
}

func CloseDB()  {
	if err :=  DB.Close(); err != nil{
		log.Fatalf("models.close err:%v", err)
	}
}