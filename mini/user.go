package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type PaidUnionId struct {
	common.WxCommonResponse
	Unionid string `json:"unionid,omitempty"` // 用户在开放平台的唯一标识符
}

func (sdk *SDK) GetPaidUnionId(c context.Context, openid string) (unionId string, err error) {

	req := &PaidUnionId{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s", sdk.AccessToken, openid)

	if err = common.DoRequestGet(c, uri, req); err != nil {
		return "", fmt.Errorf("do request get unionId: %w", err)
	}

	if req.ErrCode != 0 {
		return "", fmt.Errorf("get unionId error: %d, %s", req.ErrCode, req.ErrMsg)
	}

	return req.Unionid, nil
}
