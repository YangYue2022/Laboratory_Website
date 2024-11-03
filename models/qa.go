package models

import "fmt"

type Qa struct {
	Model

	Que string `json:"que"`
	Ans string `json:"ans"`
}

func (qa *Qa) String() string {
	return fmt.Sprintf("Qa{que:%s ans:%s}", qa.Que, qa.Ans)
}

func GetQas(pageNum int, pageSize int, queryConditions map[string]interface{}) (qa []Qa) {
	if pageNum >= 0 {
		db = applyQueryConditions(db.New(), queryConditions)
		db.Offset(pageNum).Limit(pageSize).Find(&qa)
	} else {
		db.Find(&qa)
	}

	return
}

func GetQasTotal(queryConditions map[string]interface{}) (count int) {
	db := applyQueryConditions(db.New(), queryConditions) // 使用 db.New() 确保是一个新的连接
	db.Model(&Qa{}).Count(&count)
	return
}

func AddQa(qa *Qa) Qa {
	db.Create(qa)
	return *qa
}

func ExistQaByID(id int) bool {
	var qa Qa
	db.Select("id").Where("id = ?", id).First(&qa)
	if qa.ID > 0 {
		return true
	}

	return false
}

func EditQa(id int, qa *Qa) Qa {
	db.Model(&Qa{}).Where("id =?", id).Updates(qa)

	return *qa
}

func DeleteQa(id int) bool {
	db.Where("id =?", id).Delete(&Qa{})

	return true
}
