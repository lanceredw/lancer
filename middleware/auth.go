package middleware

import (
	"github.com/gin-gonic/gin"
	"lancer/common/method"
	"lancer/constant"
	"lancer/global"
	"lancer/model"
	"lancer/response"
	"net/http"
	"sync"
)

type MapClaims map[string]interface{}

var RoleMenuMap *sync.Map

func init() {
	//store data role => []menu
	RoleMenuMap = new(sync.Map)
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		//get header auth
		authHeader := c.Request.Header.Get(constant.HeaderKey)
		if authHeader == "" {
			c.JSON(http.StatusAccepted, response.ResultData{Code: http.StatusAccepted, Msg: global.TranslateMessage["AuthHeaderEmptyError"], Data: nil})
			c.Abort()
			return
		}

		//parse jwt
		userClaims, err := method.ParseJwt(authHeader)
		if err != nil {
			c.JSON(http.StatusAccepted, response.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
			c.Abort()
			return
		}

		c.Set(constant.UserClaims, userClaims)

		expirationTime, err := userClaims.GetExpirationTime()
		if err != nil {
			c.JSON(http.StatusAccepted, response.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
			c.Abort()
			return

		}

		//TODO judge expire
		expirationTime = expirationTime

		//get role
		value, ok := RoleMenuMap.Load(userClaims.RoleId)
		if !ok {
			c.JSON(http.StatusAccepted, response.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
			c.Abort()
			return

		}

		//check hasPermission
		getMenus := value.([]model.LancerRbacMenu)
		permission := false
		for _, rbacMenu := range getMenus {
			if c.Request.URL.Path == rbacMenu.Path {
				permission = true
			}
		}

		if !permission {
			c.JSON(http.StatusUnauthorized, response.ResultData{Code: http.StatusUnauthorized, Msg: err.Error(), Data: nil})
			c.Abort()
			return
		}

		c.Next()

	}
}
