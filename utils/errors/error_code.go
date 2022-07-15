package errors

type Error struct {
	Code    int
	Message string
}

var (
	Success       = &Error{0, "ok"}
	SuccessSms    = &Error{0, "发送成功"}
	SuccessCustom = &Error{0, "自定义成功提示"}

	NoLogin = &Error{101, "未登录"}

	// 其他系统错误
	System         = &Error{500, "系统错误，请联系管理员"}
	DatabaseClient = &Error{501, "数据库连接超时"}

	// 表单或参数错误 1100 - 1199
	ParamsFailed          = &Error{1100, "参数错误"}
	MobileFailed          = &Error{1101, "手机号错误"}
	CreateTokenFailed     = &Error{1102, "token生成失败"}
	RealNameFailed        = &Error{1103, "姓名不能为空"}
	RolesFailed           = &Error{1104, "请选择角色"}
	EmptyName             = &Error{1105, "名称不能为空"}
	EmptyRegion           = &Error{1105, "请选择地区"}
	EmptyThumbImg         = &Error{1105, "缩略图不能为空"}
	EmptyMenus            = &Error{1106, "请选择权限菜单"}
	EmptyMenuMethod       = &Error{1107, "请输入菜单method"}
	EmptyMerchantName     = &Error{1108, "请输入商户名称"}
	EmptyMerchantLogo     = &Error{1109, "请选择商户logo"}
	EmptyMerchantCode     = &Error{1110, "请输入商户代号"}
	EmptyRegionCode       = &Error{1111, "请选择地区"}
	EmptyAddress          = &Error{1112, "请填写详细地址"}
	EmptyLogLat           = &Error{1113, "请选择定位"}
	UploadIllegalFileType = &Error{1131, "上传非法文件类型"}
	UploadCreateDirFailed = &Error{1132, "上传目录创建失败"}
	UploadFailed          = &Error{1133, "上传失败"}

	// 操作错误 1200 - 1299
	NoPermission         = &Error{1200, "暂无该权限"}
	ChangeMerchantFailed = &Error{1201, "切换商户失败"}
	SuperAdminNoDel      = &Error{1202, "超管账户不允许删除"}
	SaveFailed           = &Error{1203, "保存失败"}
	DeleteFailed         = &Error{1204, "删除失败"}
	DataNoFound          = &Error{1205, "找不到数据"}
	OperationFailed      = &Error{1206, "操作失败"}
	AdminUserAlreadyBind = &Error{1207, "该会员已被绑定为管理员，不允许重复绑定"}
	UserNoFound          = &Error{1208, "会员不存在"}
	EmptyUserUnionid     = &Error{1209, "该用户缺少unionid，请先让用户登录一遍小程序再操作"}
	ChangeFailed         = &Error{1210, "变更失败"}
	MerchantNOImport     = &Error{1211, "该商户下不能导入"}
	MerchantNODelete     = &Error{1212, "该商户下不可删除"}

	// 短信错误 1300 - 1349
	SmsSignName         = &Error{1300, "短信签名错误"}
	SmsSendFailed       = &Error{1301, "短信发送失败"}
	SmsCodeCancelFailed = &Error{1302, "短信发送失败"}
	SmsCodeFailed       = &Error{1303, "短信验证码错误"}

	// 业务错误 1350 - 1399
	CreateAppletsQrcodeFailed = &Error{1350, "生成小程序码失败"}

	// 管理员错误 1400 - 1449
	AdminNoExists = &Error{1400, "管理员不存在"}
	AdminDisabled = &Error{1401, "该帐号已被禁用"}
	AdminAlready  = &Error{1402, "该帐号已经是管理员了"}
)
