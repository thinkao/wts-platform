package model

import (
	"server/problem/model"
	"time"
)

/*用户表*/
type User struct {
	ID        int       `gorm:"AUTO_INCREMENT"`
	Username  string    `gorm:"type:varchar(20);unique;not null;unique_index:username"`
	Password  string    `gorm:":varchar(32)"`
	UserType  string    `gorm:"default:'normal'"`
	Email     string    `gorm:"type:varchar(30)"`
	Phone     string    `gorm:"type:varchar(11);index:phone"`
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`

	UserInfo    UserInfo        `gorm:"ForeignKey:UserID"`
	Dynamics    []Dynamic       `gorm:"ForeignKey:UserID"`
	Comments    []Comment       `gorm:"ForeignKey:UserID"`
	Problem     []model.Problem `gorm:"ForeignKey:ID"`
	FightBlues  []model.Fight   `gorm:"ForeignKey:UserBlue"`
	FightReds   []model.Fight   `gorm:"ForeignKey:UserRed"`
	FollowUsers []*User         `gorm:"many2many:user_follow_ships;association_jointable_foreignkey:follow_user_id"`
}

func (User) TableName() string {
	return "user"
}

func (UserInfo) TableName() string {
	return "user_info"
}

/*用户信息表*/
type UserInfo struct {
	UserID      int
	Declaration string    `gorm:"type:varchar(50)"`
	Level       string    `gorm:"type:varchar(10);default:'levelOne'"`
	Integral    int       `gorm:"type:int"`
	Avatar      string    `gorm:"type:varchar(200)"`
	CreatedAt   time.Time `gorm:"column:create_time"`
	UpdatedAt   time.Time `gorm:"column:update_time"`
}

/*动态表*/
type Dynamic struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	UserID    int
	Content   string    `gorm:"type:varchar(500)"`
	ImgPath   string    `gorm:"type:varchar(200)"`
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`

	Comment []Comment `gorm:"ForeignKey:CommentsId"`
}

/*评论表*/
type Comment struct {
	ID         int `gorm:"AUTO_INCREMENT"`
	UserId     int
	DynamicId  int
	CommentsId int
	ImgPath    string    `gorm:"type:varchar(200)"`
	Content    string    `gorm:"type:varchar(500)"`
	CreatedAt  time.Time `gorm:"column:create_time"`

	SubComments []*Comment `gorm:"many2many:comment_comment;association_jointable_foreignkey:sub_comment_id"`
}

/*动态信息获取*/
type Result struct {
	Avatar		string
	Username	string
	Content 	string
	ImgPath		string

	CreateAt	time.Time
}
