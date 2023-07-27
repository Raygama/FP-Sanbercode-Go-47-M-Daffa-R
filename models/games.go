package models

type Game struct {
	ID            int      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	CategoryID    int      `json:"category_id"`
	Nama          string   `json:"nama" gorm:"type:varchar(255)"`
	Deskripsi     string   `json:"deskripsi" gorm:"type:text"`
	Developer     string   `json:"developer" gorm:"type:varchar(255)"`
	YearPublished string   `json:"year_published" gorm:"type:varchar(10)"`
	Category      Category `json:"-"`
	Reviews       []Review `json:"-"`
}
