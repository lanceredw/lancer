package controller

import (
	"github.com/gin-gonic/gin"
	"lancer/constant"
	"lancer/plugin/translate"
	request "lancer/request/rbac_role_menu"
	response2 "lancer/response"
	response "lancer/response/rbac_role_menu"
	"lancer/service"
	"net/http"
)

type RbacRoleMenuController struct {
}

func NewRbacRoleMenuController() *RbacRoleMenuController {
	return &RbacRoleMenuController{}
}

// Create roleMenu create
// @Description create role_menu
// @Author Edward
// @Summary RoleMenuManage
// @Description
// @Tags RoleMenuManage
// @Accept json
// @Produce json
// @Param Request body request.RbacRoleMenuCreateRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacRoleMenuCreateResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/role_menu/create [post]
func (controller *RbacRoleMenuController) Create(c *gin.Context) {
	req := new(request.RbacRoleMenuCreateRequest)
	ret := new(response.RbacRoleMenuCreateResponse)
	err := c.ShouldBind(&req)

	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacRoleMenuService()
	ret, err = srv.Create(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// Delete roleMenu delete
// @Description delete role_menu
// @Author Edward
// @Summary RoleMenuManage
// @Description
// @Tags RoleMenuManage
// @Accept json
// @Produce json
// @Param Request body request.RbacRoleMenuCreateRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacRoleMenuCreateResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/role_menu/delete [delete]
func (controller *RbacRoleMenuController) Delete(c *gin.Context) {
	req := new(request.RbacRoleMenuDeleteRequest)
	ret := new(response.RbacRoleMenuDeleteResponse)
	err := c.ShouldBind(req)

	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacRoleMenuService()
	ret, err = srv.Delete(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// SyncData roleMenu sync data
// @Description sync data role_menu
// @Author Edward
// @Summary RoleMenuManage
// @Description
// @Tags RoleMenuManage
// @Accept json
// @Produce json
// @Param Request body request.RbacRoleMenuSyncDataRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacRoleMenuSyncDataResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/role_menu/sync_data [post]
func (controller *RbacRoleMenuController) SyncData(c *gin.Context) {
	req := new(request.RbacRoleMenuSyncDataRequest)
	ret := new(response.RbacRoleMenuSyncDataResponse)
	//err := c.ShouldBindJSON(req)

	//if err != nil {
	//	c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
	//	return
	//}

	srv := service.NewRbacRoleMenuService()
	ret, err := srv.SyncData(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// SyncDataList roleMenu sync data list
// @Description sync data list role_menu
// @Author Edward
// @Summary RoleMenuManage
// @Description
// @Tags RoleMenuManage
// @Accept json
// @Produce json
// @Param Request body request.RbacRoleMenuSyncDataListRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacRoleMenuSyncDataListResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/role_menu/sync_data_list [get]
func (controller *RbacRoleMenuController) SyncDataList(c *gin.Context) {
	req := new(request.RbacRoleMenuSyncDataListRequest)
	ret := new(response.RbacRoleMenuSyncDataListResponse)
	err := c.ShouldBindQuery(&req)

	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacRoleMenuService()
	ret, err = srv.SyncDataList(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}
