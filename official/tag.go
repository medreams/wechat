package official

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type TagData struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Count int    `json:"count,omitempty"`
}

type Tag struct {
	common.WxCommonResponse
	Tag TagData `json:"tag,omitempty"`
}

// 创建标签
func (sdk *SDK) CreateUserTag(ctx context.Context, tagName string) (*Tag, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.SetBodyMap("tag", func(b common.BodyMap) {
		b.Set("name", tagName)
	})

	req := &Tag{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request create tag: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

type Tags struct {
	common.WxCommonResponse
	Tags []TagData `json:"tags,omitempty"`
}

// 获取公众号已创建的标签
func (sdk *SDK) GetUserTagList(ctx context.Context) (*Tags, error) {
	req := &Tags{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestGet(ctx, uri, req); err != nil {
		return nil, fmt.Errorf("do request get access_token: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

// 编辑标签
func (sdk *SDK) UpdateUserTag(ctx context.Context, tagId int, tagName string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.SetBodyMap("tag", func(b common.BodyMap) {
		b.Set("id", tagId)
		b.Set("name", tagName)
	})

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request get tags: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 删除标签
func (sdk *SDK) DeleteUserTag(ctx context.Context, tagId int) error {
	bodyMap := make(common.BodyMap)
	bodyMap.SetBodyMap("tag", func(b common.BodyMap) {
		b.Set("id", tagId)
	})

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request delete tags: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 获取标签下粉丝列表
func (sdk *SDK) GetUserTagUserList(ctx context.Context, tagId int, nextOpenid string) (*UserOpenidList, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("tagid", tagId)
	bodyMap.Set("next_openid", nextOpenid)

	req := &UserOpenidList{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get tag userlist: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

// 批量为用户打标签
func (sdk *SDK) SetUserTagBatch(ctx context.Context, tagId int, openids []string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("tagid", tagId)
	bodyMap.Set("openid_list", openids)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request batch set user tag: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

// 批量为用户取消标签
func (sdk *SDK) UnSetUserTagBatch(ctx context.Context, tagId int, openids []string) error {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("tagid", tagId)
	bodyMap.Set("openid_list", openids)

	req := &common.WxCommonResponse{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return fmt.Errorf("do request batch set user tag: %w", err)
	}

	if req.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return nil
}

type TagIdListData struct {
	common.WxCommonResponse
	TagIdList []int `json:"tagid_list"`
}

// 获取用户身上的标签列表
func (sdk *SDK) GetUserUserTagIdList(ctx context.Context, openid string) (*TagIdListData, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("openid", openid)

	req := &TagIdListData{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get user usertag id list: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}
