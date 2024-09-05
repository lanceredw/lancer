package cmd

import (
	"context"
	"errors"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lancer/common/network"
	"lancer/conf"
	"lancer/global"
	"lancer/router"
	"lancer/scheduled"
	"lancer/worker"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	var (
		err error
	)

	// init setting
	conf.InitSettings()

	//init config
	conf.IniConfig()

	// init logger
	global.Logger = conf.InitLogger()

	// init lan

	// init Snowflake
	global.Snowflake = conf.InitSnowflake()

	// init mysql database
	db, err := conf.InitDB()
	if err != nil {
		panic(err)
	}
	global.DB = db

	//init redis
	rdb, err := conf.InitRedis()
	if err != nil {
		panic(err)
	}

	global.Redis = rdb

	//TODO 1.启动的时候执行一次
	//TODO 2.每天12点后会自动执行一次新的日志文件生成
	//TODO 3.判断是否需要生成新的日志，每次执行日志年月日都存储到变量，判断到不一样的时候生成新的日志文件
	//if today != time.Now().Format("2006-01-02") {
	//	//TODO 4.生成新的日志文件
	//}
	//logFilePath := filepath.Join("./log", time.Now().Format("2006-01-02")+".log")
	//
	//// 以追加模式打开日志文件
	//logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	panic(err)
	//}
	//defer logFile.Close()

	//slog日志
	//global.SLogger = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{
	//	AddSource:   true,
	//	Level:       slog.LevelDebug,
	//	ReplaceAttr: nil,
	//}))

	// init gin
	r := gin.New()
	//gin.DefaultWriter = io.MultiWriter(logFile)
	r.Use(ginzap.Ginzap(global.Logger.Desugar(), time.DateTime, true))
	r.Use(ginzap.RecoveryWithZap(global.Logger.Desugar(), true))
	//r.Use(middleware.SlogMiddleware(global.SLogger))
	router.InitRouter(r)

	//init validation
	language := global.ModeData.Language
	if language == "" {
		language = "en"
	}
	err = conf.InitValidation(language)
	if err != nil {
		panic(err)
	}

	//start service
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: r,
	}

	//run cron
	scheduled.NewCron()

	//run worker
	worker.RunWork()

	localIP := network.GetLocalIP()
	fmt.Printf("application :http://%s:%s \n", localIP, port)
	fmt.Printf("swagger doc: http://%s:%s/swagger/index.html \n", network.GetLocalIP(), port)

	//global.SLogger.Info("Server Start...")
	global.Logger.Info("Server Start...")
	go func() {
		// lister serve
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	// graceful  stop the server（setting 5 seconds timeout）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("Stop Server Error: %s", err.Error()))
	}
	global.Logger.Info("Stop Server Success")
}

func Clean() {
	global.Logger.Info("=========Clean=========")
}
