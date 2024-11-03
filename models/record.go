package models

type Record struct {
	Model

	ActivityID int      `json:"activity_id" gorm:"index"`
	Activity   Activity `json:"activity"`

	Name       string `json:"name"`
	PhotoUrl   string `json:"photo_url"`
	Visibility bool   `json:"visibility"`
}

func GetRecords(pageNum int, pageSize int, maps interface{}) (records []Record) {
	if mapsMap, ok := maps.(map[string]interface{}); ok {
		query := db.Preload("Activity")
		if activityIDs, exists := mapsMap["activity_ids"]; exists {
			query = query.Where("activity_id IN (?)", activityIDs)
		}

		if visibility, exists := mapsMap["visibility"]; exists {
			query = query.Where("visibility = ?", visibility)
		}

		if pageNum >= 0 {
			query.Offset(pageNum).Limit(pageSize).Find(&records)
		} else {
			query.Find(&records)
		}
	}

	return
}

func GetRecordsTotal(maps interface{}) (count int) {
	if mapsMap, ok := maps.(map[string]interface{}); ok {
		query := db.Model(&Record{})
		if activityIDs, exists := mapsMap["activity_ids"]; exists {
			query = query.Where("activity_id IN (?)", activityIDs)
		}

		if visibility, exists := mapsMap["visibility"]; exists {
			query = query.Where("visibility = ?", visibility)
		}

		query.Count(&count)
	}

	return
}

func AddRecord(record *Record) {
	db.Create(record)
}

func UpdateRecordVisibility(id int, visibility bool) {
	db.Model(&Record{}).Where("id = ?", id).Update("visibility", visibility)
}

func UpdateRecordName(id int, name string) {
	db.Model(&Record{}).Where("id =?", id).Update("name", name)
}

func DeleteRecord(id int) {
	db.Where("id = ?", id).Delete(&Record{})
}

func ExistRecordByID(id int) bool {
	var record Record
	db.Select("id").Where("id =?", id).First(&record)
	if record.ID > 0 {
		return true
	}

	return false
}
