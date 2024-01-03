package migratefile

import (
	"lancer/global"
	"lancer/model"
)

type User20240102171241 struct {
}

func (migrate *User20240102171241) Before() {
}

func (migrate *User20240102171241) After() {
}

func (migrate *User20240102171241) Run() {
	err := global.DB.Migrator().DropColumn(&model.LancerUser{}, "sex_e") //delete
	if err != nil {
		panic(err)
	}
}
