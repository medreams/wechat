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
	Tag TagData `json:"tag,omitempty"`
}

type Tags struct {
	Tags []TagData `json:"tags,omitempty"`
}

// 创建标签
func (sdk *SDK) CreateUserTag(ctx context.Context, tagName string) (req *Tag, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.SetBodyMap("tag", func(b common.BodyMap) {
		b.Set("name", tagName)
	})

	req = &Tag{}

	url := "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=" + sdk.AccessToken

	if err = common.DoRequestPost(ctx, url, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request create tag: %w", err)
	}

	return req, nil
}

// 编辑标签
func (s *SDK) UpdateUserTag(ctx context.Context, tagId int, tagName string) (req *common.WxCommonResponse, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.SetBodyMap("tag", func(b common.BodyMap) {
		b.Set("id", tagId)
		b.Set("name", tagName)
	})

	url := "https://api.weixin.qq.com/cgi-bin/tags/update?access_token=" + s.AccessToken

	if err = common.DoRequestPost(ctx, url, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get tags: %w", err)
	}

	return req, nil
}

// 删除标签
func (sdk *SDK) DeleteUserTag(ctx context.Context, tagId int) (req *common.WxCommonResponse, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.SetBodyMap("tag", func(b common.BodyMap) {
		b.Set("id", tagId)
	})

	url := "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=" + sdk.AccessToken

	if err = common.DoRequestPost(ctx, url, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request delete tags: %w", err)
	}

	return req, nil
}

// 获取标签下粉丝列表
func (sdk *SDK) GetUserTagUserList(ctx context.Context, tagId int, nextOpenid string) (req *UserOpenidList, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("tagid", tagId)
	bodyMap.Set("next_openid", nextOpenid)

	req = &UserOpenidList{}

	url := "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=" + sdk.AccessToken
	if err = common.DoRequestPost(ctx, url, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get tag userlist: %w", err)
	}

	return req, nil
}

// 批量为用户打标签
func (sdk *SDK) SetUserTagBatch(ctx context.Context, tagId int, openids []string) (req *common.WxCommonResponse, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("tagid", tagId)
	bodyMap.Set("openid_list", openids)

	url := "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=" + sdk.AccessToken
	if err = common.DoRequestPost(ctx, url, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request batch set user tag: %w", err)
	}

	return req, nil
}

// 批量为用户取消标签
func (sdk *SDK) UnSetUserTagBatch(ctx context.Context, tagId int, openids []string) (req *common.WxCommonResponse, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("tagid", tagId)
	bodyMap.Set("openid_list", openids)

	url := "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=" + sdk.AccessToken
	if err = common.DoRequestPost(ctx, url, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request batch set user tag: %w", err)
	}

	return req, nil
}

type TagIdListData struct {
	TagIdList []int `json:"tagid_list"`
}

// 获取用户身上的标签列表
func (sdk *SDK) GetUserUserTagIdList(ctx context.Context, openid string) (req *TagIdListData, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("openid", openid)

	req = &TagIdListData{}
	url := "https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=" + sdk.AccessToken
	if err = common.DoRequestPost(ctx, url, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get user usertag id list: %w", err)
	}

	return req, nil
}
