package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type WxMiniJumpWxa struct {
	Path       string
	Query      string
	EnvVersion string
}

type WxMiniScheme struct {
	OpenLink string `json:"openlink,omitempty"`
	Errcode  int    `json:"errcode,omitempty"`
	Errmsg   string `json:"errmsg,omitempty"`
}

type WxMiniSchemeQuery struct {
	Errcode     int              `json:"errcode,omitempty"`
	Errmsg      string           `json:"errmsg,omitempty"`
	SchemeInfo  WxMiniSchemeInfo `json:"scheme_info,omitempty"`
	SchemeQuota WxMiniQuota      `json:"scheme_quota,omitempty"`
}

type WxMiniSchemeInfo struct {
	Appid      string `json:"appid,omitempty"`
	Path       string `json:"path,omitempty"`
	Query      string `json:"query,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	ExpireTime int64  `json:"expire_time,omitempty"`
	EnvVersion string `json:"env_version,omitempty"`
}

type WxMiniQuota struct {
	LongTimeUsed  int `json:"long_time_used,omitempty"`
	LongTimeLimit int `json:"long_time_limit,omitempty"`
}

func (sdk *SDK) GenerateScheme(ctx context.Context, expireType int, expireInterval int, jw *WxMiniJumpWxa) (scheme *WxMiniScheme, err error) {

	isExpire := true //生成的 scheme 码类型，到期失效：true，永久有效：false。注意，永久有效 scheme 和有效时间超过180天的到期失效 scheme 的总数上限为10万个
	// expireType = 1   //，失效时间：0，失效间隔天数：1

	// expireInterval = 30 //到期失效的 scheme 码的失效间隔天数。生成的到期失效 scheme 码在该间隔时间到达前有效。最长间隔天数为365天。is_expire 为 true 且 expire_type 为 1 时必填

	bodyMap := make(map[string]interface{})
	bodyMap["access_token"] = sdk.AccessToken
	bodyMap["is_expire"] = isExpire
	bodyMap["expire_type"] = expireType
	bodyMap["expire_interval"] = expireInterval

	req := &WxMiniScheme{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/generatescheme?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get wxMiniScheme: %w", err)
	}

	return req, nil

}
func (sdk *SDK) QueryScheme(ctx context.Context, scheme string) (query *WxMiniSchemeQuery, err error) {

	bodyMap := make(map[string]interface{})
	bodyMap["access_token"] = sdk.AccessToken
	bodyMap["scheme"] = scheme

	req := &WxMiniSchemeQuery{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/queryscheme?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get wxMiniSchemeQuery: %w", err)
	}

	return req, nil
}
