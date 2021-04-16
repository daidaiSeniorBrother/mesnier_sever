package model

type User struct {
	Base
	UserId       int    `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	UserName     string `gorm:"column:user_name" json:"user_name"`
	UserAccount  string `gorm:"column:user_account" json:"user_account"`
	UserPassword string `gorm:"column:user_password" json:"user_password"`
	Phone        string `gorm:"column:phone" json:"phone"`
	OpenId       string `gorm:"column:open_id" json:"open_id"`
	Nick         string `gorm:"column:nick" json:"nick"`
}

//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为users（结构体+s）
func (u *User) TableName() string {
	return "sys_user"
}
