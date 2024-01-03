package router

import (
	"github.com/gin-gonic/gin"
	"lancer/controller"
	"lancer/middleware"
)

func ApiRouter(router *gin.RouterGroup) {

	//ping
	pingController := controller.NewPingController()
	pingGroup := router.Group("/ping")
	pingGroup.GET("/ping", pingController.Ping)    //ping
	pingGroup.POST("/empty", pingController.Empty) //post  maybe empty data

	//user
	userController := controller.NewUserController()
	userGroup := router.Group("/user", middleware.ResponseNoLog())
	userGroup.POST("/create", userController.Create)   //create user
	userGroup.DELETE("/delete", userController.Delete) //delete user
	userGroup.PUT("/update", userController.Update)    //update user
	userGroup.GET("/list", userController.List)        //select users
	userGroup.POST("/login", userController.Login)     //user login

	//role
	rbacRoleController := controller.NewRbacRoleController()
	rbacRoleGroup := router.Group("/role", middleware.ResponseNoLog())
	rbacRoleGroup.POST("/create", rbacRoleController.Create)   //create role
	rbacRoleGroup.GET("/list", rbacRoleController.List)        //select roles
	rbacRoleGroup.DELETE("/delete", rbacRoleController.Delete) //delete role
	rbacRoleGroup.PUT("/update", rbacRoleController.Update)    //update role

	//menu
	rbacMenuController := controller.NewRbacMenuController()
	rbacMenuGroup := router.Group("/menu", middleware.ResponseNoLog())
	rbacMenuGroup.POST("/create", rbacMenuController.Create)   //create menu
	rbacMenuGroup.GET("/list", rbacMenuController.List)        //select menus
	rbacMenuGroup.PUT("/update", rbacMenuController.Update)    //update menu
	rbacMenuGroup.DELETE("/delete", rbacMenuController.Delete) //delete menu

	//role_menu
	rbacRoleMenuController := controller.NewRbacRoleMenuController()
	rbacRoleMenuGroup := router.Group("/role_menu", middleware.ResponseNoLog())
	rbacRoleMenuGroup.POST("/create", rbacRoleMenuController.Create)              //create role_menu
	rbacRoleMenuGroup.DELETE("/delete", rbacRoleMenuController.Delete)            //delete role_menu
	rbacRoleMenuGroup.POST("/sync_data", rbacRoleMenuController.SyncData)         //sync role menu to map
	rbacRoleMenuGroup.GET("/sync_data_list", rbacRoleMenuController.SyncDataList) //role menu list

	//monitor
	monitorController := controller.NewMonitorController()
	monitorGroup := router.Group("/monitor", middleware.ResponseNoLog())
	monitorGroup.GET("/goroutines", monitorController.Goroutines) //check goroutines
	monitorGroup.GET("/file_stats", monitorController.FileStats)  //get file status
}
