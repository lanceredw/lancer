package service

import (
	"errors"
	"lancer/common/snowflake"
	"lancer/global"
	"lancer/middleware"
	"lancer/model"
	request "lancer/request/rbac_role_menu"
	response "lancer/response/rbac_role_menu"
	"net/http"
	"time"
)

type RbacRoleMenuService struct {
}

func NewRbacRoleMenuService() *RbacRoleMenuService {
	return &RbacRoleMenuService{}
}

func (service *RbacRoleMenuService) Create(req *request.RbacRoleMenuCreateRequest) (ret *response.RbacRoleMenuCreateResponse, err error) {
	ret = new(response.RbacRoleMenuCreateResponse)

	var count int64
	err = global.DB.Model(&model.LancerRbacRoleMenu{}).Where("role_id = ? AND menu_id = ?", req.RoleId, req.MenuId).First(&count).Error
	if err != nil {
		return nil, err
	}

	if count >= 1 {
		return nil, errors.New(global.TranslateMessage["RbacRoleMenuRecordRepeatError"])
	}

	rbacRoleMenu := model.LancerRbacRoleMenu{
		ID:        snowflake.Id(),
		RoleID:    req.RoleId,
		MenuID:    req.MenuId,
		CreatedAt: time.Now(),
	}

	err = global.DB.Model(&model.LancerRbacRoleMenu{}).Create(&rbacRoleMenu).Error
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["CreateSuccess"]
	return
}

func (service *RbacRoleMenuService) Delete(req *request.RbacRoleMenuDeleteRequest) (ret *response.RbacRoleMenuDeleteResponse, err error) {
	ret = new(response.RbacRoleMenuDeleteResponse)

	err = global.DB.Where("id = ?", req.Id).Unscoped().Delete(&model.LancerRbacRoleMenu{}).Error
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["DeleteSuccess"]
	return
}

func (service *RbacRoleMenuService) SyncData(req *request.RbacRoleMenuSyncDataRequest) (ret *response.RbacRoleMenuSyncDataResponse, err error) {
	ret = new(response.RbacRoleMenuSyncDataResponse)

	//select all roles
	var roles []*model.LancerRbacRole
	err = global.DB.Model(&model.LancerRbacRole{}).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	if len(roles) == 0 {
		return nil, errors.New(global.TranslateMessage["RbacRoleMenuNoRoleError"])
	}

	for _, role := range roles {
		//select role`s menus
		var roleMenus []*response.RoleMenuLineData
		err = global.DB.Model(&model.LancerRbacRoleMenu{}).
			Preload("Menu").
			Where("role_id = ?", role.ID).Find(&roleMenus).Error
		if err != nil {
			return nil, err
		}

		var menus []*model.LancerRbacMenu

		for _, roleMenu := range roleMenus {
			if roleMenu.Menu != nil {
				menus = append(menus, roleMenu.Menu)
			}
		}

		if len(menus) != 0 {

			//store role=> []menus to RoleMenuMap
			middleware.RoleMenuMap.Store(role.ID, menus)
		}

	}

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["ProcessSuccess"]

	return
}

func (service *RbacRoleMenuService) SyncDataList(req *request.RbacRoleMenuSyncDataListRequest) (ret *response.RbacRoleMenuSyncDataListResponse, err error) {
	ret = new(response.RbacRoleMenuSyncDataListResponse)

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["FindSuccess"]

	roleMenuMap := make(map[int64][]*model.LancerRbacMenu)

	if req.RoleId != 0 {
		value, ok := middleware.RoleMenuMap.Load(req.RoleId)
		if !ok {
			return
		}

		getMenus := value.([]*model.LancerRbacMenu)
		roleMenuMap[req.RoleId] = getMenus

		ret.Data.List = roleMenuMap
		return
	}

	middleware.RoleMenuMap.Range(func(key, value any) bool {

		roleId := key.(int64)
		getMenus := value.([]*model.LancerRbacMenu)
		roleMenuMap[roleId] = getMenus
		return true

	})
	ret.Data.List = roleMenuMap

	return

}
