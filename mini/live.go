package mini

import (
	"context"
	"fmt"
	"net/url"

	"github.com/medreams/wechat/common"
	"github.com/medreams/wechat/pkg/util"
)

type CreateRoomParam struct {
	Name            string `json:"name"`            //直播间名字，最短3个汉字，最长17个汉字，1个汉字相当于2个字符
	CoverImg        string `json:"coverImg"`        //背景图，填入mediaID（mediaID获取后，三天内有效）；图片 mediaID 的获取，请参考以下文档： https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html；直播间背景图，图片规则：建议像素1080*1920，大小不超过2M
	StartTime       int    `json:"startTime"`       //直播计划开始时间（开播时间需要在当前时间的10分钟后 并且 开始时间不能在 6 个月后）
	EndTime         int    `json:"endTime"`         //直播计划结束时间（开播时间和结束时间间隔不得短于30分钟，不得超过24小时）
	AnchorName      string `json:"anchorName"`      //主播昵称，最短2个汉字，最长15个汉字，1个汉字相当于2个字符
	AnchorWechat    string `json:"anchorWechat"`    //主播微信号，如果未实名认证，需要先前往“小程序直播”小程序进行实名验证, 小程序二维码链接：https://res.wx.qq.com/op_res/9rSix1dhHfK4rR049JL0PHJ7TpOvkuZ3mE0z7Ou_Etvjf-w1J_jVX0rZqeStLfwh
	SubAnchorWechat string `json:"subAnchorWechat"` //主播副号微信号，如果未实名认证，需要先前往“小程序直播”小程序进行实名验证, 小程序二维码链接：https://res.wx.qq.com/op_res/9rSix1dhHfK4rR049JL0PHJ7TpOvkuZ3mE0z7Ou_Etvjf-w1J_jVX0rZqeStLfwh
	CreaterWechat   string `json:"createrWechat"`   //创建者微信号，不传入则此直播间所有成员可见。传入则此房间仅创建者、管理员、超管、直播间主播可见
	ShareImg        string `json:"shareImg"`        //分享图，填入mediaID（mediaID获取后，三天内有效）；图片 mediaID 的获取，请参考以下文档： https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html；直播间分享图，图片规则：建议像素800*640，大小不超过1M；
	FeedsImg        string `json:"feedsImg"`        //购物直播频道封面图，填入mediaID（mediaID获取后，三天内有效）；图片 mediaID 的获取，请参考以下文档： https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html; 购物直播频道封面图，图片规则：建议像素800*800，大小不超过100KB；
	IsFeedsPublic   int    `json:"isFeedsPublic"`   //是否开启官方收录 【1: 开启，0：关闭】，默认开启收录
	Type            int    `json:"type"`            //直播间类型 【1: 推流，0：手机直播】
	CloseLike       int    `json:"closeLike"`       //是否关闭点赞 【0：开启，1：关闭】（若关闭，观众端将隐藏点赞按钮，直播开始后不允许开启）
	CloseGoods      int    `json:"closeGoods"`      //是否关闭货架 【0：开启，1：关闭】（若关闭，观众端将隐藏商品货架，直播开始后不允许开启）
	CloseComment    int    `json:"closeComment"`    //是否关闭评论 【0：开启，1：关闭】（若关闭，观众端将隐藏评论入口，直播开始后不允许开启）
	CloseReplay     int    `json:"closeReplay"`     //是否关闭回放 【0：开启，1：关闭】默认关闭回放（直播开始后允许开启）
	CloseShare      int    `json:"closeShare"`      //是否关闭分享 【0：开启，1：关闭】默认开启分享（直播开始后不允许修改）
	CloseKf         int    `json:"closeKf"`         //是否关闭客服 【0：开启，1：关闭】 默认关闭客服（直播开始后允许开启）

}

type CreateRoomRsp struct {
	common.WxCommonResponse
	RoomId    string `json:"roomId"`
	QrcodeUrl string `json:"qrcode_url"`
}

// 创建直播间 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/createRoom.html
func (sdk *SDK) CreateRoom(ctx context.Context, param *CreateRoomParam) (*CreateRoomRsp, error) {
	bodyMap := util.ConvertToMap(param)

	req := &CreateRoomRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/create?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

// 删除直播间 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/deleteRoom.html
func (sdk *SDK) DeleteRoom(ctx context.Context, id int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("id", id)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/deleteroom?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

type EditRoomParam struct {
	ID            int    `json:"id"`            //直播间id
	Name          string `json:"name"`          //直播间名字，最短3个汉字，最长17个汉字，1个汉字相当于2个字符
	CoverImg      string `json:"coverImg"`      //直播间背景图，图片规则：建议像素1080*1920，大小不超过2M，填入mediaID（mediaID获取后，三天内有效）；图片 mediaID 的获取
	StartTime     int    `json:"startTime"`     //直播计划开始时间（开播时间需要在当前时间的10分钟后 并且 开始时间不能在 6 个月后）
	EndTime       int    `json:"endTime"`       //直播计划结束时间（开播时间和结束时间间隔不得短于30分钟，不得超过24小时）
	AnchorName    string `json:"anchorName"`    //主播昵称，最短2个汉字，最长15个汉字，1个汉字相当于2个字符
	AnchorWechat  string `json:"anchorWechat"`  //主播微信号
	ShareImg      string `json:"shareImg"`      //直播间分享图，图片规则：建议像素800*640，大小不超过1M
	CloseLike     int    `json:"closeLike"`     //购物直播频道封面图，图片规则：建议像素800*800，大小不超过100KB；
	CloseGoods    int    `json:"closeGoods"`    //是否开启官方收录 【1: 开启，0：关闭】，默认开启收录
	CloseComment  int    `json:"closeComment"`  //是否关闭点赞 【0：开启，1：关闭】（若关闭，观众端不展示点赞入口，直播开始后不允许开启）
	IsFeedsPublic int    `json:"isFeedsPublic"` //是否关闭货架 【0：开启，1：关闭】（若关闭，观众端不展示商品货架，直播开始后不允许开启）
	CloseReplay   int    `json:"closeReplay"`   //是否关闭评论 【0：开启，1：关闭】（若关闭，观众端不展示评论入口，直播开始后不允许开启）
	CloseShare    int    `json:"closeShare"`    //是否关闭回放 【0：开启，1：关闭】默认关闭回放（直播开始后允许开启）
	CloseKf       int    `json:"closeKf"`       //是否关闭分享 【0：开启，1：关闭】默认开启分享（直播开始后不允许修改）
	FeedsImg      string `json:"feedsImg"`      //是否关闭客服 【0：开启，1：关闭】 默认关闭客服（直播开始后允许开启）

}

// 编辑直播间 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/editRoom.html
func (sdk *SDK) EditRoom(ctx context.Context, param *EditRoomParam) error {
	bodyMap := util.ConvertToMap(param)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/editroom?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

type ImportLiveGoodsParam struct {
	Ids    []int `json:"ids"`     //数组列表，可传入多个，里面填写 商品 ID
	RoomId int   `json:"room_id"` //房间ID
}

// 导入商品 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/importGoods.html
func (sdk *SDK) ImportLiveGoods(ctx context.Context, param *ImportLiveGoodsParam) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("ids", param.Ids)
	bodyMap.Set("roomId", param.RoomId)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/addgoods?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 推送商品 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/pushGoods.html
func (sdk *SDK) PushLiveGoods(ctx context.Context, roomId int, goodsId int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)
	bodyMap.Set("goodsId", goodsId)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/goods/push?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 上下架商品 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/SaleGoods.html
func (sdk *SDK) SaleLiveGoods(ctx context.Context, roomId int, goodsId int, onSale int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)
	bodyMap.Set("goodsId", goodsId)
	bodyMap.Set("onSale", onSale)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/goods/onsale?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 直播间商品排序 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/sortGoods.html
func (sdk *SDK) SortLiveGoods(ctx context.Context, roomId int, goodsIds []int) error {

	var goods []map[string]int
	for _, id := range goodsIds {
		goods = append(goods, map[string]int{
			"goodsId": id,
		})
	}

	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)
	bodyMap.Set("goodsId", goods)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/goods/sort?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 删除直播间商品 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/deleteDoods.html
func (sdk *SDK) DelLiveGoods(ctx context.Context, roomId int, goodsId int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)
	bodyMap.Set("goodsId", goodsId)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/goods/deleteInRoom?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

type GetRoomeListParam struct {
	Start  int    `json:"start"`    //起始拉取视频，0表示从第一个视频片段开始拉取
	Limit  int    `json:"limit"`    //每次拉取的数量，建议100以内
	Action string `json:"action"`   //只能填"get_replay"，表示获取回放。
	RoomId int    `json:"roome_id"` //当 action 有值时该字段必填，直播间ID

}

type RoomeListRsp struct {
	common.WxCommonResponse
	Total    int        `json:"total"`     //拉取房间总数
	RoomInfo []RoomInfo `json:"room_info"` //action="get_replay"不返回。
}
type LiveGoods struct {
	CoverImg        string `json:"cover_img"`         //商品封面图链接
	Url             string `json:"url"`               //商品小程序路径
	Name            string `json:"name"`              //商品名称
	Price           int    `json:"price"`             //商品价格（分）
	Price2          int    `json:"price2"`            //商品价格，使用方式看price_type
	PriceType       int    `json:"price_type"`        //价格类型，1：一口价（只需要传入price，price2不传） 2：价格区间（price字段为左边界，price2字段为右边界，price和price2必传） 3：显示折扣价（price字段为原价，price2字段为现价， price和price2必传）
	GoodsID         int    `json:"goods_id"`          //商品id
	ThirdPartyAppid string `json:"third_party_appid"` //第三方商品appid ,当前小程序商品则为空
}
type LiveReplay struct {
	CreateTime string `json:"create_time"` //回放视频创建时间
	ExpireTime string `json:"expire_time"` //回放视频 url 过期时间
	MediaUrl   string `json:"media_url"`   //回放视频链接
}
type RoomInfo struct {
	Name          string      `json:"name"`            //直播间名称
	Roomid        int         `json:"roomid"`          //直播间ID
	CoverImg      string      `json:"cover_img"`       //直播间背景图链接
	ShareImg      string      `json:"share_img"`       //直播间分享图链接
	LiveStatus    int         `json:"live_status"`     //直播间状态。101：直播中，102：未开始，103已结束，104禁播，105：暂停，106：异常，107：已过期
	StartTime     int         `json:"start_time"`      //直播间开始时间，列表按照start_time降序排列
	EndTime       int         `json:"end_time"`        //直播计划结束时间
	AnchorName    string      `json:"anchor_name"`     //主播名
	Goods         []LiveGoods `json:"goods"`           //商品
	LiveType      int         `json:"live_type"`       //直播类型，1 推流 0 手机直播
	CloseLike     int         `json:"close_like"`      //是否关闭点赞 【0：开启，1：关闭】（若关闭，观众端将隐藏点赞按钮，直播开始后不允许开启）
	CloseGoods    int         `json:"close_goods"`     //是否关闭货架 【0：开启，1：关闭】（若关闭，观众端将隐藏商品货架，直播开始后不允许开启）
	CloseComment  int         `json:"close_comment"`   //是否关闭评论 【0：开启，1：关闭】（若关闭，观众端将隐藏评论入口，直播开始后不允许开启）
	CloseKf       int         `json:"close_kf"`        //是否关闭客服 【0：开启，1：关闭】 默认关闭客服（直播开始后允许开启）
	CloseReplay   int         `json:"close_replay"`    //是否关闭回放 【0：开启，1：关闭】默认关闭回放（直播开始后允许开启）
	IsFeedsPublic int         `json:"is_feeds_public"` //是否开启官方收录，1 开启，0 关闭
	CreaterOpenid string      `json:"creater_openid"`  //创建者openid
	FeedsImg      string      `json:"feeds_img"`       //官方收录封面
}

// 获取直播间列表和回放 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/getLiveInfo.html
func (sdk *SDK) GetLiveInfo(ctx context.Context, param *GetRoomeListParam) (*RoomeListRsp, error) {
	bodyMap := util.ConvertToMap(param)

	req := &RoomeListRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getliveinfo?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

// 获取直播间推流地址 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/getPushUrl.html
func (sdk *SDK) GetPushUrl(ctx context.Context, roomId int) (string, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)

	var req struct {
		ErrCode  int64  `json:"errcode"` //错误码
		PushAddr string `json:"pushAddr"`
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/getpushurl?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, &req); err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return "", fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req.PushAddr, nil
}

type LiveSharedCodeRsp struct {
	common.WxCommonResponse
	CdnUrl    string `json:"cdnUrl"`    //分享二维码cdn url
	PagePath  string `json:"pagePath"`  //分享路径
	PosterUrl string `json:"posterUrl"` //分享海报 url
}

// 获取直播间分享二维码 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/getSharedCode.html
func (sdk *SDK) GetLiveSharedCode(ctx context.Context, roomId string, params string) (*LiveSharedCodeRsp, error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)                  //房间ID
	bodyMap.Set("params", url.QueryEscape(params)) //自定义参数

	req := &LiveSharedCodeRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/getpushurl?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

// 添架主播副号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/addSubAnchor.html
func (sdk *SDK) CreateLiveSubAnchor(ctx context.Context, roomId int, username string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)
	bodyMap.Set("username", username)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/addsubanchor?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 修改主播副号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/modifySubAnchor.html
func (sdk *SDK) EditLiveSubAnchor(ctx context.Context, roomId int, username string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)
	bodyMap.Set("username", username)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/modifysubanchor?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 获取主播副号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/getSubAnchor.html
func (sdk *SDK) GetLiveSubAnchor(ctx context.Context, roomId int) (string, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)

	var req struct {
		ErrCode  int64  `json:"errcode"` //错误码
		Username string `json:"username"`
	}

	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/getsubanchor?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, &req); err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return "", fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req.Username, nil
}

// 删除主播副号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/deleteSubAnchor.html
func (sdk *SDK) DelLiveSubAnchor(ctx context.Context, roomId int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/deletesubanchor?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

type LiveRoomeAssistant struct {
	Username string `json:"username"` //用户微信号
	Nickname string `json:"nickname"` //用户昵称
}

// 添加管理直播间小助手 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/addveAssistant.html
func (sdk *SDK) CreateLiveAssistant(ctx context.Context, roomId int, users []LiveRoomeAssistant) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)
	bodyMap.Set("users", users)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/addassistant?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 修改直播间小助手 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/modifyAssistant.html
func (sdk *SDK) EditLiveAssistant(ctx context.Context, roomId int, user LiveRoomeAssistant) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)
	bodyMap.Set("username", user.Username)
	bodyMap.Set("nickname", user.Nickname)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/modifyassistant?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 删除直播间小助手 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/removeAssistant.html
func (sdk *SDK) DelLiveAssistant(ctx context.Context, roomId int, username string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)
	bodyMap.Set("username", username)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/removeassistant?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

type LiveRoomeAssistantListRsp struct {
	List     []LiveRoomeAssistantList `json:"list"`     //小助手列表
	Count    int                      `json:"count"`    //小助手个数
	MaxCount int                      `json:"maxCount"` //小助手最大个数
	ErrCode  int                      `json:"errcode"`  //返回码
}

type LiveRoomeAssistantList struct {
	Timestamp int    `json:"timestamp"` //修改时间
	Headimg   string `json:"headimg"`   //头像
	Nickname  string `json:"nickname"`  //昵称
	Alias     string `json:"alias"`     //微信号
	Openid    string `json:"openid"`    //openid
}

// 查询直播间小助手 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/getAssistantList.html
func (sdk *SDK) GetLiveAssistant(ctx context.Context, roomId int) (*LiveRoomeAssistantListRsp, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("roomId", roomId)

	req := &LiveRoomeAssistantListRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/getassistantlist?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

// 禁言管理(整个直播间) https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/updateComment.html
func (sdk *SDK) UpdateLiveRomeComment(ctx context.Context, roomId int, banComment int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("id", roomId)             //房间ID
	bodyMap.Set("banComment", banComment) //1-禁言，0-取消禁言

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/updatecomment?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 官方收录管理 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/updateFeedPublic.html
func (sdk *SDK) UpdateLiveRoomeFeedPublic(ctx context.Context, roomId int, isFeedsPublic int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("id", roomId)
	bodyMap.Set("isFeedsPublic", isFeedsPublic) //是否开启官方收录 【1: 开启，0：关闭】

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/updatefeedpublic?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 客服功能管理 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/updateKF.html
func (sdk *SDK) UpdateLiveRoomeKf(ctx context.Context, roomId int, closeKf int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("id", roomId)
	bodyMap.Set("closeKf", closeKf) //是否关闭客服 【0：开启，1：关闭】

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/updatekf?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 回放功能管理 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/updateReplay.html
func (sdk *SDK) UpdateLiveRoomeReplay(ctx context.Context, roomId int, closeReplay int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("id", roomId)
	bodyMap.Set("closeReplay", closeReplay) //是否关闭回放 【0：开启，1：关闭】

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/room/updatereplay?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 下载商品讲解视频 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/livebroadcast/studio-management/downloadGoodsVideo.html
func (sdk *SDK) DownloadGoodsVideo(ctx context.Context, roomId int, goodsId int) (string, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("id", roomId)
	bodyMap.Set("goodsId", goodsId) //商品ID

	var req struct {
		ErrCode int64  `json:"errcode"` //错误码
		Url     string `json:"url"`
	}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxaapi/broadcast/goods/getVideo?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return "", fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return "", fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req.Url, nil
}
