package we

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
	"github.com/medreams/wechat/official"
)

//统一发送模版信息,因为小程序模版信息下线，所以这算是用小程序发公众号模版信息的一个渠道

// 公众号模版信息结构
type MpTemplateMsgParam struct {
	Appid       string                 `json:"appid"`       //公众号appid，要求与小程序有绑定且同主体
	TemplateId  string                 `json:"template_id"` //公众号模板id
	Url         string                 `json:"url"`         //公众号模板消息所要跳转的url
	Miniprogram official.WxMiniprogram `json:"miniprogram"` //公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
	Data        map[string]interface{} `json:"data"`        //公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
}

// 下发统一消息参数
type WxSendUniformTemplateMessageParam struct {
	Touser        string             `json:"touser"` //用户openid，可以是小程序的openid，也可以是mp_template_msg.appid对应的公众号的openid
	MpTemplateMsg MpTemplateMsgParam `json:"mp_template_msg"`
}

// SendUniformTemplateMessage 下发统一消息
func (sdk *SDK) SendUniformTemplateMessage(ctx context.Context, param *WxSendUniformTemplateMessageParam) error {

	if param.Touser == "" {
		return fmt.Errorf("touser is empty")
	}

	bodyMap := make(map[string]interface{})
	bodyMap["touser"] = param.Touser
	bodyMap["mp_template_msg"] = param.MpTemplateMsg

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}
