package model

import (
	"lancer/common/time_tools"

	"gorm.io/gorm"
)

const TableNameLancerRbacRole = "lancer_rbac_role"

// LancerRbacRole mapped from table <lancer_rbac_role>
type LancerRbacRole struct {
	ID        int64                     `gorm:"column:id;primaryKey" json:"id,string"`
	RoleName  string                    `gorm:"column:role_name;not null;comment:role name" json:"role_name"`
	Status    int32                     `gorm:"column:status;default:1;comment:0 disable 1 enable" json:"status"`
	CreatedAt time_tools.SerializerTime `gorm:"column:created_at;not null;comment:created time" json:"created_at"`
	UpdatedAt time_tools.SerializerTime `gorm:"column:updated_at;comment:update time" json:"updated_at"`
	DeletedAt gorm.DeletedAt            `gorm:"column:deleted_at;comment:deleted time" json:"-"`
}

// TableName LancerRbacRole's table name
func (*LancerRbacRole) TableName() string {
	return TableNameLancerRbacRole
}
