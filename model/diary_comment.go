package model

import "time"

type DiaryComment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Comment   string    `json:"comment" gorm:"not null"`
	Diary     Diary     `json:"diary" gorm:"foreignKey:DiaryId; constraint:OnDelete:CASCADE"`
	DiaryId   uint      `json:"diary_id" gorm:"foreignKey:DiaryId; constraint:OnDelete:CASCADE"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DiaryCommentResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Comment   string    `json:"comment" gorm:"not null"`
	Diary     Diary     `json:"diary" gorm:"foreignKey:DiaryId; constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
