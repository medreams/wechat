package official

type SDK struct {
	Appid       string
	Secret      string
	AccessToken string
}

func New(appid, secret, token string) *SDK {
	return &SDK{
		Appid:       appid,
		Secret:      secret,
		AccessToken: token,
	}
}

func (sdk *SDK) UpdateAccessToken(accessToken string) error {
	sdk.AccessToken = accessToken
	return nil
}
