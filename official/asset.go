package official

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
	"github.com/medreams/wechat/pkg/util"
)

// MediaType 媒体文件类型
type MediaType string

const (
	// MediaTypeImage 媒体文件:图片
	MediaTypeImage MediaType = "image"
	// MediaTypeVoice 媒体文件:声音
	MediaTypeVoice MediaType = "voice"
	// MediaTypeVideo 媒体文件:视频
	MediaTypeVideo MediaType = "video"
	// MediaTypeThumb 媒体文件:缩略图
	MediaTypeThumb MediaType = "thumb"
	// MediaTypeImageTex 永久素材上传图文中的图片，不占用素材库数量限制
	MediaTypeNewsImage MediaType = "newsimage"
	// MediaTypeNew 素材类型：图文
	MediaTypeNews MediaType = "news"
)

type UploadTempAssetsRsp struct {
	common.WxCommonResponse
	Type      string `json:"type"`       //媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb，主要用于视频与音乐格式的缩略图）
	MediaID   string `json:"media_id"`   //媒体文件上传后，获取标识
	CreatedAt int    `json:"created_at"` //媒体文件上传时间戳
}

// 上传临时素材 媒体文件在微信后台保存时间为3天，即3天后media_id失效 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
func (sdk *SDK) UploadTempAssets(ctx context.Context, fileType MediaType, file *util.File) (*UploadTempAssetsRsp, error) {

	bodyMap := make(common.BodyMap)
	bodyMap.SetFormFile("media", file)

	fmt.Printf("bm : %#+v\n", bodyMap)

	req := &UploadTempAssetsRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", sdk.AccessToken, fileType)

	if err := common.DoUploadFile(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

type GetAssetsRsp struct {
	common.WxCommonResponse
	VideoURL string `json:"video_url"`
}

// 获取临时素材 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html
func (sdk *SDK) GetTempAssets(ctx context.Context, mediaId string) (*GetAssetsRsp, error) {

	req := &GetAssetsRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s", sdk.AccessToken, mediaId)

	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

type UploadPermanentAssetsRsp struct {
	common.WxCommonResponse
	MediaID string `json:"media_id"` //媒体文件上传后，获取标识
	Url     string `json:"url"`      //新增的图片素材的图片URL（仅新增图片素材时会返回该字段）
}

// 视频素材描述
type VideDescription struct {
	Title        string `json:"title"`        //视频素材的标题
	Introduction string `json:"introduction"` //视频素材的描述
}

// 上传永久图片素材 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (sdk *SDK) UploadPermanentAssets(ctx context.Context, fileType MediaType, file *util.File, vd *VideDescription) (*UploadPermanentAssetsRsp, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.SetFormFile("media", file)
	uri := ""
	if fileType == MediaTypeNewsImage {
		//上传图文消息内的图片
		uri = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s", sdk.AccessToken)
	} else {
		//新增其他类型永久素材
		uri = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s", sdk.AccessToken, fileType)

		if fileType == MediaTypeVideo {
			bodyMap.SetBodyMap("description", func(b common.BodyMap) {
				b.Set("title", vd.Title)
				b.Set("introduction", vd.Introduction)
			})
		}
	}
	req := &UploadPermanentAssetsRsp{}
	if err := common.DoUploadFile(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

type GetMaterialRsp struct {
	common.WxCommonResponse
	NewsItem []MaterialNewsItem `json:"news_item"`
}

type MaterialNewsItem struct {
	Title            string `json:"title"`              //图文消息的标题
	ThumbMediaID     string `json:"thumb_media_id"`     //图文消息的封面图片素材id（必须是永久mediaID）
	ShowCoverPic     int    `json:"show_cover_pic"`     //是否显示封面，0为false，即不显示，1为true，即显示
	Author           string `json:"author"`             //作者
	Digest           string `json:"digest"`             //图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空
	Content          string `json:"content"`            //图文消息的具体内容，支持 HTML 标签，必须少于2万字符，小于1M，且此处会去除JS
	Url              string `json:"url"`                //图文页的URL
	ContentSourceURL string `json:"content_source_url"` //图文消息的原文地址，即点击“阅读原文”后的URL
	//视频
	Description string `json:"description"`
	DownUrl     string `json:"down_url"`
}

// 获取永久素材 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (sdk *SDK) GetPermanentAssets(ctx context.Context, mediaId string) (*GetMaterialRsp, error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("media_id", mediaId)

	req := &GetMaterialRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

// 删除永久素材 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Deleting_Permanent_Assets.html
func (sdk *SDK) DelPermanentAssets(ctx context.Context, mediaId string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("media_id", mediaId)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

type GetPermanentAssetsCountRsp struct {
	VoiceCount int `json:"voice_count"` //语音总数量
	VideoCount int `json:"video_count"` //视频总数量
	ImageCount int `json:"image_count"` //图片总数量
	NewsCount  int `json:"news_count"`  //图文总数量
}

// 获取素材总数 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_the_total_of_all_materials.html
func (sdk *SDK) GetPermanentAssetsCount(ctx context.Context) error {

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

type PermanentAssetsListRsp struct {
	common.WxCommonResponse
	TotalCount int                   `json:"total_count"`
	ItemCount  int                   `json:"item_count"`
	Item       []PermanentAssetsItem `json:"item"`
}

type PermanentAssetsNewsItem struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	URL              string `json:"url"`
	ContentSourceURL string `json:"content_source_url"`
}

// 图片内容
type PermanentAssetsContent struct {
	NewsItem []PermanentAssetsNewsItem `json:"news_item"`
}

type PermanentAssetsItem struct {
	MediaID    string                 `json:"media_id"`
	Name       string                 `json:"name"` //
	Url        string                 `json:"url"`
	Content    PermanentAssetsContent `json:"content"` //图文
	UpdateTime string                 `json:"update_time"`
}

// 获取素材列表 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func (sdk *SDK) GetPermanentAssetsList(ctx context.Context, mediaType MediaType, offset, count int) (*PermanentAssetsListRsp, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("type", mediaType)
	bodyMap.Set("offset", offset)
	bodyMap.Set("count", count)

	req := &PermanentAssetsListRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}
