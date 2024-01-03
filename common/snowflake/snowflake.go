package snowflake

import "lancer/global"

func Id() int64 {
	id, _ := global.Snowflake.NextID()
	return int64(id)
}
