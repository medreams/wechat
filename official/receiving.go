package official

import (
	"encoding/xml"
	"fmt"
)

/**
消息类型 MsgType

文本为 text
图片为 image
语音为 voice
视频为 video
小视频为 shortvideo
地理位置为 location
事件为 event
**/

/**
事件类型 MsgType 为 event
菜单事件类型 Event
CLICK 点击菜单拉取消息时的事件推送
VIEW 点击菜单跳转链接时的事件推送
scancode_push 扫码推事件的事件推送
scancode_waitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
pic_sysphoto 弹出系统拍照发图的事件推送
pic_photo_or_album 弹出拍照或者相册发图的事件推送
pic_weixin 弹出微信相册发图器的事件推送
location_select 弹出地理位置选择器的事件推送
view_miniprogram 点击菜单跳转小程序的事件推送

其它事件类型 Event
subscribe/unsubscribe 关注/取消关注事件
SCAN 扫描带参数二维码事件
LOCATION 上报地理位置事件
**/

// 扫码推事件
type EventScanCodeInfo struct {
	ScanType   string `xml:"ScanType"`
	ScanResult string `xml:"ScanResult"`
}

// 弹出系统拍照发图的事件
type EventSendPicsInfo struct {
	Count   string `xml:"Count"`
	PicList struct {
		Item struct {
			PicMd5Sum string `xml:"PicMd5Sum"`
		} `xml:"item"`
	} `xml:"PicList"`
}

type EventSendLocationInfo struct {
	LocationX string `xml:"Location_X"`
	LocationY string `xml:"Location_Y"`
	Scale     string `xml:"Scale"`
	Label     string `xml:"Label"`
	Poiname   string `xml:"Poiname"`
}

type ReceivingEventMsg struct {
	Event            string                `xml:"Event,omitempty"`            //事件类型，CLICK
	EventKey         string                `xml:"EventKey,omitempty"`         //事件 KEY 值，与自定义菜单接口中 KEY 值对应
	MenuId           string                `xml:"MenuId,omitempty"`           //指菜单ID，如果是个性化菜单，则可以通过这个字段，知道是哪个规则的菜单被点击了。
	ScanCodeInfo     EventScanCodeInfo     `xml:"ScanCodeInfo,omitempty"`     //扫码推事件且弹出“消息接收中”提示框的事件推送
	SendPicsInfo     EventSendPicsInfo     `xml:"SendPicsInfo,omitempty"`     //弹出系统拍照发图的事件推送
	SendLocationInfo EventSendLocationInfo `xml:"SendLocationInfo,omitempty"` //弹出地理位置选择器的事件推送
	Ticket           string                `xml:"Ticket,omitempty"`           //二维码的ticket，可用来换取二维码图片
	Latitude         string                `xml:"Latitude,omitempty"`         //地理位置纬度
	Longitude        string                `xml:"Longitude,omitempty"`        //地理位置经度
	Precision        string                `xml:"Precision,omitempty"`        //地理位置精度
}

// 图片消息内容
type ReceivingImageMsg struct {
	PicUrl  string `json:"pic_url"`
	MediaId string `json:"media_id"`
}

// 语音信息内容
type ReceivingVoiceMsg struct {
	MediaId     string `json:"media_id"`
	Format      string `xml:"Format,omitempty"`      //语音格式，如amr，speex等
	Recognition string `xml:"Recognition,omitempty"` //语音识别结果，UTF8编码
}

// 视频信息
type ReceivingVideoMsg struct {
	MediaId      string `json:"media_id"`
	ThumbMediaId string `xml:"ThumbMediaId,omitempty"` //视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
}

// 地理位置信息
type ReceivingLocationMsg struct {
	LocationX string `xml:"Location_X,omitempty"` //地理位置纬度
	LocationY string `xml:"Location_Y,omitempty"` //地理位置经度
	Scale     string `xml:"Scale,omitempty"`      //地图缩放大小
	Label     string `xml:"Label,omitempty"`      //地理位置信息
}

// 链接消息
type ReceivingLinkMsg struct {
	Title       string `xml:"Title,omitempty"`       //消息标题
	Description string `xml:"Description,omitempty"` //消息描述
	URL         string `xml:"Url,omitempty"`         //消息链接
}

type ReceivingMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`   //开发者微信号
	FromUserName string   `xml:"FromUserName"` //发送方帐号（一个OpenID）
	CreateTime   string   `xml:"CreateTime"`   //消息创建时间 （整型）
	MsgType      string   `xml:"MsgType"`      //消息类型，文本为text

	//文本信息
	Content string `xml:"Content,omitempty"` //文本消息内容
	//图片信息
	PicUrl string `xml:"PicUrl,omitempty"` //图片链接（由系统生成）
	//语音视频消息
	Format       string `xml:"Format,omitempty"`       //语音格式，如amr，speex等
	MediaId      string `xml:"MediaId,omitempty"`      //图片消息媒体id，可以调用获取临时素材接口拉取数据。
	Recognition  string `xml:"Recognition,omitempty"`  //语音识别结果，UTF8编码
	ThumbMediaId string `xml:"ThumbMediaId,omitempty"` //视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	//地理位置消息
	ReceivingLocationMsg
	//链接消息
	ReceivingLinkMsg

	MsgId     string `xml:"MsgId,omitempty"`     //消息id，64位整型
	MsgDataId string `xml:"MsgDataId,omitempty"` //消息的数据ID（消息如果来自文章时才有）
	Idx       string `xml:"Idx,omitempty"`       //多图文时第几篇文章，从1开始（消息如果来自文章时才有）
	//事件部分消息
	ReceivingEventMsg
}

func NewReceivingMessage() *ReceivingMessage {
	return &ReceivingMessage{}
}

// 解析接收到的参数
func (rm *ReceivingMessage) Unmarshal(xmldata []byte) (*ReceivingMessage, error) {
	req := &ReceivingMessage{}
	err := xml.Unmarshal(xmldata, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (rm *ReceivingMessage) UnmarshalForString(xmldata string) (*ReceivingMessage, error) {
	return rm.Unmarshal([]byte(xmldata))
}

// 获取消息类型
func (rm *ReceivingMessage) GetMsgType() string {
	return rm.MsgType
}

// 获取公众号的原始ID
func (rm *ReceivingMessage) GetOfficialId() string {
	return rm.ToUserName
}

// 获取发送者的openid
func (rm *ReceivingMessage) GetSendOpenid() string {
	return rm.FromUserName
}

// 获取消息创建时间
func (rm *ReceivingMessage) GetCreateTime() string {
	return rm.CreateTime
}

// 获取事件消息
func (rm *ReceivingMessage) GetEventMsg() (*ReceivingEventMsg, error) {

	if rm.GetMsgType() == "event" {
		eventMsg := &ReceivingEventMsg{
			Event:        rm.Event,
			EventKey:     rm.EventKey,
			MenuId:       rm.MenuId,
			ScanCodeInfo: rm.ScanCodeInfo,
			SendPicsInfo: rm.SendPicsInfo,
			Ticket:       rm.Ticket,
			Latitude:     rm.Latitude,
			Longitude:    rm.Longitude,
			Precision:    rm.Precision,
		}

		return eventMsg, nil
	}

	return nil, fmt.Errorf("该消息不是事件消息，消息类型为 %s", rm.GetMsgType())
}

// 获取文本信息中的内容
func (rm *ReceivingMessage) GetTextMsgContent() (string, error) {
	if rm.GetMsgType() == "text" {
		return rm.Content, nil
	}

	return "", fmt.Errorf("该消息不是文本类消息，消息类型为 %s", rm.GetMsgType())
}

// 获取图片信息中的内容
func (rm *ReceivingMessage) GetImageMsgContent() (*ReceivingImageMsg, error) {

	if rm.GetMsgType() == "image" {
		return &ReceivingImageMsg{
			PicUrl:  rm.PicUrl,
			MediaId: rm.MediaId,
		}, nil
	}

	return nil, fmt.Errorf("该消息不是图片类消息，消息类型为 %s", rm.GetMsgType())
}

// 获取语音信息中的内容
func (rm *ReceivingMessage) GetVoiceMsgContent() (*ReceivingVoiceMsg, error) {

	if rm.GetMsgType() == "voice" {
		return &ReceivingVoiceMsg{
			MediaId:     rm.MediaId,
			Format:      rm.Format,
			Recognition: rm.Recognition,
		}, nil
	}

	return nil, fmt.Errorf("该消息不是语音类消息，消息类型为 %s", rm.GetMsgType())
}

// 获取视频信息内容
func (rm *ReceivingMessage) GetVideoMsgContent() (*ReceivingVideoMsg, error) {

	if rm.GetMsgType() == "video" || rm.GetMsgType() == "shortvideo" {
		return &ReceivingVideoMsg{
			MediaId:      rm.MediaId,
			ThumbMediaId: rm.ThumbMediaId,
		}, nil
	}

	return nil, fmt.Errorf("该消息不是视频类消息，消息类型为 %s", rm.GetMsgType())
}

// 获取地理位置信息内容
func (rm *ReceivingMessage) GetLocationMsgContent() (*ReceivingLocationMsg, error) {

	if rm.GetMsgType() == "location" {
		return &ReceivingLocationMsg{
			LocationX: rm.LocationX,
			LocationY: rm.LocationY,
			Scale:     rm.Scale,
			Label:     rm.Label,
		}, nil
	}

	return nil, fmt.Errorf("该消息不是地理位置类消息，消息类型为 %s", rm.GetMsgType())
}

// 获取链接信息内容
func (rm *ReceivingMessage) GetLinkMsgContent() (*ReceivingLinkMsg, error) {

	if rm.GetMsgType() == "location" {
		return &ReceivingLinkMsg{
			Title:       rm.Title,
			Description: rm.Description,
			URL:         rm.URL,
		}, nil
	}

	return nil, fmt.Errorf("该消息不是连接类消息，消息类型为 %s", rm.GetMsgType())
}
