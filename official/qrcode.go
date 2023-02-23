package official

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

// 帐号管理 /生成带参数的二维码

type QrCodeRsp struct {
	common.WxCommonResponse
	Ticket        string `json:"ticket"`         //获取的二维码ticket，凭借此 ticket 可以在有效时间内换取二维码。
	ExpireSeconds int    `json:"expire_seconds"` //该二维码有效时间，以秒为单位。 最大不超过2592000（即30天）。
	URL           string `json:"url"`            //二维码图片解析后的地址，开发者可根据该地址自行生成需要的二维码图片
}

// 生成带参数的二维码 actionName QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值 https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
func (sdk *SDK) GenerateQRCode(ctx context.Context, expireSeconds int, actionName string, sceneId int, sceneStr string) (*QrCodeRsp, error) {
	bodyMap := make(common.BodyMap)
	bodyMap.Set("expire_seconds", expireSeconds)
	bodyMap.Set("action_name", actionName)
	bodyMap.SetBodyMap("action_info", func(bm common.BodyMap) {
		bm.SetBodyMap("scene", func(b common.BodyMap) {
			if sceneId > 0 {
				b.Set("scene_id", sceneId)
			}
			if sceneStr != "" {
				b.Set("scene_id", sceneId)
			}
		})
	})

	req := &QrCodeRsp{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s", sdk.AccessToken)

	if err := common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if req.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode(%d) != 0", req.ErrCode)
	}

	return req, nil
}

// 通过 ticket 换取二维码
func (sdk *SDK) TicketGetQRCode(ctx context.Context, ticket string) ([]byte, error) {
	uri := fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s", ticket)

	bs, err := common.DoRequestGetByte(ctx, uri)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	return bs, nil
}
