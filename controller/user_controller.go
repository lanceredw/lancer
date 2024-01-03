package controller

import (
	"github.com/gin-gonic/gin"
	"lancer/constant"
	"lancer/plugin/translate"
	request "lancer/request/user"
	response2 "lancer/response"
	response "lancer/response/user"
	"lancer/service"
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// Create user create
// @Description create user
// @Author Edward
// @Summary UserManage
// @Description
// @Tags UserManage
// @Accept json
// @Produce json
// @Param Request body request.UserCreateRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.UserCreateResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/user/create [post]
func (controller *UserController) Create(c *gin.Context) {
	req := new(request.UserCreateRequest)
	ret := new(response.UserCreateResponse)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewUserService()
	ret, err = srv.Create(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// Delete user delete
// @Description delete user
// @Author Edward
// @Summary UserManage
// @Description
// @Tags UserManage
// @Accept json
// @Produce json
// @Param Request query request.UserDeleteRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.UserDeleteResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/user/delete [delete]
func (controller *UserController) Delete(c *gin.Context) {
	req := new(request.UserDeleteRequest)
	ret := new(response.UserDeleteResponse)

	err := c.ShouldBindQuery(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewUserService()
	ret, err = srv.Delete(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// Update user info update
// @Description user info update
// @Author Edward
// @Summary UserManage
// @Description
// @Tags UserManage
// @Accept json
// @Produce json
// @Param Request body request.UserUpdateRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.UserUpdateResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/user/update [put]
func (controller *UserController) Update(c *gin.Context) {
	req := new(request.UserUpdateRequest)
	ret := new(response.UserUpdateResponse)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewUserService()
	ret, err = srv.Update(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// List user list
// @Description select users
// @Author Edward
// @Summary UserManage
// @Description
// @Tags UserManage
// @Accept json
// @Produce json
// @Param Request query request.UserListRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.UserListResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/user/list [get]
func (controller *UserController) List(c *gin.Context) {
	req := new(request.UserListRequest)
	ret := new(response.UserListResponse)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewUserService()
	ret, err = srv.List(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// Login user login
// @Description  user login
// @Author Edward
// @Summary UserManage
// @Description
// @Tags UserManage
// @Accept json
// @Produce json
// @Param Request body request.UserLoginRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.UserLoginResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/user/login [post]
func (controller *UserController) Login(c *gin.Context) {
	req := new(request.UserLoginRequest)
	ret := new(response.UserLoginResponse)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}
	srv := service.NewUserService()
	ret, err = srv.Login(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return

}
