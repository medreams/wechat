package open

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

// WxSendTemplateMessageParam 发送订阅消息
type WxSendSubscribeMessageParam struct {
	Touser     string                 `json:"touser"`        //接收用户openid
	TemplateID string                 `json:"template_id"`   //消息模版ID
	Url        string                 `json:"url,omitempty"` //点击消息跳转的链接，需要有 ICP 备案
	Scene      string                 `json:"scene"`         //订阅场景值
	Title      string                 `json:"title"`         //消息标题，15 字以内
	Data       map[string]interface{} `json:"data"`          //模板内容
}

// SendSubscribeMessage 发送模版信息
// https://developers.weixin.qq.com/doc/oplatform/Mobile_App/One-time_subscription_info.html
func (sdk *SDK) SendSubscribeMessage(ctx context.Context, param *WxSendSubscribeMessageParam) error {

	if param.TemplateID == "" {
		return fmt.Errorf("template_id is empty")
	}

	bodyMap := make(common.BodyMap)
	bodyMap.Set("touser", param.Touser)
	bodyMap.Set("template_id", param.TemplateID)
	bodyMap.Set("scene", param.Scene)
	bodyMap.Set("title", param.Title)
	bodyMap.Set("data", param.Data)

	if param.Url != "" {
		bodyMap.Set("url", param.Url)
	}

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/subscribe?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}
