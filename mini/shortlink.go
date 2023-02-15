package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type WxShortLink struct {
	common.WxCommonResponse
	Link string `json:"link,omitempty"`
}

func (sdk *SDK) GetWxShortLink(ctx context.Context, pageUrl, title string, isPeermanent bool) (req *WxShortLink, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("page_url", pageUrl)
	bodyMap.Set("page_title", title)
	bodyMap.Set("is_permanent", isPeermanent)

	req = &WxShortLink{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/wxa/genwxashortlink?access_token=%s", sdk.AccessToken)

	if err = common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get wxShortLink: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("get wxShortLink error: %d, %s", req.ErrCode, req.ErrMsg)
	}

	return req, nil
}
