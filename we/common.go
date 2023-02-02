package we

type SDK struct {
	Appid          string
	Secret         string
	AccessToken    string
	WebAccessToekn string
}

func New(appid, secret, token string) *SDK {
	return &SDK{
		Appid:       appid,
		Secret:      secret,
		AccessToken: token,
	}
}

func (sdk *SDK) UpdateWebAccessToken(webAccessToken string) error {
	sdk.WebAccessToekn = webAccessToken
	return nil
}
