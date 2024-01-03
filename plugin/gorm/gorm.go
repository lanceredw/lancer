package plugin

import (
	"fmt"
	"lancer/global"
)

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {

	//took over gorm 的 SQL printout
	global.Logger.Info(fmt.Sprintf(format, args...))

}
