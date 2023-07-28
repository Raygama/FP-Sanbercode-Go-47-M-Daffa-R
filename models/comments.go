package models

import "time"

type Comment struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	ReviewID  int       `json:"review_id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content" gorm:"type:text"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Review    Review    `json:"-"`
	User      User      `json:"-"`
}
