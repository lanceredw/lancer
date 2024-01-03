package model

import (
	"time"
)

const TableNameLancerMigrateLog = "lancer_migrate_log"

// LancerMigrateLog mapped from table <lancer_migrate_log>
type LancerMigrateLog struct {
	ID        int64     `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name;not null;comment:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;not null;comment:exec migrate time" json:"created_at"`
}

// TableName LancerMigrateLog's table name
func (*LancerMigrateLog) TableName() string {
	return TableNameLancerMigrateLog
}
