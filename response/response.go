package response

type ResultData struct {
	Code int         `json:"code" gorm:"code"`
	Msg  string      `json:"msg" gorm:"msg"`
	Data interface{} `json:"data" gorm:"data"`
}
