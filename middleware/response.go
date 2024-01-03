package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"lancer/common/snowflake"
	"lancer/common/time_tools"
	"lancer/constant"
	"lancer/global"
	"lancer/model"
	"lancer/worker"
	"net/http"
)

// Response response middleware
func Response() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		responseBody, ok := c.Get(constant.ResponseBody)
		if ok {
			c.JSON(http.StatusOK, responseBody)
			//log.Println(responseBody)
			marshal, _ := json.Marshal(responseBody)
			global.Logger.Info(string(marshal))

			//get not record
			if c.Request.Method == http.MethodGet {
				return
			}
			//record
			ActionLog(c, marshal)
			return
		} else {
			global.Logger.Error("response data has errors")
			return
		}
	}
}

// ResponseNoLog no logging response middleware
func ResponseNoLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		responseBody, ok := c.Get(constant.ResponseBody)
		if ok {
			c.JSON(http.StatusOK, responseBody)
			marshal, _ := json.Marshal(responseBody)
			global.Logger.Info(string(marshal))

			return
		} else {
			global.Logger.Error("response data has errors")
			return
		}
	}
}

func ActionLog(c *gin.Context, marshal []byte) {
	//snowId, _ := global.Snowflake.NextID()
	snowId := snowflake.Id()

	//获取用户
	var userId int64
	claims, exists := c.Get(constant.UserClaims)
	if exists {
		mapClaims := claims.(MapClaims)
		userId = int64(mapClaims["identity"].(float64))
	}

	//get body
	requestBody := c.GetString(constant.RequestBody)

	//assembly data
	log := &model.LancerActionLog{
		ID:            int64(snowId),
		UserID:        userId,
		RequestType:   c.Request.Method,
		URL:           c.Request.RequestURI,
		ClientIP:      c.ClientIP(),
		Params:        requestBody,
		RequestResult: string(marshal),
		CreatedAt:     time_tools.SerializerTimeNow(),
	}
	worker.GLogSink.Send(log)
}
