package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type WxCode2Session struct {
	Openid     string `json:"openid,omitempty"`      // 用户唯一标识
	SessionKey string `json:"session_key,omitempty"` // 会话密钥
	Unionid    string `json:"unionid,omitempty"`     // 用户在开放平台的唯一标识符
	Errcode    int    `json:"errcode,omitempty"`     // 错误码
	Errmsg     string `json:"errmsg,omitempty"`      // 错误信息
}

func (s *SDK) Code2Session(c context.Context, code string) (session *WxCode2Session, err error) {

	uri := "https://api.weixin.qq.com/sns/jscode2session?appid=" + s.Appid + "&secret=" + s.Secret + "&js_code=" + code + "&grant_type=authorization_code"

	session = &WxCode2Session{}
	if err = common.DoRequestGet(c, uri, session); err != nil {
		return nil, fmt.Errorf("do request get session: %w", err)
	}

	if session.Errcode != 0 {
		return nil, fmt.Errorf("get session error: %d, %s", session.Errcode, session.Errmsg)
	}

	return session, nil

}
