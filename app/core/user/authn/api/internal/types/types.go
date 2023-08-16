// Code generated by goctl. DO NOT EDIT.
package types

type QuickAuthSmsReq struct {
	MobileNo      string `json:"mobile_no"`                  // 手机号
	MobileCountry string `json:"mobile_country,default=+86"` // 国家代码: 默认值: +86
	Code          string `json:"code"`                       // 短信校验码
}

type QuickAuthEmailReq struct {
	Email string `json:"email"` // 邮箱
	Code  string `json:"code"`  // 校验码
}

type QuickAuthOAuthReq struct {
	OauthToken    string `json:"oauth_token"`    // oauth token
	OauthProvider string `json:"oauth_provider"` // oauth 平台: [wechat, twitter, facebook, google, github,]
}

type AuthResp struct {
	UserName  string `json:"username"`              // 用户名
	UserID    string `json:"user_id"`               // 用户 id
	Token     string `json:"token"`                 // 认证 token
	HasMfa    bool   `json:"has_mfa"`               // 是否有 MFA(Multi-Factor Authentication) 设备
	MfaType   uint32 `json:"mfa_type,range=[1:20]"` // MFA 设备类型: [0: 未设置, 1: google authenticator, 2: sms, 3: email]
	MfaTicket string `json:"mfa_ticket"`            // MFA ticket 有效期: 5分钟 (5 * 60), 用途: 二次认证 MFA code 时, 一起提交给后台, 作为用户识别+比对
}

type UserRegisterCheckReq struct {
	Email         string `json:"email,optional"`                      // 邮箱, 可选
	MobileNo      string `json:"mobile_no,optional"`                  // 手机号, 可选
	MobileCountry string `json:"mobile_country,optional,default=+86"` // 国家代码, 可选, 默认值: +86
	RegisterType  uint32 `json:"register_type,range=[1:10]"`          // 注册类型: [1: mobile, 2: email, 3: oauth]
}

type UserRegisterCheckResp struct {
	Status  int32  `json:"status"`  // 状态码: [1, 已经存在, 0, 不存在]
	Message string `json:"message"` // 状态信息
}

type UserRegisterReq struct {
	Email         string `json:"email"`          // 邮箱
	MobileNo      string `json:"mobile_no"`      // 手机号
	MobileCountry string `json:"mobile_country"` // 国家代码
	UserName      string `json:"user_name"`      // 用户名
	Password      string `json:"password"`       // 密码
	RegisterType  string `json:"register_type"`  // 注册类型: [mobile, email, username]
}

type UserRegisterResp struct {
	UserName string `json:"username"` // 用户名
	UserID   string `json:"user_id"`  // 用户 id
	Token    string `json:"token"`    // 认证 token
}

type UserLoginReq struct {
	Email         string `json:"email"`          // 邮箱
	MobileNo      string `json:"mobile_no"`      // 手机号
	MobileCountry string `json:"mobile_country"` // 国家代码
	UserName      string `json:"username"`       // 用户名
	UserNameSn    string `json:"username_sn"`    // 用户名编号: 类似暴雪游戏ID方案, 是 用户名+数字编号, 组合去重
	Password      string `json:"password"`       // 密码
	LoginType     string `json:"login_type"`     // 登录类型: [mobile, email, username]
}

type UserLogoutReq struct {
	Token string `json:"token"` // 认证 token
}

type UserLogoutResp struct {
	UserName string `json:"username"` // 用户名
	Message  string `json:"message"`  // 状态信息
	Status   int32  `json:"status"`   // 状态码: [1, 登出成功, -1, 登出失败]
}

type CaptchaReq struct {
	CaptchaID string `json:"captcha_id"` // 验证码 id
	Captcha   string `json:"captcha"`    // 验证码
}

type CaptchaResp struct {
	Status  int32  `json:"status"`  // 验证结果 状态码: [1, 验证成功, -1, 验证失败]
	Message string `json:"message"` // 状态信息
}

type MfaAddReq struct {
	Secret  string `json:"google_secret"`         // secret
	Code    string `json:"google_code"`           // code
	MfaType uint32 `json:"mfa_type,range=[1:10]"` // MFA 设备类型: [0: 未设置, 1: google authenticator, 2: sms, 3: email]}
}

type MfaAddResp struct {
	Status  int32  `json:"status"`  // 状态码: [1, 添加成功, -1, 添加失败]
	Message string `json:"message"` // 状态信息
}

type MfaVerifyReq struct {
	Code      string `json:"google_code"`           // code
	MfaTicket string `json:"mfa_ticket"`            // MFA ticket 有效期: 5分钟 (5 * 60), 用途: 二次认证 MFA code 时, 一起提交给后台, 作为用户识别+比对
	MfaType   uint32 `json:"mfa_type,range=[1:10]"` // MFA 设备类型: [0: 未设置, 1: google authenticator, 2: sms, 3: email]}
}

type MfaVerifyResp struct {
	Username string `json:"username"` // 用户名
	UserID   string `json:"user_id"`  // 用户 id
	Token    string `json:"token"`    // 认证 token
	Status   int32  `json:"status"`   // 状态码: [1, 验证成功, -1, 验证失败]
	Message  string `json:"message"`  // 状态信息
}

type MfaRemoveReq struct {
	Code    string `json:"code"`                  // mfa code
	MfaType uint32 `json:"mfa_type,range=[1:10]"` // MFA 设备类型: [0: 未设置, 1: google authenticator, 2: sms, 3: email]}
}

type MfaRemoveResp struct {
	Status  int32  `json:"status"`  // 状态码: [1, 移除成功, -1, 移除失败]
	Message string `json:"message"` // 状态信息
}

type SendEmailCodeReq struct {
	Email string `json:"email"` // 邮箱
	Code  string `json:"code"`  // 验证码
}

type SendSmsCodeReq struct {
	MobileNo      string `json:"mobile_no"`      // 手机号
	MobileCountry string `json:"mobile_country"` // 国家代码
	Code          string `json:"code"`           // 验证码
}

type SendCodeResp struct {
	Message string `json:"message"` // 状态信息
	Status  int32  `json:"status"`  // 状态码: [1, 发送成功, -1, 发送失败, -2 限速]
}