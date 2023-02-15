package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type WxCode2Session struct {
	common.WxCommonResponse
	Openid     string `json:"openid,omitempty"`      // 用户唯一标识
	SessionKey string `json:"session_key,omitempty"` // 会话密钥
	Unionid    string `json:"unionid,omitempty"`     // 用户在开放平台的唯一标识符
}

func (sdk *SDK) Code2Session(c context.Context, code string) (req *WxCode2Session, err error) {

	req = &WxCode2Session{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", sdk.Appid, sdk.Secret, code)

	if err = common.DoRequestGet(c, uri, req); err != nil {
		return nil, fmt.Errorf("do request get session: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("get session error: %d, %s", req.ErrCode, req.ErrMsg)
	}

	return req, nil

}
