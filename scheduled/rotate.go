package scheduled

import "lancer/global"

func StartRotate() {
	err := global.LumberjackLog.Rotate()
	if err != nil {
		global.Logger.Error("执行日志滚动错误:", err)
	}
}
