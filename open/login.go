package open

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
func (sdk *SDK) Code2WebAccessToken(ctx context.Context, code string) (req *WxWebAccessToekn, err error) {
	req = &WxWebAccessToekn{}
	URL := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", sdk.Appid, sdk.Secret, code)

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

func (sdk *SDK) RefreshAccessToken(ctx context.Context, refreshToken string) (req *WxWebAccessToekn, err error) {
	req = &WxWebAccessToekn{}
	URL := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?grant_type=refresh_token&appid=%s&refresh_token=%s", sdk.Appid, refreshToken)

	if err = common.DoRequestGet(ctx, URL, req); err != nil {
		return nil, fmt.Errorf("do request get access_token: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("get access_token error: %d, %s", req.ErrCode, req.ErrMsg)
	}

	req.ExpiresTime = time.Now().Unix() + int64(req.ExpiresIn)

	return req, nil
}

// 网站应用微信登录
func (sdk *SDK) WebLoginUrl(ctx context.Context, redirectUri string, state string) string {
	encodURL := url.QueryEscape(redirectUri)
	uri := fmt.Sprintf("https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_login&state=%s&lang=zh_CN#wechat_redirect", sdk.Appid, encodURL, state)

	return uri
}
