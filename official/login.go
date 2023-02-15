package official

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/medreams/wechat/common"
)

type WxWebAccessToekn struct {
	common.WxCommonResponse
	Openid       string `json:"openid,omitempty"`        // 用户唯一标识
	AccessToken  string `json:"access_token,omitempty"`  // 会话密钥
	ExpiresIn    int    `json:"expires_in,omitempty"`    // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	ExpiresTime  int64  `json:"expires_time,omitempty"`  // 凭证过期时间
	RefreshToken string `json:"refresh_token,omitempty"` // 用户刷新access_token
	Scope        string `json:"scope,omitempty"`         //用户授权的作用域，使用逗号（,）分隔
}

// 网页授权和开放平台网页code获取access_token(此access_token,只能在网页授权和开放平台网页中使用)
func (s *SDK) Code2WebAccessToken(ctx context.Context, code string) (req *WxWebAccessToekn, err error) {
	req = &WxWebAccessToekn{}
	URL := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", s.Appid, s.Secret, code)

	if err = common.DoRequestGet(ctx, URL, req); err != nil {
		return nil, fmt.Errorf("do request get access_token: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("get access_token error: %d, %s", req.ErrCode, req.ErrMsg)
	}

	req.ExpiresTime = time.Now().Unix() + int64(req.ExpiresIn)
	//s.UpdateAccessToken(at.AccessToken)

	return req, nil
}

func (s *SDK) RefreshAccessToken(ctx context.Context, refreshToken string) (req *WxWebAccessToekn, err error) {

	req = &WxWebAccessToekn{}
	URL := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?grant_type=refresh_token&appid=%s&refresh_token=%s", s.Appid, refreshToken)

	if err = common.DoRequestGet(ctx, URL, req); err != nil {
		return nil, fmt.Errorf("do request get access_token: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("get access_token error: %d, %s", req.ErrCode, req.ErrMsg)
	}

	req.ExpiresTime = time.Now().Unix() + int64(req.ExpiresIn)

	return req, nil
}

// 公众号网页授权获取的access_token拉取用户信息
func (s *SDK) WebAccessTokenAndOpenid2UserInfo(ctx context.Context, webAccessToken string, openid string) (req *UserInfo, err error) {
	req = &UserInfo{}

	URL := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", webAccessToken, openid)

	if err = common.DoRequestGet(ctx, URL, req); err != nil {
		return nil, fmt.Errorf("do request get mp userinfo: %w", err)
	}

	return req, nil
}

// 公众号网页获取授权code
func (s *SDK) WebLoginUrl(ctx context.Context, redirectUri string, state string, scope string) string {
	if scope == "" {
		scope = "snsapi_base"
	}
	encodURL := url.QueryEscape(redirectUri)
	uri := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s&lang=zh_CN#wechat_redirect", s.Appid, encodURL, scope, state)

	return uri
}
