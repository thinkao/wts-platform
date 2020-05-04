package serializer

import (
	"server/account/model"
	"time"
)

type UserSerialize struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	UserType string `json:"userType"`
	Match    string	`json:"match"`
	UserInfo model.UserInfo `gorm:"ForeignKey:UserID"`

	Email string `json:"email"`
	Phone string `json:"phone"`

}

type UserInfoSerialize struct {
	UserId      int    `json:"user_id"`
	Declaration string `json:"declaration"`
	Avatar     string `json:"avatar"`

	Level    string `json:"level"`
	Integral int    `json:"integral"`
}

type DynamicSerialize struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
	ImgPath string `json:"img_path"`

	CreatedAt time.Time `json:"create_time"`
	UpdatedAt time.Time `json:"update_time"`
}

type CommentsSerialize struct {
	ID         int    `json:"id"`
	UserId     int    `json:"user_id"`
	DynamicID  int    `json:"dynamic_id"`
	CommentsID int    `json:"comments_id"`
	Content    string `json:"content"`
	ImgPath    string `json:"img_path"`

	CreatedAt time.Time `json:"create_time"`
}
