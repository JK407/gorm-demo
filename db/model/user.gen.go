// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	UserID   int64  `gorm:"column:user_id;type:int;primaryKey;autoIncrement:true;comment:用户ID" json:"user_id"` // 用户ID
	UserName string `gorm:"column:user_name;type:varchar(50);not null;comment:用户名" json:"user_name"`           // 用户名
	Password string `gorm:"column:password;type:varchar(100);not null;comment:密码" json:"password"`             // 密码
	Email    string `gorm:"column:email;type:varchar(100);not null;comment:邮箱" json:"email"`                   // 邮箱
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}