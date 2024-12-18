package common

import (
	"context"
	"fmt"
	"time"
)

type WxAccessToken struct {
	AccessToken string `json:"access_token,omitempty"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in,omitempty"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	ExpiresTime int64  `json:"expires_time,omitempty"` // 凭证过期时间
	ErrCode     int    `json:"errcode,omitempty"`      // 错误码
	ErrMsg      string `json:"errmsg,omitempty"`       // 错误信息
}

func GetAccessToken(ctx context.Context, appid, appSecret string) (at *WxAccessToken, err error) {

	at = &WxAccessToken{}
	URL := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appid, appSecret)

	if err = DoRequestGet(ctx, URL, at); err != nil {
		return nil, fmt.Errorf("do request get access_token: %w", err)
	}

	if at.ErrCode != 0 {
		return nil, fmt.Errorf("get access_token error: %d, %s", at.ErrCode, at.ErrMsg)
	}

	at.ExpiresTime = time.Now().Unix() + int64(at.ExpiresIn-200)

	return at, nil
}

func GetStableToken(ctx context.Context, appid, appSecret string) (at *WxAccessToken, err error) {

	at = &WxAccessToken{}
	URL := "https://api.weixin.qq.com/cgi-bin/stable_token"

	bodyMap := make(BodyMap)
	bodyMap.Set("grant_type", "client_credential")
	bodyMap.Set("appid", appid)
	bodyMap.Set("secret", appSecret)
	bodyMap.Set("force_refresh", false)

	if err = DoRequestPost(ctx, URL, bodyMap, at); err != nil {
		return nil, fmt.Errorf("do request get access_token: %w", err)
	}

	if at.ErrCode != 0 {
		return nil, fmt.Errorf("get access_token error: %d, %s", at.ErrCode, at.ErrMsg)
	}

	at.ExpiresTime = time.Now().Unix() + int64(at.ExpiresIn)

	return at, nil
}
