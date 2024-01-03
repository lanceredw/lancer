package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"lancer/global"
	plugin "lancer/plugin/gorm"
	"time"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if !global.ModeData.Develop {
		logMode = logger.Error
	}
	newLogger := logger.New(
		plugin.Writer{},
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logMode,     // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(global.MysqlData.Dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.MysqlData.MaxIdleConn)
	sqlDB.SetMaxOpenConns(global.MysqlData.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
