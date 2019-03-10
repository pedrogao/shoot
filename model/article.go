package model

import "time"

type Article struct {
	Id       int    `xorm:"not null pk autoincr INT(11)" json:"id"`
	Title    string `xorm:"not null VARCHAR(80)" json:"title"`
	Tags     string `xorm:"VARCHAR(100)" json:"tags"` // 最多5个标签
	Author   string `xorm:"VARCHAR(30)" json:"author"`
	AuthorId string `xorm:"not null INT(11)" json:"author_id"`
	Summary  string `xorm:"VARCHAR(2000)" json:"summary"`
	Content  string `xorm:"TEXT" json:"content"`
	Image    string `xorm:"VARCHAR(250)" json:"image"`

	CreateTime time.Time `xorm:"TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `xorm:"TIMESTAMP" json:"update_time"`
	DeleteTime time.Time `xorm:"TIMESTAMP" json:"delete_time"`
}
