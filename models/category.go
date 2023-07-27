package models

type Category struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Nama      string `json:"nama" gorm:"type:varchar(255)"`
	Deskripsi string `json:"deskripsi" gorm:"type:text"`
	Games     []Game `json:"-"`
}
