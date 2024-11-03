package models

import "fmt"

type Member struct {
	Model

	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Identity    string `json:"identity" gorm:"type:enum('教职员', '学生');not null"`
	Title       string `json:"title" gorm:"type:enum('教授', '副教授','讲师','工程师','博士研究生','硕士研究生','本科生','毕业生');not null"`
	Description string `json:"description" gorm:"type:text"`
	PhotoUrl    string `json:"photo_url"`
}

func GetMembers(pageNum int, pageSize int) (members []*Member, err error) {
	if pageNum >= 0 {
		db.Offset(pageNum).Limit(pageSize).Find(&members)
	} else {
		db.Find(&members)
	}
	fmt.Println(members)
	return members, nil
}

func GetMembersTotal() (count int, err error) {
	result := db.Model(&Member{}).Count(&count)
	if result.Error != nil {
		err = result.Error
		return 0, err
	}
	return count, nil
}

func AddMember(member *Member) error {
	err := db.Create(member).Error
	if err != nil {
		return err
	}
	return nil
}

func ExistMemberByID(id int) (bool, error) {
	var member Member
	result := db.New().Where("id = ?", id).First(&member)
	if result.Error != nil {
		return false, result.Error
	}
	if member.ID > 0 {
		return true, nil
	}

	return false, nil
}

func EditMember(member *Member) error {
	err := db.Model(&Member{}).Where("id = ?", member.ID).Updates(member).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteMember(id int) error {
	err := db.Where("id = ?", id).Delete(&Member{}).Error
	if err != nil {
		return err
	}
	return nil
}
