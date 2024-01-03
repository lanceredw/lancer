package service

import (
	"errors"
	"lancer/common/snowflake"
	"lancer/common/time_tools"
	"lancer/global"
	"lancer/model"
	request "lancer/request/rbac_menu"
	response "lancer/response/rbac_menu"
	"net/http"
	"unicode/utf8"
)

type RbacMenuService struct {
}

func NewRbacMenuService() *RbacMenuService {
	return &RbacMenuService{}
}

func (service *RbacMenuService) Create(req *request.RbacMenuCreateRequest) (ret *response.RbacMenuCreateResponse, err error) {
	ret = new(response.RbacMenuCreateResponse)

	//judge length
	length := utf8.RuneCount([]byte(req.Title))
	if length > 20 || length < 2 {
		return nil, errors.New(global.TranslateMessage["RbacMenuTitleLengthError"])
	}

	//judge repeat
	var count int64
	err = global.DB.Model(&model.LancerRbacMenu{}).Where("title = ?", req.Title).Count(&count).Error
	if err != nil {
		return nil, err
	}

	if count > 1 {
		return nil, errors.New(global.TranslateMessage["RbacMenuTitleRepeatError"])
	}

	rbacMenu := model.LancerRbacMenu{
		ID:        snowflake.Id(),
		ParentID:  req.ParentId,
		Title:     req.Title,
		Path:      req.Path,
		MenuType:  req.MenuType,
		Icon:      req.Icon,
		Sort:      req.Sort,
		Status:    req.Status,
		CreatedAt: time_tools.SerializerTimeNow(),
	}

	err = global.DB.Model(&model.LancerRbacMenu{}).Create(&rbacMenu).Error
	if err != nil {
		return nil, err
	}
	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["CreateSuccess"]

	return
}

func (service *RbacMenuService) List(req *request.RbacMenuListRequest) (ret *response.RbacMenuListResponse, err error) {
	ret = new(response.RbacMenuListResponse)

	//select all menu ,then using pointers to assemble data
	var rbacMenus []*model.LancerRbacMenu
	err = global.DB.Model(&model.LancerRbacMenu{}).Order("created_at ASC").Find(&rbacMenus).Error
	if err != nil {
		return nil, err
	}
	list := make([]*response.RbacMenuLineData, 0)
	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["FindSuccess"]

	if len(rbacMenus) == 0 {
		ret.Data.List = list
		return
	}

	var menus []*response.RbacMenuLineData
	for _, rbacMenu := range rbacMenus {
		menuLineData := response.RbacMenuLineData{
			LancerRbacMenu: *rbacMenu,
			Children:       nil,
		}

		menus = append(menus, &menuLineData)
	}

	for _, menu1 := range menus {
		for _, menu2 := range menus {
			if menu1.ID == menu2.ParentID {
				menu1.Children = append(menu1.Children, menu2)
			}
		}
	}

	for _, menuTree := range menus {
		if menuTree.ParentID == 0 {
			list = append(list, menuTree)
		}
	}

	ret.Data.List = list

	return
}

func (service *RbacMenuService) Update(req *request.RbacMenuUpdateRequest) (ret *response.RbacMenuUpdateResponse, err error) {
	ret = new(response.RbacMenuUpdateResponse)

	//judge length
	length := utf8.RuneCount([]byte(req.Title))
	if length > 20 || length < 2 {
		return nil, errors.New(global.TranslateMessage["RbacMenuTitleLengthError"])
	}

	updateMap := make(map[string]interface{})
	updateMap["title"] = req.Title
	updateMap["parent_id"] = req.ParentId
	updateMap["path"] = req.Path
	updateMap["menu_type"] = req.MenuType
	updateMap["icon"] = req.Icon
	updateMap["sort"] = req.Sort
	updateMap["status"] = req.Status

	err = global.DB.Model(&model.LancerRbacMenu{}).Where("id = ?", req.Id).Updates(updateMap).Error
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["UpdateSuccess"]
	return
}

func (service *RbacMenuService) Delete(req *request.RbacMenuDeleteRequest) (ret *response.RbacMenuDeleteResponse, err error) {
	ret = new(response.RbacMenuDeleteResponse)

	err = global.DB.Where("id = ?", req.Id).Delete(&model.LancerRbacMenu{}).Error
	if err != nil {
		return nil, err
	}
	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["DeleteSuccess"]
	return
}
