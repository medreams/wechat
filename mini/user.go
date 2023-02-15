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

func (s *SDK) GetPaidUnionId(c context.Context, openid string) (unionId string, err error) {

	PaidUnionId := &PaidUnionId{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/getpaidunionid?access_token=%s&openid=%s", s.AccessToken, openid)

	if err = common.DoRequestGet(c, uri, PaidUnionId); err != nil {
		return "", fmt.Errorf("do request get unionId: %w", err)
	}

	if PaidUnionId.ErrCode != 0 {
		return "", fmt.Errorf("get unionId error: %d, %s", PaidUnionId.ErrCode, PaidUnionId.ErrMsg)
	}

	return PaidUnionId.Unionid, nil
}
