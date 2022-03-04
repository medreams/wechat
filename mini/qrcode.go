package mini

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/medreams/wechat/common"
)

type LineColorParam struct {
	R string `form:"r" json:"r" default:"0"`
	G string `form:"g" json:"g" default:"0"`
	B string `form:"b" json:"b" default:"0"`
}

//WxMiniSceneParam 小程复杂二维码入参
type WxMiniSceneParam struct {
	Scene     string `form:"scene" json:"scene" binding:"required,max=32"`                     //最大32个可见字符，只支持数字，大小写英文以及部分特殊字符
	Page      string `form:"page" json:"page" binding:"required,max=64"`                       //必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	LineColor string `form:"line_color" json:"line_color" default:"{'r':'0','g':'0','b':'0'}"` //二维码线条颜色
	Width     int    `form:"width" json:"width" default:"430"`                                 //二维码的宽度，单位 px，最小 280px，最大 1280px
	IsHyaline bool   `form:"is_hyaline" json:"is_hyaline" default:"true"`                      //是否需要透明底
}

//WxMiniPathParam 小程简单二维码入参
type WxMiniPathParam struct {
	Path  string `form:"path" json:"path" binding:"required,max=128"` //不能为空，最大长度 128 字节
	Width int    `form:"width" json:"width" default:"430"`            //二维码的宽度，单位 px，最小 280px，最大 1280px
}

//ReturnMiniCodeData 返回二维码图片数据
type ReturnMiniCodeData struct {
	ImageBase64 string `form:"image_base64" json:"image_base64"` //图片的base64格式
}

//CreateMiniSceneCode 创建小程序二维码（较多业务场景）
func (sdk *SDK) CreateMiniSceneCode(ctx context.Context, param *WxMiniSceneParam) (ret []byte, err error) {
	//log.Info.Println("创建小程序二维码（较多业务场景）", param)

	if param.LineColor == "" {
		param.LineColor = `{"r":"0","g":"0","b":"0"}`
	}

	var lineColorMap LineColorParam
	json.Unmarshal([]byte(param.LineColor), &lineColorMap)

	if param.Width == 0 {
		param.Width = 430
	}

	bodyMap := make(map[string]interface{})
	bodyMap["scene"] = param.Scene //最大32个可见字符，只支持数字
	bodyMap["width"] = param.Width
	bodyMap["page"] = param.Page            //必须是已经发布的小程序存在的页面
	bodyMap["auto_color"] = false           //自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	bodyMap["is_hyaline"] = param.IsHyaline //是否需要透明底色
	bodyMap["line_color"] = lineColorMap

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", sdk.AccessToken)

	bs, err := common.DoRequestPostGetByte(ctx, uri, bodyMap)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

//CreateMiniDefaultCode 创建小程序二维码（较少业务场景）
func (sdk *SDK) CreateMiniDefaultCode(ctx context.Context, param *WxMiniPathParam) (ret []byte, err error) {

	bodyMap := make(map[string]interface{})
	bodyMap["path"] = param.Path
	bodyMap["width"] = param.Width

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s", sdk.AccessToken)

	retByte, err := common.DoRequestPostGetByte(ctx, uri, bodyMap)
	if err != nil {
		return nil, err
	}

	return retByte, nil
}
