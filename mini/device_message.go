package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type WxSendHardwareDeviceMessageParam struct {
	Sn               string                 `json:"sn"`                          //设备唯一序列号。由厂商分配，长度不能超过128字节。字符只接受数字，大小写字母，下划线（_）和连字符（-）。
	TemplateID       string                 `json:"template_id"`                 //消息模版ID
	ToOpenidList     []string               `json:"to_openid_list"`              //接收者（用户）的 openid 列表
	MiniprogramState string                 `json:"miniprogram_state,omitempty"` //（小程序必填）跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	ModelId          string                 `json:"modelId"`                     //设备型号 id ，通过注册设备获得。
	Data             map[string]interface{} `json:"data"`                        //模板内容
}

// 发送设备消息 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/sendHardwareDeviceMessage.html
func (sdk *SDK) SendHardwareDeviceMessage(ctx context.Context, param *WxSendHardwareDeviceMessageParam) error {
	if param.TemplateID == "" {
		return fmt.Errorf("template_id is empty")
	}

	bodyMap := make(common.BodyMap)
	bodyMap.Set("sn", param.Sn)
	bodyMap.Set("template_id", param.TemplateID)
	bodyMap.Set("to_openid_list", param.ToOpenidList)
	bodyMap.Set("modelId", param.ModelId)
	bodyMap.Set("data", param.Data)

	if len(param.MiniprogramState) > 0 {
		bodyMap["miniprogram_state"] = param.MiniprogramState //小程序
	}

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/device/subscribe/send?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

type GetSnTicketParam struct {
	Sn      string `json:"sn"`
	ModelId string `json:"model_id"`
}

type GetSnTicketRsp struct {
	common.WxCommonResponse
	SnTicket string `json:"sn_ticket"`
}

// 获取设备票据 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/hardware-device/getSnTicket.html
func (sdk *SDK) GetSnTicket(ctx context.Context, param *GetSnTicketParam) (*GetSnTicketRsp, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("sn", param.Sn)
	bodyMap.Set("model_id", param.ModelId)

	req := &GetSnTicketRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/getsnticket?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}
