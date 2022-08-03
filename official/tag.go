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

//创建标签
func (s *SDK) CreateUserTag(ctx context.Context, tagName string) (tag *Tag, err error) {

	params := map[string]interface{}{
		"tag": map[string]interface{}{
			"name": tagName,
		},
	}

	tag = &Tag{}

	url := "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=" + s.AccessToken

	if err = common.DoRequestPost(ctx, url, params, tag); err != nil {
		return nil, fmt.Errorf("do request create tag: %w", err)
	}

	return tag, nil
}

//编辑标签
func (s *SDK) UpdateUserTag(ctx context.Context, tagId int, tagName string) (rst *common.WxCommonResponse, err error) {

	params := map[string]interface{}{
		"tag": map[string]interface{}{
			"id":   tagId,
			"name": tagName,
		},
	}

	url := "https://api.weixin.qq.com/cgi-bin/tags/update?access_token=" + s.AccessToken

	if err = common.DoRequestPost(ctx, url, params, rst); err != nil {
		return nil, fmt.Errorf("do request get tags: %w", err)
	}

	return rst, nil
}

// 删除标签
func (s *SDK) DeleteUserTag(ctx context.Context, tagId int) (rst *common.WxCommonResponse, err error) {

	params := map[string]interface{}{
		"tag": map[string]interface{}{
			"id": tagId,
		},
	}

	url := "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=" + s.AccessToken

	if err = common.DoRequestPost(ctx, url, params, rst); err != nil {
		return nil, fmt.Errorf("do request delete tags: %w", err)
	}

	return rst, nil
}

// 获取标签下粉丝列表
func (s *SDK) GetUserTagUserList(ctx context.Context, tagId int, nextOpenid string) (list *UserOpenidList, err error) {
	params := map[string]interface{}{
		"tagid":       tagId,
		"next_openid": nextOpenid,
	}

	list = &UserOpenidList{}

	url := "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=" + s.AccessToken
	if err = common.DoRequestPost(ctx, url, params, list); err != nil {
		return nil, fmt.Errorf("do request get tag userlist: %w", err)
	}

	return list, nil
}

// 批量为用户打标签
func (s *SDK) SetUserTagBatch(ctx context.Context, tagId int, openids []string) (rst *common.WxCommonResponse, err error) {
	params := map[string]interface{}{
		"tagid":       tagId,
		"openid_list": openids,
	}

	url := "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=" + s.AccessToken
	if err = common.DoRequestPost(ctx, url, params, rst); err != nil {
		return nil, fmt.Errorf("do request batch set user tag: %w", err)
	}

	return rst, nil
}

// 批量为用户取消标签
func (s *SDK) UnSetUserTagBatch(ctx context.Context, tagId int, openids []string) (rst *common.WxCommonResponse, err error) {
	params := map[string]interface{}{
		"tagid":       tagId,
		"openid_list": openids,
	}

	url := "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=" + s.AccessToken
	if err = common.DoRequestPost(ctx, url, params, rst); err != nil {
		return nil, fmt.Errorf("do request batch set user tag: %w", err)
	}

	return rst, nil
}

type TagIdListData struct {
	TagIdList []int `json:"tagid_list"`
}

// 获取用户身上的标签列表
func (s *SDK) GetUserUserTagIdList(ctx context.Context, openid string) (tagIdList *TagIdListData, err error) {
	params := map[string]interface{}{
		"openid": openid,
	}
	tagIdList = &TagIdListData{}
	url := "https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=" + s.AccessToken
	if err = common.DoRequestPost(ctx, url, params, tagIdList); err != nil {
		return nil, fmt.Errorf("do request get user usertag id list: %w", err)
	}

	return tagIdList, nil
}
