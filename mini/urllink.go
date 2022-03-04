package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type WxMiniLink struct {
	UrlLink string `json:"url_link,omitempty"`
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type WxMiniCloudBase struct {
	Env           string `json:"env"`
	Domain        string `json:"domain,omitempty"`
	Path          string `json:"path,omitempty"`
	Query         string `json:"query,omitempty"`
	ResourceAppid string `json:"resource_appid,omitempty"`
}

type WxMiniUrlLinkInfo struct {
	Appid      string           `json:"appid"`
	Path       string           `json:"path,omitempty"`
	Query      string           `json:"query,omitempty"`
	CreateTime int64            `json:"create_time,omitempty"`
	ExpireTime int64            `json:"expire_time,omitempty"`
	EnvVersion string           `json:"env_version,omitempty"`
	CloudBase  *WxMiniCloudBase `json:"cloud_base,omitempty"`
}

type WxMiniUrlLinkQuery struct {
	Errcode      int               `json:"errcode,omitempty"`
	Errmsg       string            `json:"errmsg,omitempty"`
	UrlLinkInfo  WxMiniUrlLinkInfo `json:"url_link_info,omitempty"`
	UrlLinkQuota WxMiniQuota       `json:"url_link_quota,omitempty"`
}

func (sdk *SDK) GenerateUrlLink(ctx context.Context, path, query, envVersion string, expire *WxMiniExpireParam, cb *WxMiniCloudBase) (scheme *WxMiniLink, err error) {

	bodyMap := make(map[string]interface{})
	bodyMap["access_token"] = sdk.AccessToken
	bodyMap["path"] = path
	bodyMap["query"] = query
	bodyMap["env_version"] = envVersion

	if expire != nil {
		bodyMap["is_expire"] = expire.IsExpire     //到期失效：true，永久有效：false。注意，永久有效 Link 和有效时间超过180天的到期失效 Link 的总数上限为10万个，详见获取 URL Link，生成 Link 前请仔细确认。
		bodyMap["expire_type"] = expire.ExpireType //小程序 URL Link 失效类型，失效时间：0，失效间隔天数：1
		switch expire.ExpireType {
		case 0:
			bodyMap["expire_time"] = expire.ExpireTime //到期失效的URL Link的失效时间，UNIX 时间戳，单位：秒。expire_type 为 0 必填
		case 1:
			bodyMap["expire_interval"] = expire.ExpireInterval //到期失效的URL Link的失效间隔天数。生成的到期失效URL Link在该间隔时间到达前有效。最长间隔天数为365天。expire_type 为 1 必填
		}

	} else {
		bodyMap["is_expire"] = true
		bodyMap["expire_type"] = 1
		bodyMap["expire_interval"] = 30
	}

	if cb != nil {
		bodyMap["cloud_base"] = *cb
	}

	req := &WxMiniLink{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/generate_urllink?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get wxMiniLink: %w", err)
	}

	return req, nil

}
func (sdk *SDK) QueryUrlLink(ctx context.Context, urlLink string) (query *WxMiniUrlLinkQuery, err error) {

	bodyMap := make(map[string]interface{})
	bodyMap["access_token"] = sdk.AccessToken
	bodyMap["url_link"] = urlLink

	req := &WxMiniUrlLinkQuery{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/query_urllink?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get wxMiniUrlLinkQuery: %w", err)
	}

	return req, nil
}
