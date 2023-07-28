package models

import "time"

type Review struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	GameID    int       `json:"game_id"`
	RatingID  int       `json:"rating_id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Content   string    `json:"content" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Game      Game      `json:"-"`
	Rating    Rating    `json:"-"`
	User      User      `json:"-"`
	Comments  []Comment `json:"-"`
}
