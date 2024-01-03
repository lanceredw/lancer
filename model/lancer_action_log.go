package model

import (
	"lancer/common/time_tools"
)

const TableNameLancerActionLog = "lancer_action_log"

// LancerActionLog mapped from table <lancer_action_log>
type LancerActionLog struct {
	ID            int64  `gorm:"column:id;not null;comment:id" json:"id,string"`
	UserID        int64  `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`
	RequestType   string `gorm:"column:request_type;comment:请求类型" json:"request_type"`
	URL           string `gorm:"column:url;comment:请求url" json:"url"`
	ClientIP      string `gorm:"column:client_ip;comment:ip" json:"client_ip"`
	Params        string `gorm:"column:params;comment:请求参数" json:"params"`
	RequestResult string `gorm:"column:request_result;comment:响应" json:"request_result"`
	//TODO times record start to end time
	CreatedAt time_tools.SerializerTime `gorm:"column:created_at;comment:创建时间" json:"created_at"`
}

// TableName LineActionLog's table name
func (*LancerActionLog) TableName() string {
	return TableNameLancerActionLog
}
