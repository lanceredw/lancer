package service

import (
	"errors"
	"gorm.io/gorm"
	"lancer/common/method"
	"lancer/common/page"
	"lancer/common/snowflake"
	"lancer/common/time_tools"
	enum "lancer/enum/user"
	"lancer/global"
	"lancer/model"
	request "lancer/request/user"
	response "lancer/response/user"
	"net/http"
	"time"
	"unicode/utf8"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (service *UserService) Create(req *request.UserCreateRequest) (ret *response.UserCreateResponse, err error) {
	ret = new(response.UserCreateResponse)

	//judge some user name
	var count int64
	err = global.DB.Model(&model.LancerUser{}).Where("user_name = ?", req.UserName).Count(&count).Error
	if err != nil {
		return nil, err
	}

	if count >= 1 {
		return nil, errors.New(global.TranslateMessage["UserCreateUserNameRepeatError"])
	}

	salt := method.GenerateRandomString(4)

	password := method.Md5(req.Password + salt)

	user := model.LancerUser{
		ID:        snowflake.Id(),
		UserName:  req.UserName,
		NickName:  "",
		Avatar:    "",
		RoleID:    0,
		Password:  password,
		Salt:      salt,
		Sex:       0,
		Email:     "",
		Phone:     "",
		Status:    0,
		CreatedAt: time_tools.SerializerTimeNow(),
	}

	err = global.DB.Model(&model.LancerUser{}).Create(&user).Error
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Data = global.TranslateMessage["CreateSuccess"]

	return
}

func (service *UserService) Delete(req *request.UserDeleteRequest) (ret *response.UserDeleteResponse, err error) {
	ret = new(response.UserDeleteResponse)

	err = global.DB.Where("id = ?", req.Id).Delete(&model.LancerUser{}).Error
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["DeleteSuccess"]

	return
}

func (service *UserService) Update(req *request.UserUpdateRequest) (ret *response.UserUpdateResponse, err error) {
	ret = new(response.UserUpdateResponse)

	updateMap := make(map[string]interface{})
	updateMap["created_at"] = time.Now().Format(time.DateTime)

	//check req
	if req.NickName != "" {
		length := utf8.RuneCount([]byte(req.NickName))
		if length > 10 || length < 2 {
			return nil, errors.New(global.TranslateMessage["UserUpdateNickNameLengthError"])
		}
		updateMap["nick_name"] = req.NickName
	}

	if req.Avatar != "" {
		updateMap["avatar"] = req.Avatar
	}

	if req.RoleID != 0 {
		updateMap["role_id"] = req.RoleID
	}

	if (req.Sex != enum.UnKnowSex) && (req.Sex == enum.Male || req.Sex == enum.Female) {
		updateMap["sex"] = req.Sex
	}

	if req.Email != "" {
		if !method.IsEmail(req.Email) {
			return nil, errors.New(global.TranslateMessage["UserUpdateEmailError"])
		}
		updateMap["email"] = req.Email
	}

	if req.Phone != "" {
		updateMap["phone"] = req.Phone
	}

	if req.Status == enum.UserDisable || req.Status == enum.UserEnable {
		updateMap["status"] = req.Status
	}

	err = global.DB.Model(&model.LancerUser{}).Where("id = ?", req.Id).Updates(updateMap).Error
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["UpdateSuccess"]

	return
}

func (service *UserService) List(req *request.UserListRequest) (ret *response.UserListResponse, err error) {
	ret = new(response.UserListResponse)

	userModel := global.DB.Model(&model.LancerUser{})

	if req.UserName != "" {
		userModel.Where("user_name LIKE ?", "%"+req.UserName+"%")
	}

	if req.Status != -1 {
		userModel.Where("status = ?", req.Status)
	}

	var count int64
	err = userModel.Count(&count).Error
	if err != nil {
		return nil, err
	}

	list := make([]*model.LancerUser, 0)
	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["FindSuccess"]
	if count == 0 {
		ret.Data.List = list
		return
	}

	pageResponse := page.Paging(count, req.Request)
	var users []*model.LancerUser
	err = userModel.Order("created_at DESC").Scopes(page.Paginate(pageResponse)).Find(&users).Error
	if err != nil {
		return nil, err
	}

	//TODO for range users to list

	list = users

	ret.Data.List = list
	ret.Data.Count = count
	ret.Data.PageSize = pageResponse.PageSize
	ret.Data.PageIndex = pageResponse.PageIndex
	ret.Data.TotalPage = pageResponse.TotalPage

	return
}

func (service *UserService) Login(req *request.UserLoginRequest) (ret *response.UserLoginResponse, err error) {
	ret = new(response.UserLoginResponse)

	//select user name
	var user *model.LancerUser
	err = global.DB.Model(&model.LancerUser{}).Where("user_name = ?", req.UserName).Order("created_at DESC").First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(global.TranslateMessage["UserLoginUserNameNotExistError"])
		}
		return nil, err
	}

	p := method.Md5(req.Password + user.Salt)
	if user.Password != p {
		return nil, errors.New(global.TranslateMessage["UserLoginPasswordError"])
	}

	jwt, err := method.GenerateJwt(user.ID, user.RoleID, user.UserName)
	if err != nil {
		return nil, err
	}

	//register redis expire

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["UserLoginSuccess"]
	ret.Data.Token = jwt
	return
}
