package mini

import (
	"encoding/json"
	"errors"

	"github.com/medreams/wechat/common"
)

// 微信小程序解密后 用户信息
type WxUserInfo struct {
	OpenId    string         `json:"openId,omitempty"`
	NickName  string         `json:"nickName,omitempty"`
	Gender    int            `json:"gender,omitempty"`
	City      string         `json:"city,omitempty"`
	Province  string         `json:"province,omitempty"`
	Country   string         `json:"country,omitempty"`
	AvatarUrl string         `json:"avatarUrl,omitempty"`
	UnionId   string         `json:"unionId,omitempty"`
	Language  string         `json:"language,omitempty"`
	Watermark *watermarkInfo `json:"watermark,omitempty"`
}

func (sdk *SDK) DecryptUserInfo(session_key, iv, encrypted_data string) (*WxUserInfo, error) {

	dataBytes, err := common.Dncrypt(encrypted_data, session_key, iv)
	if err != nil {
		return nil, err
	}

	var result WxUserInfo
	err = json.Unmarshal(dataBytes, &result)

	watermark := result.Watermark
	if watermark.Appid != sdk.Appid {
		return nil, errors.New("invalid appid data")
	}

	return &result, err
}

func (sdk *SDK) DecryptUserPhone(session_key, iv, encrypted_data string) (*WxUserPhone, error) {

	dataBytes, err := common.Dncrypt(encrypted_data, session_key, iv)
	if err != nil {
		return nil, err
	}

	var result WxUserPhone
	err = json.Unmarshal(dataBytes, &result)

	watermark := result.Watermark
	if watermark.Appid != sdk.Appid {
		return nil, errors.New("invalid appid data")
	}

	return &result, err
}
