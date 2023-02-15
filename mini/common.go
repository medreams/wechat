package mini

type SDK struct {
	Appid       string
	Secret      string
	AccessToken string
}

type WxMiniQuota struct {
	LongTimeUsed  int `json:"long_time_used,omitempty"`
	LongTimeLimit int `json:"long_time_limit,omitempty"`
}

// 微信小程序链接过期参数
type WxMiniExpireParam struct {
	IsExpire       bool  `form:"is_expire" json:"is_expire" binding:"required"`             //是否过期
	ExpireType     int32 `form:"expire_type" json:"expire_type" binding:"required"`         //过期类型
	ExpireTime     int32 `form:"expire_time" json:"expire_time" binding:"required"`         //过期时间
	ExpireInterval int32 `form:"expire_interval" json:"expire_interval" binding:"required"` //过期间隔
}

func New(appid, secret, token string) *SDK {
	return &SDK{
		Appid:       appid,
		Secret:      secret,
		AccessToken: token,
	}
}
