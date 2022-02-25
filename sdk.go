package wechat

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
	"github.com/medreams/wechat/mini"
	"github.com/medreams/wechat/official"
	"github.com/medreams/wechat/open"
)

type WeChatSDK struct {
	ctx       context.Context
	AppId     string
	AppSecret string
	tokenInfo common.WxAccessToken
}

func NewWeChatSDK(ctx context.Context, appId, appSecret string, isAccessToken ...bool) *WeChatSDK {

	sdk := &WeChatSDK{
		ctx:       context.Background(),
		AppId:     appId,
		AppSecret: appSecret,
	}

	//如果需要自动获取access_token，则自动获取
	if len(isAccessToken) > 0 && isAccessToken[0] {
		acc, err := common.GetAccessToken(sdk.ctx, sdk.AppId, sdk.AppSecret)
		if err != nil {
			fmt.Println(err)
		}
		if acc != nil {
			sdk.SetAccessToken(*acc)
		}
	}

	return sdk
}

//小程序
func (sdk *WeChatSDK) NewMini() *mini.SDK {
	return mini.New(sdk.AppId, sdk.AppSecret, sdk.tokenInfo.AccessToken)
}

//公众号
func (sdk *WeChatSDK) NewOfficial() *official.SDK {
	return official.New(sdk.AppId, sdk.AppSecret, sdk.tokenInfo.AccessToken)
}

//开放平台
func (sdk *WeChatSDK) NewOpen() *open.SDK {
	return open.New(sdk.AppId, sdk.AppSecret, sdk.tokenInfo.AccessToken)
}

func (sdk *WeChatSDK) SetAccessToken(token common.WxAccessToken) (err error) {
	sdk.tokenInfo = token
	return nil
}

func (sdk *WeChatSDK) GetAccessToken() (access_token string) {
	return sdk.tokenInfo.AccessToken
}
