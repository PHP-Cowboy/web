package ecode

var (
	Success                  = New(200, "ok") // 正确
	RequestError             = New(400, "请求错误")
	IllegalRequest           = New(401, "非法请求")
	RequestNotFound          = New(404, "请求不存在")
	ServerErr                = New(500, "服务器错误")
	DataInsertError          = New(501, "数据添加失败")
	DataDeleteError          = New(502, "数据删除失败")
	DataSaveError            = New(503, "数据保存失败")
	DataQueryError           = New(504, "数据查询失败")
	DataNotExist             = New(505, "数据不存在")
	DataAlreadyExist         = New(506, "数据已存在")
	DataCannotBeModified     = New(507, "数据不允许修改")
	GetContextUserInfoFailed = New(508, "获取上下文用户数据失败")
	DataQueryTimeOutOfRange  = New(599, "超出查询时间范围")

	CommunalSignInvalid    = New(10, "sign参数异常")
	ParamInvalid           = New(11, "参数不合法")
	CommunalSessionInvalid = New(12, "session参数异常")
	CommunalParamInvalid   = New(13, "公共参数异常")
	UserNotLogin           = New(14, "用户未登录")
	TokenExpired           = New(15, "token已过期")
	RedisKeyDelFailed      = New(16, "redis key del failed")
	CreateCaptchaError     = New(17, "生成验证码错误")

	UserNotFound            = New(1000, "用户未找到")
	RoleNotFound            = New(1001, "角色未找到")
	DataTransformationError = New(1002, "数据转换出错")
	PasswordCheckFailed     = New(1003, "密码校验有误")
	WarehouseNotFound       = New(1004, "仓库未找到")
	RedisFailedToGetData    = New(1005, "redis获取数据失败")
	RedisFailedToSetData    = New(1006, "redis设置数据失败")
	NoOrderFound            = New(1007, "未查询到订单")
	MapKeyNotExist          = New(1008, "map key not exist")
	OrderPickingCompleted   = New(1009, "订单已拣货完成")
	OrderNotToBeReviewed    = New(1010, "订单不是待复核状态，无法复核")
	OrderDataNotFound       = New(1011, "订单数据未找到")

	WarehouseSelectError = New(2000, "仓库选择有误，请重试")
	GetWarehouseFailed   = New(2001, "获取仓库数据失败")

	InvTaskNotGoing    = New(3000, "盘点任务不在进行中")
	InvTaskAlreadyBind = New(3001, "盘点任务已被绑定")
	UnBindTaskNotAllow = New(3002, "未绑定盘点任务的不允许复盘")

	RegularOrExpressDeliveryBatchInvalid = New(4000, "订单创建常规或快递批次选择不正确")
)
