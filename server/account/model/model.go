package model

import (
	"server/problem/model"
	"time"
)

/*用户表*/
type User struct {
	ID       int    `gorm:"AUTO_INCREMENT"`
	Username string `gorm:"type:varchar(20);unique;not null;unique_index:username"`
	Password string `gorm:":varchar(32)"`

	UserType string `gorm:"default:'normal'"`

	Email string `gorm:"type:varchar(30)"`
	Phone string `gorm:"type:varchar(11);index:phone"`

	UserInfo      UserInfo      `gorm:"ForeignKey:UserID"`
	Comments      Comments      `gorm:"ForeignKey:UserID"`
	CommentsReply CommentsReply `gorm:"ForeignKey:UserID"`
	FightBlue     model.Fight   `gorm:"ForeignKey:UserBlue"`
	FightRed      model.Fight   `gorm:"ForeignKey:UserRed"`

	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
}

/*用户信息表*/
type UserInfo struct {
	UserID      int
	Declaration string          `gorm:"type:varchar(50)"`
	Level       string          `gorm:"type:varchar(10);default:'levelOne'"`
	Integral    int             `gorm:"type:int"`
	HeadUrl     string          `gorm:"type:varchar(200)"`
	ProblemId   []model.Problem `gorm:"ForeignKey:ID"`
	UserFocusId []User          `gorm:"ForeignKey:ID"`
	CreatedAt   time.Time       `gorm:"column:create_time"`
	UpdatedAt   time.Time       `gorm:"column:update_time"`
}

/*评论表*/
type Comments struct {
	ID      int `gorm:"AUTO_INCREMENT"`
	UserID  int
	Content string `gorm:"type:varchar(500)"`
	ImgPath string `gorm:"type:varchar(200)"`

	CommentsReply CommentsReply `gorm:"ForeignKey:CommentsId"`

	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
}

/*评论回复表*/
type CommentsReply struct {
	ID              int `gorm:"AUTO_INCREMENT"`
	CommentsReplyId int
	UserId          int
	CommentsId      int
	ImgPath         string `gorm:"type:varchar(200)"`
	Content         string `gorm:"type:varchar(500)"`

	CreatedAt time.Time `gorm:"column:create_time"`
}
