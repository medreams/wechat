package official

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
	"github.com/medreams/wechat/pkg/util"
)

// 客服管理  https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Manage

func toError(code int64, msg string) error {

	switch code {
	case 0:
		return fmt.Errorf("成功")
	case 65400:
		return fmt.Errorf("65400 API不可用，即没有开通/升级到新版客服功能")
	case 65401:
		return fmt.Errorf("65401 无效客服帐号")
	case 65403:
		return fmt.Errorf("65403 客服昵称不合法")
	case 65404:
		return fmt.Errorf("65404 客服帐号不合法")
	case 65405:
		return fmt.Errorf("65405 帐号数目已达到上限，不能继续添加")
	case 65406:
		return fmt.Errorf("65406 已经存在的客服帐号")
	case 65407:
		return fmt.Errorf("65407 邀请对象已经是该公众号客服")
	case 65408:
		return fmt.Errorf("65408 本公众号已经有一个邀请给该微信")
	case 65409:
		return fmt.Errorf("65409 无效的微信号")
	case 65410:
		return fmt.Errorf("65410 邀请对象绑定公众号客服数达到上限（目前每个微信号可以绑定5个公众号客服帐号）")
	case 65411:
		return fmt.Errorf("65411 该帐号已经有一个等待确认的邀请，不能重复邀请")
	case 65412:
		return fmt.Errorf("65412 该帐号已经绑定微信号，不能进行邀请")
	case 65413:
		return fmt.Errorf("65413 不存在对应用户的会话信息")
	case 65414:
		return fmt.Errorf("65414 粉丝正在被其他客服接待")
	case 65415:
		return fmt.Errorf("65415 指定的客服不在线")
	case 65416:
		return fmt.Errorf("65416 查询参数不合法")
	case 65417:
		return fmt.Errorf("65417 查询时间段超出限制")
	case 40005:
		return fmt.Errorf("40005 不支持的媒体类型")
	case 40009:
		return fmt.Errorf("40009 媒体文件长度不合法")
	default:
		return fmt.Errorf("%d %s", code, msg)
	}
}

type CustomerInfo struct {
	KfAccount        string `json:"kf_account"`         //完整客服帐号，格式为：帐号前缀@公众号微信号
	KfHeadimgurl     string `json:"kf_headimgurl"`      //客服头像
	KfID             string `json:"kf_id"`              //客服编号
	KfNick           string `json:"kf_nick"`            //客服昵称
	KfWx             string `json:"kf_wx"`              //如果客服帐号已绑定了客服人员微信号， 则此处显示微信号
	InviteWx         string `json:"invite_wx"`          //如果客服帐号尚未绑定微信号，但是已经发起了一个绑定邀请， 则此处显示绑定邀请的微信号
	InviteExpireTime int64  `json:"invite_expire_time"` //如果客服帐号尚未绑定微信号，但是已经发起过一个绑定邀请， 邀请的过期时间，为unix 时间戳
	InviteStatus     string `json:"invite_status"`      //邀请的状态，有等待确认“waiting”，被拒绝“rejected”， 过期“expired”

}

type CustomerList struct {
	common.WxCommonResponse
	KfList []CustomerInfo `json:"kf_list"`
}

// 客服列表
func (sdk *SDK) GetCustomerList(ctx context.Context) (*CustomerList, error) {

	req := &CustomerList{}

	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s", sdk.AccessToken)
	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, toError(req.ErrCode, req.ErrMsg)
	}

	return req, nil
}

type CustomerOnlineInfo struct {
	KfAccount    string `json:"kf_account"`    //完整客服帐号，格式为：帐号前缀@公众号微信号
	Status       int    `json:"status"`        //客服在线状态，目前为：1、web 在线
	KfID         int    `json:"kf_id"`         //客服编号
	AcceptedCase int    `json:"accepted_case"` //客服当前正在接待的会话数
}

type CustomerOnlineList struct {
	common.WxCommonResponse
	KfOnlineList []*CustomerOnlineInfo `json:"kf_online_list"`
}

// 在线客服列表
func (sdk *SDK) GetOnlineCustomerList(ctx context.Context) (*CustomerOnlineList, error) {

	req := &CustomerOnlineList{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=%s", sdk.AccessToken)
	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, toError(req.ErrCode, req.ErrMsg)
	}

	return req, nil
}

// 添加客服帐号
func (sdk *SDK) CreateCustomer(ctx context.Context, account, nickname string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("kf_account", account)
	bodyMap.Set("nickname", nickname)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return toError(req.ErrCode, req.ErrMsg)
	}
	return nil
}

// 邀请绑定客服帐号
func (sdk *SDK) InviteCustomer(ctx context.Context, account, inviteWx string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("kf_account", account)
	bodyMap.Set("invite_wx", inviteWx)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return toError(req.ErrCode, req.ErrMsg)
	}
	return nil
}

// 设置客服昵称
func (sdk *SDK) UpdateCustomerNickname(ctx context.Context, account, nickname string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("kf_account", account)
	bodyMap.Set("nickname", nickname)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return toError(req.ErrCode, req.ErrMsg)
	}
	return nil
}

// 上传客服头像
func (sdk *SDK) UploadCustomerHeadimg(ctx context.Context, account string, headimg *util.File) error {
	bodyMap := make(common.BodyMap)
	bodyMap.SetFormFile("media", headimg)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%s&kf_account=%s", sdk.AccessToken, account)

	if err := common.DoUploadFile(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return toError(req.ErrCode, req.ErrMsg)
	}
	return nil
}

// 删除客服帐号
func (sdk *SDK) DeleteCustomer(ctx context.Context, account string) error {

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s&kf_account=%s", sdk.AccessToken, account)
	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return toError(req.ErrCode, req.ErrMsg)
	}
	return nil
}

// 创建会话
func (sdk *SDK) CreateCustomerSession(ctx context.Context, account, openid string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("kf_account", account)
	bodyMap.Set("openid", openid)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfsession/create?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return toError(req.ErrCode, req.ErrMsg)
	}
	return nil
}

// 关闭会话
func (sdk *SDK) CloseCustomerSession(ctx context.Context, account, openid string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("kf_account", account)
	bodyMap.Set("openid", openid)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfsession/close?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return toError(req.ErrCode, req.ErrMsg)
	}
	return nil
}

type CustomerSessionStatus struct {
	common.WxCommonResponse
	Createtime int    `json:"createtime"` //正在接待的客服，为空表示没有人在接待
	KfAccount  string `json:"kf_account"` //会话接入的时间
}

// 获取客户会话状态
func (sdk *SDK) GetCustomerSession(ctx context.Context, openid string) (*CustomerSessionStatus, error) {

	req := &CustomerSessionStatus{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, toError(req.ErrCode, req.ErrMsg)
	}
	return req, nil
}

type CustomerSessionList struct {
	common.WxCommonResponse
	Sessionlist []CustomerSession `json:"sessionlist"`
}

type CustomerSession struct {
	Createtime int    `json:"createtime"`
	Openid     string `json:"openid"`
}

// 获取客服会话列表
func (sdk *SDK) GetCustomerSessionList(ctx context.Context, account string) (*CustomerSessionList, error) {

	req := &CustomerSessionList{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, toError(req.ErrCode, req.ErrMsg)
	}
	return req, nil
}

type CustomerWaitSessionList struct {
	common.WxCommonResponse
	Count        int                   `json:"count"`        //未接入会话数量
	Waitcaselist []CustomerWaitSession `json:"waitcaselist"` //未接入会话列表，最多返回100条数据，按照来访顺序
}

type CustomerWaitSession struct {
	LatestTime int    `json:"latest_time"` //粉丝的最后一条消息的时间
	Openid     string `json:"openid"`      //粉丝的openid
}

// 获取未接入会话列表
func (sdk *SDK) GetCustomerWaitSessionList(ctx context.Context, account string) (*CustomerWaitSessionList, error) {

	req := &CustomerWaitSessionList{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, toError(req.ErrCode, req.ErrMsg)
	}
	return req, nil
}

type CustomerMsgList struct {
	Recordlist []CustomerMsg `json:"recordlist"`
	Number     int           `json:"number"`
	Msgid      int           `json:"msgid"`
}

type CustomerMsg struct {
	Openid   string `json:"openid"`   //用户标识
	Opercode int    `json:"opercode"` //操作码，2002（客服发送信息），2003（客服接收消息）
	Text     string `json:"text"`     //聊天记录
	Time     int    `json:"time"`     //操作时间，unix时间戳
	Worker   string `json:"worker"`   //完整客服帐号，格式为：帐号前缀@公众号微信号
}

// 获取聊天记录
func (sdk *SDK) GetCustomerMsg(ctx context.Context, starttime, endtime, msgid, number int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("starttime", starttime)
	bodyMap.Set("endtime", endtime)
	bodyMap.Set("msgid", msgid)
	bodyMap.Set("number", number)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/customservice/msgrecord/getmsglist?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return toError(req.ErrCode, req.ErrMsg)
	}
	return nil
}
