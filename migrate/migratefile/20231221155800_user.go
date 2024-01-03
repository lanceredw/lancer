package migratefile

import (
	"fmt"
	"lancer/global"
	"lancer/model"
)

type User20231221155800 struct {
}

func (migrate *User20231221155800) Before() {
	fmt.Println("before")
}

func (migrate *User20231221155800) After() {
	fmt.Println("after")
}

func (migrate *User20231221155800) Run() {
	fmt.Println("run")

	err := global.DB.Migrator().AddColumn(&model.LancerUser{}, "sex_e") //Must add database does not exist, it exists in the model
	if err != nil {
		panic(err)
	}
}
