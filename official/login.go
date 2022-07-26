package official

import (
	"context"
	"fmt"
	"time"

	"github.com/medreams/wechat/common"
)

type WxWebAccessToekn struct {
	Openid       string `json:"openid,omitempty"`        // 用户唯一标识
	AccessToken  string `json:"access_token,omitempty"`  // 会话密钥
	ExpiresIn    int    `json:"expires_in,omitempty"`    // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	ExpiresTime  int64  `json:"expires_time,omitempty"`  // 凭证过期时间
	RefreshToken string `json:"refresh_token,omitempty"` // 用户刷新access_token
	Scope        string `json:"scope,omitempty"`         //用户授权的作用域，使用逗号（,）分隔
	Errcode      int    `json:"errcode,omitempty"`       // 错误码
	Errmsg       string `json:"errmsg,omitempty"`        // 错误信息
}

//公众帐号code 获取 access_token(此access_token,只能在网页授权中使用)
func (s *SDK) Code2WebAccessToken(ctx context.Context, code string) (at *WxWebAccessToekn, err error) {
	at = &WxWebAccessToekn{}
	URL := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", s.Appid, s.Secret, code)

	if err = common.DoRequestGet(ctx, URL, at); err != nil {
		return nil, fmt.Errorf("do request get access_token: %w", err)
	}

	if at.Errcode != 0 {
		return nil, fmt.Errorf("get access_token error: %d, %s", at.Errcode, at.Errmsg)
	}

	at.ExpiresTime = time.Now().Unix() + int64(at.ExpiresIn)
	//s.UpdateAccessToken(at.AccessToken)

	return at, nil
}

//公众号网页授权获取的access_token拉取用户信息
func (s *SDK) WebAccessTokenAndOpenid2UserInfo(ctx context.Context, webAccessToken string, openid string) (userinfo *PublicUserInfo, err error) {
	userinfo = &PublicUserInfo{}

	URL := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", webAccessToken, openid)

	if err = common.DoRequestGet(ctx, URL, userinfo); err != nil {
		return nil, fmt.Errorf("do request get mp userinfo: %w", err)
	}

	return userinfo, nil
}
