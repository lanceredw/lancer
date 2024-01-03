package translations

var MessagesZh = map[string]string{

	//for public use
	"CreateSuccess":  "创建成功",
	"UpdateSuccess":  "更新成功",
	"DeleteSuccess":  "删除成功",
	"FindSuccess":    "查询成功",
	"AddSuccess":     "添加成功",
	"ProcessSuccess": "处理成功",

	//middleware
	"AuthHeaderEmptyError": "header头为空",

	//user controller
	"UserUserNameRepeatError":        "账户名重复",
	"UserUpdateNickNameLengthError":  "昵称长度在2-10",
	"UserUpdateEmailError":           "邮箱格式错误",
	"UserLoginUserNameNotExistError": "账户不存在",
	"UserLoginPasswordError":         "密码错误",
	"UserLoginSuccess":               "登录成功",

	//role controller
	"RbacRoleNameLengthError": "角色名长度在2-10",
	"RbacRoleNameRepeatError": "角色名重复",

	//menu controller
	"RbacMenuTitleLengthError": "菜单名长度在2-20",
	"RbacMenuTitleRepeatError": "菜单名重复",

	//roleMenu controller
	"RbacRoleMenuRecordRepeatError": "重复添加",
	"RbacRoleMenuNoRoleError":       "没有角色",
}
