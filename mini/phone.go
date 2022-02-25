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

func (sdk *SDK) MiniCodeGetPhone(ctx context.Context, phoneCode string) (phone *WxUserPhone, err error) {
	uri := "https://api.weixin.qq.com/wxa/genwxashortlink?access_token=" + sdk.AccessToken

	bodyMap := make(map[string]interface{})
	bodyMap["code"] = phoneCode

	phone = &WxUserPhone{}
	if err = common.DoRequestPost(ctx, uri, bodyMap, phone); err != nil {
		return nil, fmt.Errorf("do request get phone: %w", err)
	}

	if phone.Watermark.Appid != sdk.Appid {
		return nil, fmt.Errorf("get phone error: %s", "appid not match")
	}

	return phone, nil
}

func DecryptWeChatOpenDataGetPhone(sessionKey, encryptedData, iv string) (phone *WxUserPhone, err error) {
	phone = &WxUserPhone{}
	if err = common.DecryptOpenDataToStruct(sessionKey, encryptedData, iv, phone); err != nil {
		return nil, fmt.Errorf("decrypt wechat open data get phone: %w", err)
	}

	return phone, nil
}
