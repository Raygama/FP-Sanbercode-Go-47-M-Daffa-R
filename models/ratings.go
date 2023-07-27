package models

type Rating struct {
	ID        int      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Score     string   `json:"score" gorm:"type:varchar(10)"`
	Deskripsi string   `json:"deskripsi" gorm:"type:text"`
	Reviews   []Review `json:"-"`
}
