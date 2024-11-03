package models

import "fmt"

type Reward struct {
	Model

	Type    string `gorm:"type:enum('竞赛','其他');not null" json:"type"`
	Year    int    `gorm:"type:year;not null" json:"year"`
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	Project string `gorm:"type:varchar(255);not null" json:"project"`
	Winner  string `gorm:"type:varchar(255);not null" json:"winner"`
}

func (reward *Reward) String() string {
	return fmt.Sprintf("Reward{type:%s year:%d name:%s project:%s winner:%s}", reward.Type, reward.Year, reward.Name, reward.Project, reward.Winner)
}

func GetRewards(pageNum int, pageSize int, queryConditions map[string]interface{}) (rewards []Reward) {
	if pageNum >= 0 {
		db = applyQueryConditions(db.New(), queryConditions)
		db.Offset(pageNum).Limit(pageSize).Find(&rewards)
	} else {
		db.Find(&rewards)
	}

	return
}

func GetRewardsTotal(maps interface{}) (count int) {
	db.Model(&Reward{}).Where(maps).Count(&count)

	return
}

func AddReward(reward *Reward) bool {
	db.Create(reward)

	return true
}

func ExistRewardByID(id int) bool {
	var reward Reward
	db.New().Select("id").Where("id = ?", id).First(&reward)
	if reward.ID > 0 {
		return true
	}

	return false
}

func EditReward(id int, reward *Reward) bool {
	db.Model(&Reward{}).Where("id = ?", id).Updates(reward)

	return true
}

func DeleteReward(id int) bool {
	db.Where("id = ?", id).Delete(&Reward{})
	return true
}
