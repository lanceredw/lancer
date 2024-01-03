package controller

import (
	"github.com/gin-gonic/gin"
	"lancer/constant"
	"lancer/plugin/translate"
	request "lancer/request/rbac_menu"
	response2 "lancer/response"
	response "lancer/response/rbac_menu"
	"lancer/service"
	"net/http"
)

type RbacMenuController struct {
}

func NewRbacMenuController() *RbacMenuController {
	return &RbacMenuController{}
}

// Create menu create
// @Description create menu
// @Author Edward
// @Summary MenuManage
// @Description
// @Tags MenuManage
// @Accept json
// @Produce json
// @Param Request body request.RbacMenuCreateRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacMenuCreateResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/menu/create [post]
func (controller *RbacMenuController) Create(c *gin.Context) {
	req := new(request.RbacMenuCreateRequest)
	ret := new(response.RbacMenuCreateResponse)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacMenuService()
	ret, err = srv.Create(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// List menu list
// @Description select menus
// @Author Edward
// @Summary MenuManage
// @Description
// @Tags MenuManage
// @Accept json
// @Produce json
// @Param Request query request.RbacMenuListRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacMenuListResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/menu/list [get]
func (controller *RbacMenuController) List(c *gin.Context) {
	req := new(request.RbacMenuListRequest)
	ret := new(response.RbacMenuListResponse)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacMenuService()
	ret, err = srv.List(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

// Update menu update
// @Description menu update
// @Author Edward
// @Summary MenuManage
// @Description
// @Tags MenuManage
// @Accept json
// @Produce json
// @Param Request query request.RbacMenuUpdateRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacMenuUpdateResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/menu/update [put]
func (controller *RbacMenuController) Update(c *gin.Context) {
	req := new(request.RbacMenuUpdateRequest)
	ret := new(response.RbacMenuUpdateResponse)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacMenuService()
	ret, err = srv.Update(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return

}

// Delete menu delete
// @Description menu delete
// @Author Edward
// @Summary MenuManage
// @Description
// @Tags MenuManage
// @Accept json
// @Produce json
// @Param Request query request.RbacMenuDeleteRequest true "request message"
// @Param Authorization header string true "request message"
// @Success 200 {object} response.RbacMenuDeleteResponse
// @Failure 202 {object}   response2.ResultData
// @Router /api/menu/delete [delete]
func (controller *RbacMenuController) Delete(c *gin.Context) {
	req := new(request.RbacMenuDeleteRequest)
	ret := new(response.RbacMenuDeleteResponse)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewRbacMenuService()
	ret, err = srv.Delete(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}
