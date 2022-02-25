package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type PaidUnionId struct {
	Unionid string `json:"unionid,omitempty"` // 用户在开放平台的唯一标识符
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

func (s *SDK) GetPaidUnionId(c context.Context, openid string) (unionId string, err error) {

	uri := "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=ACCESS_TOKEN&openid=OPENID"

	PaidUnionId := &PaidUnionId{}
	if err = common.DoRequestGet(c, uri, PaidUnionId); err != nil {
		return "", fmt.Errorf("do request get unionId: %w", err)
	}

	if PaidUnionId.Errcode != 0 {
		return "", fmt.Errorf("get unionId error: %d, %s", PaidUnionId.Errcode, PaidUnionId.Errmsg)
	}

	return PaidUnionId.Unionid, nil
}
