package conf

import (
	"github.com/sony/sonyflake"
)

func InitSnowflake() (snowflake *sonyflake.Sonyflake) {
	snowflake = sonyflake.NewSonyflake(sonyflake.Settings{})

	//if snowflake is nil ,need to do
	if snowflake == nil {
		snowflake = sonyflake.NewSonyflake(sonyflake.Settings{
			MachineID: func() (uint16, error) {
				return 1, nil
			},
		})
	}

	return snowflake
}
