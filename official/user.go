package official

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

// 微信公众号用户信息
type PublicUserInfo struct {
	Subscribe      int    `json:"subscribe,omitempty"`       // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	Openid         string `json:"openid,omitempty"`          // 用户唯一标识
	Nickname       string `json:"nickname,omitempty"`        // 用户的昵称
	Sex            int    `json:"sex,omitempty"`             // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City           string `json:"city,omitempty"`            // 用户所在城市
	Province       string `json:"province,omitempty"`        // 用户所在省份
	Country        string `json:"country,omitempty"`         // 用户所在国家
	Language       string `json:"language,omitempty"`        // 用户的语言，简体中文为zh_CN
	Headimgurl     string `json:"headimgurl,omitempty"`      // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	SubscribeTime  int    `json:"subscribe_time,omitempty"`  // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	Unionid        string `json:"unionid,omitempty"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark         string `json:"remark,omitempty"`          // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Groupid        int    `json:"groupid,omitempty"`         // 用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      []int  `json:"tagid_list,omitempty"`      // 用户被打上的标签ID列表
	SubscribeScene string `json:"subscribe_scene,omitempty"` // 返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENEPROFILE LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_OTHERS 其他
	QrScene        int    `json:"qr_scene,omitempty"`        // 二维码扫码场景（开发者自定义）
	QrSceneStr     string `json:"qr_scene_str,omitempty"`    // 二维码扫码场景描述（开发者自定义）
	Errcode        int    `json:"errcode,omitempty"`         // 错误码
	Errmsg         string `json:"errmsg,omitempty"`          // 错误信息
}

//获取用户基本信息(UnionID机制)
//文档地址 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func (s *SDK) Openid2UserInfo(ctx context.Context, openid string) (user *PublicUserInfo, err error) {

	url := "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + s.AccessToken + "&openid=" + openid + "&lang=zh_CN"

	user = &PublicUserInfo{}
	if err = common.DoRequestGet(ctx, url, user); err != nil {
		return nil, fmt.Errorf("do request get userinfo: %w", err)
	}

	return user, nil
}
