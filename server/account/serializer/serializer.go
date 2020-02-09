package serializer

import "time"

type UserSerialize struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`

	Email string `json:"email"`
	Phone string `json:"phone"`

	CreatedAt time.Time `json:"create_time"`
	UpdatedAt time.Time `json:"update_time"`
}

type UserInfoSerialize struct {
	UserId      int    `json:"user_id"`
	Declaration string `json:"declaration"`
	HeadUrl     string `json:"head_url"`

	Level    string `json:"level"`
	Integral int    `json:"integral"`
}

type CommentsSerialize struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
	ImgPath string `json:"img_path"`

	CreatedAt time.Time `json:"create_time"`
	UpdatedAt time.Time `json:"update_time"`
}

type CommentsReplySerialize struct {
	ID              int    `json:"id"`
	UserId          int    `json:"user_id"`
	CommentsID      int    `json:"comments_id"`
	CommentsReplyId int    `json:"comments_reply_id"`
	Content         string `json:"content"`
	ImgPath         string `json:"img_path"`

	CreatedAt time.Time `json:"create_time"`
}
