package we

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

// WxSubscribeMessageTemplate
type WxSubscribeMessageTemplate struct {
	PriTmplId string `json:"priTmplId"` //模板ID
	Title     string `json:"title"`     //模板标题
	Content   string `json:"content"`   //模板内容
	Example   string `json:"example"`   //模板示例
	Type      int32  `json:"type"`      //模版类型
}

type WxGetTemplateRes struct {
	common.WxCommonResponse
	Data []WxSubscribeMessageTemplate `json:"data"`
}

// WxSubscribeMessageParam 发送订阅消息
type WxSubscribeMessageParam struct {
	Touser           string                 `json:"touser"`                      //接收用户openid
	TemplateID       string                 `json:"template_id"`                 //消息模版ID
	Page             string                 `json:"page"`                        //点击后的跳转页面，仅限本小程序内的页面
	Data             map[string]interface{} `json:"data"`                        //模板内容
	MiniprogramState string                 `json:"miniprogram_state,omitempty"` //（小程序必填）跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Miniprogram      map[string]interface{} `json:"miniprogram,omitempty"`       //（公众号可选）跳转小程序时填写，格式如{ "appid": "pagepath": { "value": any } }
	Lang             string                 `json:"lang"`                        //进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

// GetSubscribeTemplateList 获取私有订阅模版
func (sdk *SDK) GetSubscribeTemplateList(ctx context.Context, appid string) (req *WxGetTemplateRes, err error) {
	req = &WxGetTemplateRes{}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token=%s", sdk.AccessToken)
	if err = common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	return req, nil
}

// SendSubscribeMessage 发送模版信息
func (sdk *SDK) SendSubscribeMessage(ctx context.Context, param *WxSubscribeMessageParam) error {

	if param.TemplateID == "" {
		return fmt.Errorf("template_id is empty")
	}

	bodyMap := make(common.BodyMap)
	bodyMap.Set("touser", param.Touser)
	bodyMap.Set("template_id", param.TemplateID)
	bodyMap.Set("data", param.Data)
	bodyMap.Set("lang", param.Lang)

	if param.Page != "" {
		bodyMap.Set("page", param.Page)
	}
	if len(param.Miniprogram) > 0 {
		bodyMap.Set("miniprogram", param.Miniprogram) //公众号
	} else {
		bodyMap.Set("miniprogram_state", param.MiniprogramState) //小程序
	}

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}
