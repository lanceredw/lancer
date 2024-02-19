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
		//&model.LancerMigrateLog{},
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

/**
need to create lancer_migrate_log

create table lancer_migrate_log
(
    id         bigint       not null
        primary key,
    name       varchar(255) not null comment 'name',
    created_at datetime     not null comment 'exec migrate time'
)
    comment 'migrate log';
*/
