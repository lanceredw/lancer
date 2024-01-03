package request

type RbacMenuUpdateRequest struct {
	Id       int64  `json:"id,string" form:"id,string" binding:"required"` //id
	ParentId int64  `json:"parent_id,string" form:"parent_id,string"`      //parent id
	Title    string `json:"title" form:"title" `                           //menu title
	Path     string `json:"path" form:"path"`                              //menu path
	MenuType int32  `json:"menu_type" form:"menu_type"`                    //menu type 1.menu 2.button
	Icon     string `json:"icon" form:"icon"`                              //menu ico
	Sort     int32  `json:"sort" form:"sort"`                              //menu sort
	Status   int32  `json:"status" form:"status"`                          //status 0.disable 1.enable
}
