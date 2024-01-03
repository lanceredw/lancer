package model

import (
	"time"
)

const TableNameLancerRbacRoleMenu = "lancer_rbac_role_menu"

// LancerRbacRoleMenu mapped from table <lancer_rbac_role_menu>
type LancerRbacRoleMenu struct {
	ID        int64     `gorm:"column:id;primaryKey;comment:id" json:"id,string"`
	RoleID    int64     `gorm:"column:role_id;not null;comment:role id" json:"role_id,string"`
	MenuID    int64     `gorm:"column:menu_id;not null;comment:menu id" json:"menu_id,string"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
}

// TableName LancerRbacRoleMenu's table name
func (*LancerRbacRoleMenu) TableName() string {
	return TableNameLancerRbacRoleMenu
}
