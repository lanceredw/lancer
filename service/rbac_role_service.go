package service

import (
	"errors"
	"lancer/common/page"
	"lancer/common/snowflake"
	"lancer/common/time_tools"
	enum "lancer/enum/rbac_role"
	"lancer/global"
	"lancer/model"
	request "lancer/request/rbac_role"
	response "lancer/response/rbac_role"
	"net/http"
	"time"
	"unicode/utf8"
)

type RbacRoleService struct {
}

func NewRbacRoleService() *RbacRoleService {
	return &RbacRoleService{}
}

func (service *RbacRoleService) Create(req *request.RbacRoleCreateRequest) (ret *response.RbacRoleCreateResponse, err error) {
	ret = new(response.RbacRoleCreateResponse)

	length := utf8.RuneCount([]byte(req.RoleName))

	if length < 2 || length > 10 {
		return nil, errors.New(global.TranslateMessage["RbacRoleNameLengthError"])
	}

	var count int64

	err = global.DB.Model(&model.LancerRbacRole{}).Where("name = ?", req.RoleName).Count(&count).Error
	if err != nil {
		return nil, err
	}

	if count >= 1 {
		return nil, errors.New(global.TranslateMessage["RbacRoleNameRepeatError"])
	}

	role := model.LancerRbacRole{
		ID:        snowflake.Id(),
		RoleName:  req.RoleName,
		Status:    enum.RoleEnable,
		CreatedAt: time_tools.SerializerTimeNow(),
	}

	err = global.DB.Model(&model.LancerRbacRole{}).Create(&role).Error
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["CreateSuccess"]

	return
}

func (service *RbacRoleService) List(req *request.RbacRoleListRequest) (ret *response.RbacRoleListResponse, err error) {
	ret = new(response.RbacRoleListResponse)

	rbacRoleModel := global.DB.Model(&model.LancerRbacRole{})

	if req.RoleName != "" {
		rbacRoleModel.Where("role_name LIKE ?", "%"+req.RoleName+"%")
	}

	var count int64
	err = rbacRoleModel.Count(&count).Error
	if err != nil {
		return nil, err
	}

	list := make([]*model.LancerRbacRole, 0)
	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["FindSuccess"]
	if count == 0 {
		ret.Data.List = list
		return
	}

	pageResponse := page.Paging(count, req.Request)
	var roles []*model.LancerRbacRole
	err = rbacRoleModel.Order("created_at DESC").Scopes(page.Paginate(pageResponse)).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	list = roles

	ret.Data.List = list
	ret.Data.Count = count
	ret.Data.PageSize = pageResponse.PageSize
	ret.Data.PageIndex = pageResponse.PageIndex
	ret.Data.TotalPage = pageResponse.TotalPage

	return
}

func (service *RbacRoleService) Delete(req *request.RbacRoleDeleteRequest) (ret *response.RbacRoleDeleteResponse, err error) {
	ret = new(response.RbacRoleDeleteResponse)

	err = global.DB.Where("id = ?", req.Id).Delete(&model.LancerRbacRole{}).Error
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["DeleteSuccess"]

	return
}

func (service *RbacRoleService) Update(req *request.RbacRoleUpdateRequest) (ret *response.RbacRoleUpdateResponse, err error) {
	ret = new(response.RbacRoleUpdateResponse)

	length := utf8.RuneCount([]byte(req.RoleName))
	if length < 2 || length > 10 {
		return nil, errors.New(global.TranslateMessage["RbacRoleNameLengthError"])
	}

	var count int64
	err = global.DB.Model(&model.LancerRbacRole{}).Where("id = ? AND role_name = ?", req.Id, req.RoleName).Count(&count).Error
	if err != nil {
		return nil, err
	}

	if count >= 1 {
		return nil, errors.New(global.TranslateMessage["RbacRoleNameRepeatError"])
	}

	updateMap := make(map[string]interface{})
	updateMap["created_at"] = time.Now().Format(time.DateTime)
	updateMap["role_name"] = req.RoleName

	err = global.DB.Model(&model.LancerRbacRole{}).Where("id = ?", req.Id).Updates(updateMap).Error
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["UpdateSuccess"]

	return
}
