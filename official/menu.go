package official

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
	"github.com/medreams/wechat/pkg/util"
)

type CreateMenuParams struct {
	Button []MenuButton `json:"button"` //一级菜单数组，个数应为1~3个
}

type SubMenuButton struct {
	Type     string `json:"type"`     //菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型
	Name     string `json:"name"`     //菜单标题，不超过16个字节，子菜单不超过60个字节
	Url      string `json:"url"`      //网页 链接，用户点击菜单可打开链接，不超过1024字节。 type为 miniprogram 时，不支持小程序的老版本客户端将打开本url。
	Appid    string `json:"appid"`    //小程序的appid（仅认证公众号可配置）
	Pagepath string `json:"pagepath"` //小程序的页面路径
	Key      string `json:"key"`      //菜单 KEY 值，用于消息接口推送，不超过128字节
}

type MenuButton struct {
	Type      string          `json:"type"`       //菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型
	Name      string          `json:"name"`       //菜单标题，不超过16个字节，子菜单不超过60个字节
	Key       string          `json:"key"`        //菜单 KEY 值，用于消息接口推送，不超过128字节
	SubButton []SubMenuButton `json:"sub_button"` //二级菜单数组，个数应为1~5个
	MediaId   string          `json:"media_id"`   //调用新增永久素材接口返回的合法media_id
	ArticleId string          `json:"article_id"` //发布后获得的合法 article_id
}

// 创建自定义菜单（单个） https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html
func (sdk *SDK) CreateCustomMenu(ctx context.Context, param *CreateMenuParams) error {
	bodyMap := util.ConvertToMap(param)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

type GetMenuRsp struct {
	IsMenuOpen   int         `json:"is_menu_open"`  //菜单是否开启，0代表未开启，1代表开启
	SelfMenuInfo GetMenuInfo `json:"selfmenu_info"` //菜单信息
}

type GetMenuSubButton struct {
	List []SubMenuButton `json:"list"`
}

type GetMenuButton struct {
	Type      string           `json:"type"` //菜单的类型，公众平台官网上能够设置的菜单类型有view（跳转网页）、text（返回文本，下同）、img、photo、video、voice。使用 API 设置的则有8种，详见《自定义菜单创建接口》
	Name      string           `json:"name"` //菜单名称
	Key       string           `json:"key"`
	SubButton GetMenuSubButton `json:"sub_button"`
}

type GetMenuInfo struct {
	Button []GetMenuButton `json:"button"` //菜单按钮
}

// 查询自定义菜单
func (sdk *SDK) QueryCustomMenu(ctx context.Context) (*GetMenuRsp, error) {
	req := &GetMenuRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, nil, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	return req, nil
}

// 删除自定义菜单（调用此接口会删除默认菜单及全部个性化菜单）
func (sdk *SDK) DelCustomMenu(ctx context.Context) error {
	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, nil, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}
