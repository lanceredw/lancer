package controller

import (
	"github.com/gin-gonic/gin"
	"lancer/constant"
	"lancer/plugin/translate"
	request "lancer/request/rbac_role"
	response2 "lancer/response"
	response "lancer/response/rbac_role"
	"lancer/service"
	"net/http"
)

type RbacRoleController struct {
}

func NewRbacRoleController() *RbacRoleController {
	return &RbacRoleController{}
}

// Create role create
// @Description create role
// @Author Edward
// @Summary RoleManage
// @Description
// @Tags RoleManage
// @Accept json
// @Produce json
// @Param Request body request.RbacRoleCreateRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacRoleCreateResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/role/create [post]
func (controller *RbacRoleController) Create(c *gin.Context) {
	req := new(request.RbacRoleCreateRequest)
	ret := new(response.RbacRoleCreateResponse)
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacRoleService()
	ret, err = srv.Create(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// List role list
// @Description select roles
// @Author Edward
// @Summary RoleManage
// @Description
// @Tags RoleManage
// @Accept json
// @Produce json
// @Param Request query request.RbacRoleListRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacRoleListResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/role/list [get]
func (controller *RbacRoleController) List(c *gin.Context) {
	req := new(request.RbacRoleListRequest)
	ret := new(response.RbacRoleListResponse)
	err := c.ShouldBindQuery(req)

	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacRoleService()
	ret, err = srv.List(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// Delete role delete
// @Description delete role
// @Author Edward
// @Summary RoleManage
// @Description
// @Tags RoleManage
// @Accept json
// @Produce json
// @Param Request query request.RbacRoleDeleteRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacRoleDeleteResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/role/delete [delete]
func (controller *RbacRoleController) Delete(c *gin.Context) {
	req := new(request.RbacRoleDeleteRequest)
	ret := new(response.RbacRoleDeleteResponse)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacRoleService()
	ret, err = srv.Delete(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// Update role info update
// @Description role info update
// @Author Edward
// @Summary RoleManage
// @Description
// @Tags RoleManage
// @Accept json
// @Produce json
// @Param Request body request.RbacRoleUpdateRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacRoleUpdateResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/role/update [put]
func (controller *RbacRoleController) Update(c *gin.Context) {
	req := new(request.RbacRoleUpdateRequest)
	ret := new(response.RbacRoleUpdateResponse)

	err := c.ShouldBindJSON(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacRoleService()
	ret, err = srv.Update(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}
