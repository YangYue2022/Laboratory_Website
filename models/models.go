package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strings"

	"UISwebsite_backend/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key;auto_increment" json:"id"`
}

func Setup() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = setting.DatabaseSetting.Type
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func applyQueryConditions(db *gorm.DB, queryConditions map[string]interface{}) *gorm.DB {
	for key, value := range queryConditions {
		// 检测是否包含 'LIKE' 条件
		if strings.Contains(key, "LIKE") {
			// 适当的 LIKE 语法，去除额外的 `?`
			likeValue := "%" + value.(string) + "%"
			db = db.Where(key, likeValue)
		} else {
			// 对于普通的条件，直接使用
			db = db.Where(key+" = ?", value)
		}
	}
	return db
}
