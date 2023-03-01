package official

import (
	"context"
	"errors"
	"fmt"

	"github.com/medreams/wechat/common"
)

// 获取公众号的黑名单列表 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
func (sdk *SDK) GetUserBlackList(ctx context.Context, beginOpenid string) (list *UserOpenidList, err error) {
	bodyMap := make(common.BodyMap)
	bodyMap["begin_openid"] = beginOpenid

	list = &UserOpenidList{}
	uri := "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=" + sdk.AccessToken
	if err = common.DoRequestPost(ctx, uri, bodyMap, list); err != nil {
		return nil, fmt.Errorf("do request get user usertag id list: %w", err)
	}

	return list, nil
}

// 拉黑用户(一次20个)
func (sdk *SDK) AddUsersToBlackList(ctx context.Context, openids []string) (rst *common.WxCommonResponse, err error) {

	if len(openids) > 20 {
		return nil, errors.New("一次最多20个openid")
	}

	bodyMap := make(common.BodyMap)
	bodyMap.Set("openid_list", openids)

	uri := "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=" + sdk.AccessToken
	if err = common.DoRequestPost(ctx, uri, bodyMap, rst); err != nil {
		return nil, fmt.Errorf("do request get user usertag id list: %w", err)
	}

	return rst, nil
}

// 取消拉黑用户(一次20个)
func (sdk *SDK) CancelUsersFromBlackList(ctx context.Context, openids []string) (req *common.WxCommonResponse, err error) {

	if len(openids) > 20 {
		return nil, errors.New("一次最多20个openid")
	}

	bodyMap := make(common.BodyMap)
	bodyMap.Set("openid_list", openids)

	uri := "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=" + sdk.AccessToken
	if err = common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get user usertag id list: %w", err)
	}

	return req, nil
}
