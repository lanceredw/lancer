package method

import (
	"lancer/global"
	"lancer/model"
)

var (
	MigrateNameMap = map[string]struct{}{}
)

func Init() {
	InitMap()
}

func InitMap() {
	MigrateNameMap = map[string]struct{}{}

	var migrateLogs []model.LancerMigrateLog
	err := global.DB.Model(&model.LancerMigrateLog{}).Find(&migrateLogs).Error
	if err != nil {
		panic(err)
	}

	for _, migrateLog := range migrateLogs {
		MigrateNameMap[migrateLog.Name] = struct{}{}
	}
}
