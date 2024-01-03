package model

import (
	"gorm.io/gorm"
	"lancer/common/time_tools"
)

const TableNameLancerUser = "lancer_user"

// LancerUser mapped from table <lancer_user>
type LancerUser struct {
	ID        int64                     `gorm:"column:id;primaryKey;comment:user id" json:"id,string"`
	UserName  string                    `gorm:"column:user_name;not null;comment:account" json:"user_name"`
	NickName  string                    `gorm:"column:nick_name;comment:nick name" json:"nick_name"`
	Avatar    string                    `gorm:"column:avatar;comment:avatar url" json:"avatar"`
	RoleID    int64                     `gorm:"column:role_id;comment:role id" json:"role_id"`
	Password  string                    `gorm:"column:password;comment:password" json:"password"`
	Salt      string                    `gorm:"column:salt;comment:password md5 salt" json:"salt"`
	Sex       int32                     `gorm:"column:sex;comment:1 male 2 female" json:"sex"`
	Email     string                    `gorm:"column:email;comment:email" json:"email"`
	Phone     string                    `gorm:"column:phone;comment:phone" json:"phone"`
	Status    int32                     `gorm:"column:status;default:1;comment:0 disable 1 enable" json:"status"`
	SexE      int32                     `gorm:"column:sex_e;default:1;comment:0 disable 1 enable" json:"sex_e"`
	CreatedAt time_tools.SerializerTime `gorm:"column:created_at;not null;comment:created time" json:"created_at"`
	UpdatedAt time_tools.SerializerTime `gorm:"column:updated_at;comment:update time" json:"updated_at"`
	DeletedAt gorm.DeletedAt            `gorm:"column:deleted_at;comment:deleted time" json:"-"`
}

// TableName LancerUser's table name
func (*LancerUser) TableName() string {
	return TableNameLancerUser
}
