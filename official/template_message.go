package official

import (
	"context"
	"fmt"
	"strconv"

	"github.com/medreams/wechat/common"
)

// WxMessageTemplate
type WxMessageTemplate struct {
	TemplateId      string `json:"template_id"`      //模板ID
	Title           string `json:"title"`            //模板标题
	PrimaryIndustry string `json:"primary_industry"` //模板所属行业的一级行业
	DeputyIndustry  string `json:"deputy_industry"`  //模板所属行业的二级行业
	Content         string `json:"content"`          //模板内容
	Example         string `json:"example"`          //模板示例

}

type WxSendTemplateMessageRes struct {
	common.WxCommonResponse
	Msgid int64 `json:"msgid"`
}

type WxGetTemplateRes struct {
	common.WxCommonResponse
	TemplateList []WxMessageTemplate `json:"template_list"`
}

// 小程序跳转参数
type WxMiniprogram struct {
	Appid    string `json:"appid"`              //所需跳转到的小程序appid（该小程序 appid 必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	Pagepath string `json:"pagepath,omitempty"` //所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}

// WxSendTemplateMessageParam 发送订阅消息
type WxSendTemplateMessageParam struct {
	Touser      string                 `json:"touser"`          //接收用户openid
	TemplateID  string                 `json:"template_id"`     //消息模版ID
	Url         string                 `json:"url,omitempty"`   //点击后的跳转页面，仅限本小程序内的页面
	Data        map[string]interface{} `json:"data"`            //模板内容
	Color       string                 `json:"color,omitempty"` //模板内容字体颜色，不填默认为黑色
	Miniprogram WxMiniprogram          `json:"miniprogram,omitempty"`
	ClientMsgId string                 `json:"client_msg_id,omitempty"` //防重入id。对于同一个openid + client_msg_id, 只发送一条消息,10分钟有效,超过10分钟不保证效果。若无防重入需求，可不填
}

// GetTemplateList 获取私有模版
func (sdk *SDK) GetMessageTemplateList(ctx context.Context, appid string) (*WxGetTemplateRes, error) {
	req := &WxGetTemplateRes{}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%s", sdk.AccessToken)
	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

// SendTemplateMessage 发送模版信息
func (sdk *SDK) SendTemplateMessage(ctx context.Context, param *WxSendTemplateMessageParam) (string, error) {

	if param.TemplateID == "" {
		return "", fmt.Errorf("template_id is empty")
	}
	if param.Touser == "" {
		return "", fmt.Errorf("touser is empty")
	}

	bodyMap := make(common.BodyMap)
	bodyMap.Set("touser", param.Touser)
	bodyMap.Set("template_id", param.TemplateID)
	bodyMap.Set("data", param.Data)

	if param.Color != "" {
		bodyMap.Set("color", param.Color)
	}

	if param.Miniprogram.Appid != "" {
		bodyMap.Set("miniprogram", param.Miniprogram) //公众号
	}

	if param.Color != "" {
		bodyMap.Set("url", param.Url)
	}

	if param.ClientMsgId != "" {
		bodyMap.Set("client_msg_id", param.ClientMsgId)
	}

	req := &WxSendTemplateMessageRes{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return "", fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return strconv.FormatInt(req.Msgid, 10), nil
}
