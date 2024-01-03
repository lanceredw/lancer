package data

type MysqlData struct {
	Dsn         string `json:"dsn"`
	MaxIdleConn int    `json:"maxIdleConn"`
	MaxOpenConn int    `json:"maxOpenConn"`
}
