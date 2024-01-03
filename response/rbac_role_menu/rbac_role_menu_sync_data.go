package response

import (
	"lancer/model"
	"lancer/response"
)

type RbacRoleMenuSyncDataResponse struct {
	response.ResultData
}

type RoleMenuLineData struct {
	model.LancerRbacRoleMenu
	Menu *model.LancerRbacMenu `json:"menus" gorm:"foreignKey:ID;references:MenuID"`
}
