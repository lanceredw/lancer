package model

import (
	"gorm.io/gorm"
	"lancer/common/time_tools"
)

const TableNameLancerRbacMenu = "lancer_rbac_menu"

// LancerRbacMenu mapped from table <lancer_rbac_menu>
type LancerRbacMenu struct {
	ID        int64                     `gorm:"column:id;primaryKey" json:"id,string"`
	ParentID  int64                     `gorm:"column:parent_id;comment:menu parent" json:"parent_id"`
	Title     string                    `gorm:"column:title;not null;comment:menu title" json:"title"`
	Path      string                    `gorm:"column:path;not null;comment:menu path" json:"path"`
	MenuType  int32                     `gorm:"column:menu_type;default:1;comment:1 menu 2 button" json:"menu_type"`
	Icon      string                    `gorm:"column:icon;comment:menu icon" json:"icon"`
	Sort      int32                     `gorm:"column:sort;default:99;comment:sort desc" json:"sort"`
	Status    int32                     `gorm:"column:status;default:1;comment:0 disable 1 enable" json:"status"`
	CreatedAt time_tools.SerializerTime `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time_tools.SerializerTime `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt            `gorm:"column:deleted_at" json:"-"`
}

// TableName LancerRbacMenu's table name
func (*LancerRbacMenu) TableName() string {
	return TableNameLancerRbacMenu
}
