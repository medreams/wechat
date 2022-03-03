package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

//WxSubscribeMessageTemplate
type WxSubscribeMessageTemplate struct {
	PriTmplId string `form:"priTmplId" json:"priTmplId"` //模板ID
	Title     string `form:"title" json:"title"`         //模板标题
	Content   string `form:"content" json:"content"`     //模板内容
	Example   string `form:"example" json:"example"`     //模板示例
	Type      int32  `form:"type" json:"type"`           //模版类型
}

type WxGetTemplateRes struct {
	Errmsg  string                       `form:"errmsg" json:"errmsg"`
	Errcode int32                        `form:"errcode" json:"errcode"`
	Data    []WxSubscribeMessageTemplate `form:"data" json:"data"`
}

//WxSendTemplateMessageParam 发送订阅消息
type WxSendTemplateMessageParam struct {
	Touser           string                 `form:"touser" json:"touser" binding:"required,min=2,max=30"`    //接收用户openid
	TemplateID       string                 `form:"template_id" json:"template_id" binding:"required,min=1"` //消息模版ID
	Page             string                 `form:"page" json:"page" binding:"min=4"`                        //点击后的跳转页面，仅限本小程序内的页面
	Data             map[string]interface{} `form:"data" json:"data" binding:"required"`                     //模板内容
	MiniprogramState string                 `form:"miniprogram_state" json:"miniprogram_state"`              //跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string                 `form:"lang" json:"lang"`                                        //进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

// GetSubscribeTemplateList 获取订阅模版
func (sdk *SDK) GetSubscribeTemplateList(ctx context.Context, appid string) (template *WxGetTemplateRes, err error) {
	template = &WxGetTemplateRes{}

	apiurl := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token=%s", sdk.AccessToken)

	if err = common.DoRequestGet(ctx, apiurl, template); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	return template, nil
}

// SendSubscribeMessage 发送模版信息
func (sdk *SDK) SendSubscribeMessage(ctx context.Context, param *WxSendTemplateMessageParam) error {

	if param.TemplateID == "" {
		return fmt.Errorf("template_id is empty")
	}

	bodyMap := make(map[string]interface{})

	bodyMap["access_token"] = sdk.AccessToken
	bodyMap["touser"] = param.Touser
	bodyMap["template_id"] = param.TemplateID
	bodyMap["page"] = param.Page
	bodyMap["data"] = param.Data
	bodyMap["miniprogram_state"] = param.MiniprogramState
	bodyMap["lang"] = param.Lang

	apiurl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", sdk.AccessToken)

	req := &common.WxCommonResponse{}
	if err := common.DoRequestPost(ctx, apiurl, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}
