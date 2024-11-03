package models

type Activity struct {
	Model

	Name string `json:"name"`
	Year string `json:"year"`
}

func AddActivity(activity *Activity) {
	db.Create(activity)
}

func GetActivityIDByName(name string) []int {
	var activities []Activity
	db.Where("name LIKE ?", "%"+name+"%").Find(&activities)

	var ids []int
	for _, activity := range activities {
		ids = append(ids, activity.ID)
	}
	return ids
}
