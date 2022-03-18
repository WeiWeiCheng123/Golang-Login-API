package model

type User struct {
	Username string `xorm:"pk" json:"username"`
	Passwd string `json:"passwd"`
}

func (u *User) TableName() string {
	return "user_table"
}
