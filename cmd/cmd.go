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
	"lancer/worker"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {

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

	// init gin
	r := gin.New()
	r.Use(ginzap.Ginzap(global.Logger.Desugar(), time.DateTime, true))
	r.Use(ginzap.RecoveryWithZap(global.Logger.Desugar(), true))
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

	//run worker
	worker.RunWork()

	localIP := network.GetLocalIP()
	fmt.Printf("application :http://%s:%s \n", localIP, port)
	fmt.Printf("swagger doc: http://%s:%s/swagger/index.html \n", network.GetLocalIP(), port)

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
