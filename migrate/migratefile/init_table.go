package migratefile

import (
	"fmt"
	"lancer/global"
	"lancer/model"
)

type InitTable struct {
}

func (migrate *InitTable) Before() {
}

func (migrate *InitTable) After() {
}

func (migrate *InitTable) Run() {
	err := global.DB.Migrator().AutoMigrate(&model.LancerActionLog{},
		&model.LancerMigrateLog{},
		&model.LancerUser{},
		&model.LancerRbacMenu{},
		&model.LancerRbacRole{},
		&model.LancerRbacRoleMenu{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("init table successfully")

}
