package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type watermarkInfo struct {
	Appid     string `json:"appid,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}

type WxUserPhone struct {
	PhoneNumber     string         `json:"phoneNumber,omitempty"`
	PurePhoneNumber string         `json:"purePhoneNumber,omitempty"`
	CountryCode     string         `json:"countryCode,omitempty"`
	Watermark       *watermarkInfo `json:"watermark,omitempty"`
}

func (sdk *SDK) MiniCodeGetPhone(ctx context.Context, phoneCode string) (*WxUserPhone, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("code", phoneCode)

	req := &WxUserPhone{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get phone: %w", err)
	}

	if req.Watermark.Appid != sdk.Appid {
		return nil, fmt.Errorf("get phone error: %s", "appid not match")
	}

	return req, nil
}
