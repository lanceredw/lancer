package method

import (
	"fmt"
	"lancer/common/snowflake"
	"lancer/global"
	"lancer/model"
	"reflect"
	"time"
)

func MigrateRun(obj interface{}) {

	//judge exist
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	structName := objType.Name() //Get the name of the structure
	fmt.Println("struct name:", structName)

	if _, ok := MigrateNameMap[structName]; ok {
		fmt.Println(structName, "already exec!")
		return
	}

	callMethod(obj, "Before")
	callMethod(obj, "Run")
	callMethod(obj, "After")

	//when exec, insert log

	err := global.DB.Model(&model.LancerMigrateLog{}).Create(&model.LancerMigrateLog{
		ID:        snowflake.Id(),
		Name:      structName,
		CreatedAt: time.Now(),
	}).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(structName, "exec successfully")
}

// callMethod  call function
func callMethod(obj interface{}, methodName string) {
	// Obtain reflection value
	reflectValue := reflect.ValueOf(obj)

	if reflectValue.Kind() == reflect.Ptr && !reflectValue.IsNil() {
		// get function
		method := reflectValue.MethodByName(methodName)

		// judge function exist
		if method.IsValid() {
			// calling method
			method.Call(nil)
		} else {
			fmt.Println("function not exist:", methodName)
		}
	} else {
		fmt.Println("obj nil")
	}
}
