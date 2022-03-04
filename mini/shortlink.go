package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type WxShortLink struct {
	Link    string `json:"link,omitempty"`
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func (sdk *SDK) GetWxShortLink(ctx context.Context, pageUrl, title string, isPeermanent bool) (link *WxShortLink, err error) {

	bodyMap := make(map[string]interface{})
	bodyMap["page_url"] = pageUrl
	bodyMap["page_title"] = title
	bodyMap["is_permanent"] = isPeermanent

	wxShortLink := &WxShortLink{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/genwxashortlink?access_token=%s", sdk.AccessToken)

	if err = common.DoRequestPost(ctx, uri, bodyMap, wxShortLink); err != nil {
		return nil, fmt.Errorf("do request get wxShortLink: %w", err)
	}

	if wxShortLink.Errcode != 0 {
		return nil, fmt.Errorf("get wxShortLink error: %d, %s", wxShortLink.Errcode, wxShortLink.Errmsg)
	}

	return wxShortLink, nil
}
