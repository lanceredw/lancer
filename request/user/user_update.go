package request

type UserUpdateRequest struct {
	Id       int64  `json:"id,string" form:"id,string" label:"id" binding:"required"` //id
	NickName string `json:"nick_name" form:"nick_name" label:"nick_name"`             //nickname
	Avatar   string `json:"avatar" form:"avatar" label:"avatar"`                      //avatar url
	RoleID   int64  `json:"role_id" form:"role_id" label:"role"`                      //role id
	Sex      int32  `json:"sex" form:"sex" label:"sex"`                               //sex 1.male 2.female
	Email    string `json:"email" form:"email" label:"email"`                         // email
	Phone    string `json:"phone" form:"phone" label:"phone"`                         //phone
	Status   int32  `json:"status" form:"status" label:"status"`                      //status 1.enable 2.disable
}
