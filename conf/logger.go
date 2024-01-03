package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"lancer/global"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	// getWriteSyncer:output to file;  os.Stdout：output to console
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)
	if !global.ModeData.Develop {
		logMode = zapcore.InfoLevel
		core = zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer()), logMode)
	}

	logger := zap.New(core, zap.AddCaller()).Sugar() //zap.AddCaller() Display file name and line number

	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".log"
	fmt.Println(stLogFilePath)

	lumberJackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"),    // maximum size of log file (M), automatically split after exceeding the limit
		MaxBackups: viper.GetInt("log.MaxBackups"), // maximum number of retained old files
		MaxAge:     viper.GetInt("log.MaxAge"),     // maximum number of days to retain old files
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackSyncer)
}
