package official

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

// 文章
type Article struct {
	Title              string `json:"title"`                           //标题
	Author             string `json:"author,omitempty"`                //作者
	Digest             string `json:"digest,omitempty"`                //图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。
	Content            string `json:"content"`                         //图文消息的具体内容，支持 HTML 标签，必须少于2万字符，小于1M，且此处会去除 JS ,涉及图片 url 必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片 url 将被过滤。
	ContentSourceURL   string `json:"content_source_url,omitempty"`    //图文消息的原文地址，即点击“阅读原文”后的URL
	ThumbMediaID       string `json:"thumb_media_id"`                  //图文消息的封面图片素材id（必须是永久MediaID）
	NeedOpenComment    int    `json:"need_open_comment,omitempty"`     //Uint32 是否打开评论，0不打开(默认)，1打开
	OnlyFansCanComment int    `json:"only_fans_can_comment,omitempty"` //Uint32 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论
	ShowCoverPic       string `json:"show_cover_pic,omitempty"`        //是否在正文显示封面。平台已不支持此功能，因此默认为0，即不展示,只有查询时返回，新建时不用提交
	Url                string `json:"url,omitempty"`                   //草稿的临时链接,只有查询时返回，新建时不用提交
}

type AddDraftRsp struct {
	common.WxCommonResponse
	MediaId string `json:"media_id,omitempty"`
}

// 新建草稿 https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html
func (sdk *SDK) CreateDraft(ctx context.Context, param []*Article) (*AddDraftRsp, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("articles", param)

	req := &AddDraftRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/draft/add?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, common.ToError(req.ErrCode, req.ErrMsg)
	}

	return req, nil
}

type GetDraftRsp struct {
	common.WxCommonResponse
	NewsItem []Article `json:"news_item,omitempty"`
}

// 获取草稿 https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Get_draft.html
func (sdk *SDK) GetDraft(ctx context.Context, mediaId string) (*GetDraftRsp, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("media_id", mediaId)

	req := &GetDraftRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/draft/get?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, common.ToError(req.ErrCode, req.ErrMsg)
	}

	return req, nil
}

// 删除草稿 https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Delete_draft.html
func (sdk *SDK) DelDraft(ctx context.Context, mediaId string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("media_id", mediaId)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/draft/delete?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return common.ToError(req.ErrCode, req.ErrMsg)
	}

	return nil
}

// 修改草稿 https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Update_draft.html
func (sdk *SDK) UpdateDraft(ctx context.Context, mediaId string, index int, article Article) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("media_id", mediaId)
	bodyMap.Set("index", index)
	bodyMap.Set("articles", article)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/draft/update?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return common.ToError(req.ErrCode, req.ErrMsg)
	}

	return nil
}

type GetDraftTotalRsp struct {
	common.WxCommonResponse
	TotalCount int `json:"total_count"`
}

// 获取草稿总数 https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Count_drafts.html
func (sdk *SDK) GetDraftTotal(ctx context.Context) (*GetDraftTotalRsp, error) {
	req := &GetDraftTotalRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/draft/count?access_token=%s", sdk.AccessToken)
	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, common.ToError(req.ErrCode, req.ErrMsg)
	}
	return req, nil
}

type GetDraftListRsp struct {
	common.WxCommonResponse
	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
	Item       []struct {
		MediaId    string `json:"media_id"`
		UpdateTime string `json:"update_time"`
		Content    struct {
			NewsItem []Article `json:"news_item"`
		} `json:"content"`
	} `json:"item"`
}

// 获取草稿列表 https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Get_draft_list.html
func (sdk *SDK) GetDraftList(ctx context.Context, offset, count, noContent int) (*GetDraftRsp, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("offset", offset)        //从全部素材的该偏移位置开始返回，0表示从第一个素材返回
	bodyMap.Set("count", count)          //返回素材的数量，取值在1到20之间
	bodyMap.Set("no_content", noContent) //1 表示不返回 content 字段，0 表示正常返回，默认为 0

	req := &GetDraftRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/draft/batchget?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, common.ToError(req.ErrCode, req.ErrMsg)
	}

	return req, nil
}

type DraftSwitchRsp struct {
	common.WxCommonResponse
	IsOpen int `json:"is_open"` //仅 errcode==0 (即调用成功) 时返回，0 表示开关处于关闭，1 表示开启成功（或开关已开启）
}

// MP端开关（仅内测期间使用）https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Temporary_MP_Switch.html
func (sdk *SDK) MpDraftSwitch(ctx context.Context, checkonly int) (*DraftSwitchRsp, error) {
	req := &DraftSwitchRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/draft/switch?access_token=%s", sdk.AccessToken)
	if checkonly == 1 {
		uri = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/draft/switch?access_token=%s&checkonly=%d", sdk.AccessToken, checkonly)
	}

	if err := common.DoRequestPost(ctx, uri, nil, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, common.ToError(req.ErrCode, req.ErrMsg)
	}

	return req, nil
}
