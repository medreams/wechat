package open

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
