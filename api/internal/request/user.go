package request

type Register struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
	SmsCode  string `form:"sms_code" binding:"required"`
}
type Login struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
	SmsCode  string `form:"sms_code" binding:"required"`
}
type SendSms struct {
	Phone  string `form:"phone" binding:"required"`
	Source string `form:"source" binding:"required"`
}
type UserCenterAdd struct {
	Name          string `form:"name" binding:"required"`          //名称
	NickName      string `form:"nickName" binding:"required"`      //昵称
	UserCode      string `form:"userCode" binding:"required"`      //编号
	Signature     string `form:"signature" binding:"required"`     //签名
	Sex           string `form:"sex" binding:"required"`           //性别
	IpAddress     string `form:"ipAddress" binding:"required"`     //ip地址
	Constellation string `form:"constellation" binding:"required"` //星座
	AttendCount   int64  `form:"attendCount" binding:"required"`   //关注数
	FanCount      int64  `form:"fanCount" binding:"required"`      //粉丝数
	ZanCount      int64  `form:"zanCount" binding:"required"`      //点赞数
	Status        string `form:"status" binding:"required"`        //用户状态
	ImageId       int64  `form:"imageId" binding:"required"`       //头像
	Info          string `form:"info" binding:"required"`          //认证信息
	Mobile        string `form:"mobile" binding:"required"`        //手机号
	RealNameAuth  string `form:"realNameAuth" binding:"required"`  //实名认证状态
	Age           int64  `form:"age" binding:"required"`           //年龄
	OnlineStatus  string `form:"onlineStatus" binding:"required"`  //在线状态
	Type          string `form:"type" binding:"required"`          //认证类型
	Level         string `form:"level" binding:"required"`         //用户等级
	Balance       string `form:"balance" binding:"required"`       //余额
}
