package model

import "time"

/*题库表*/
type Problem struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	UserID    int
	Content   string    `gorm:"type:varchar(500)"`
	Option    string    `gorm:"type:varchar(500)"`
	Answer    string    `gorm:"type:varchar(500)"`
	Type      string    `gorm:"type:varchar(200)"`
	Difficult string    `gorm:"type:varchar(10)"`
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`

	Fights []*Fight `gorm:"many2many:problem_fight"`
}

/*答题对战表*/
type Fight struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	UserBlue  int
	UserRed   int
	ScoreBlue int       `gorm:"type:int"`
	ScoreRed  int       `gorm:"type:int"`
	Result    int       `gorm:"type:int"`
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
}
