package dao

type Links struct {
	Uid         int64  `gorm:"column:uid;" json:"uid"`
	Url         string `gorm:"column:url;" json:"url"`
	InsertAt    int64  `gorm:"column:insert_at;" json:"insertAt"`
	AvailableAt int64  `gorm:"column:available_at;" json:"availableAt"`
	Available   int32  `gorm:"column:available;" json:"available"`
}
