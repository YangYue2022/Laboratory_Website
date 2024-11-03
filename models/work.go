package models

import "fmt"

type Work struct {
	Model

	Type string `gorm:"type:enum('期刊论文', '会议论文', '专著', '软著/专利', '其他');not null" json:"type"`
	Year int    `gorm:"type:year;not null" json:"year"`
	Desc string `gorm:"type:varchar(255);not null" json:"desc"`
}

func (word *Work) String() string {
	return fmt.Sprintf("Work{type:%s year:%d desc:%s}", word.Type, word.Year, word.Desc)
}

func GetWorks(pageNum int, pageSize int, queryConditions map[string]interface{}) (works []Work) {
	if pageNum >= 0 {
		db = applyQueryConditions(db.New(), queryConditions)
		db.Offset(pageNum).Limit(pageSize).Find(&works)
	} else {
		db.Find(&works)
	}
	return
}

func GetWorksTotal(queryConditions map[string]interface{}) (count int) {
	db := applyQueryConditions(db.New(), queryConditions) // 使用 db.New() 确保是一个新的连接
	db.Model(&Work{}).Count(&count)
	return
}

func AddWork(work *Work) bool {
	db.Create(work)
	return true
}

func ExistWorkByID(id int) bool {
	var work Work
	db.Select("id").Where("id = ?", id).First(&work)
	if work.ID > 0 {
		return true
	}

	return false
}

func EditWork(id int, work *Work) bool {
	db.Model(&Work{}).Where("id = ?", id).Updates(work)

	return true
}

func DeleteWork(id int) bool {
	db.New().Where("id = ?", id).Delete(&Work{})

	return true
}
