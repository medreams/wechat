package official

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type GetApiDomainIpRsp struct {
	common.WxCommonResponse
	IPList []string `json:"ip_list"`
}

// 获取微信服务器 IP 地址  https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (sdk *SDK) GetApiDomainIp(ctx context.Context) ([]string, error) {
	req := &GetApiDomainIpRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/get_api_domain_ip?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request get access_token: %w", err)
	}
	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req.IPList, nil
}

// 获取微信callback IP地址
func (sdk *SDK) GetCallbackDomainIp(ctx context.Context) ([]string, error) {
	req := &GetApiDomainIpRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request get access_token: %w", err)
	}
	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req.IPList, nil
}
