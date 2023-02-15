package official

import (
	"encoding/xml"
	"time"
)

// CDATA  使用该类型,在序列化为 xml 文本时文本会被解析器忽略
type CDATA string

// MarshalXML 实现自己的序列化方法
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

type ReplyMsgCommon struct {
	XMLName      xml.Name `xml:"xml"`
	Text         CDATA    `xml:",cdata"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
}
type ReplyText struct {
	Content *CDATA `xml:"Content"`
}
type ReplyMedia struct {
	MediaId     CDATA  `xml:"MediaId"`
	Title       *CDATA `xml:"Title,omitempty"`
	Description *CDATA `xml:"Description,omitempty"`
}

type ReplyMusic struct {
	Title        CDATA `xml:"Title"`
	Description  CDATA `xml:"Description,omitempty"`
	MusicUrl     CDATA `xml:"MusicUrl"`
	HQMusicUrl   CDATA `xml:"HQMusicUrl"`
	ThumbMediaId CDATA `xml:"ThumbMediaId"`
}

type ReplyArticleInfo struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
	URL         string `xml:"Url"`
}

type ReplyArticles struct {
	Item []ReplyArticleInfo `xml:"item"`
}

type ReplyMessage struct {
	ReplyMsgCommon
	ReplyText
	Image        *ReplyMedia    `xml:"Image,omitempty"`
	Voice        *ReplyMedia    `xml:"Voice,omitempty"`
	Video        *ReplyMedia    `xml:"Video,omitempty"`
	Music        *ReplyMusic    `xml:"Music,omitempty"`
	ArticleCount *int           `xml:"ArticleCount,omitempty"`
	Articles     *ReplyArticles `xml:"Articles,omitempty"`
}

func NewReplyMessage() *ReplyMessage {
	return &ReplyMessage{}
}

// 置文本信息的内容
func (rm *ReplyMessage) SetContent(text CDATA) {
	rm.CreateTime = time.Now().Unix()
	rm.MsgType = "text"
	rm.Content = &text
}

// 设置图片信息的mediaId
func (rm *ReplyMessage) SetImage(mediaId CDATA) {
	rm.CreateTime = time.Now().Unix()
	rm.MsgType = "image"
	rm.Image = &ReplyMedia{
		MediaId: mediaId,
	}
}

// 设置语音信息的mediaId
func (rm *ReplyMessage) SetVoice(mediaId CDATA) {
	rm.CreateTime = time.Now().Unix()
	rm.MsgType = "voice"
	rm.Voice = &ReplyMedia{
		MediaId: mediaId,
	}
}

// 设置视频信息
func (rm *ReplyMessage) SetVideo(mediaId, title, description CDATA) {
	rm.CreateTime = time.Now().Unix()
	rm.MsgType = "video"
	rm.Video = &ReplyMedia{
		MediaId:     mediaId,
		Title:       &title,
		Description: &description,
	}
}

// 设置音乐信息
func (rm *ReplyMessage) SetMusic(title, description, musicUrl, HQMusicUrl, thumbMediaId CDATA) {
	rm.CreateTime = time.Now().Unix()
	rm.MsgType = "music"
	rm.Music = &ReplyMusic{
		Title:        title,
		Description:  description,
		MusicUrl:     musicUrl,
		HQMusicUrl:   HQMusicUrl,
		ThumbMediaId: thumbMediaId,
	}
}

// 设置图文信息
func (rm *ReplyMessage) SetArticles(articleCount int, items []ReplyArticleInfo) {
	rm.CreateTime = time.Now().Unix()
	rm.MsgType = "news"
	rm.ArticleCount = &articleCount
	rm.Articles = &ReplyArticles{
		Item: items,
	}
}

func (rm *ReplyMessage) Marshal() []byte {
	rawXMLbyte, err := xml.Marshal(rm)
	if err != nil {
		return nil
	}
	return rawXMLbyte
}
