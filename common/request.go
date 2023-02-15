package common

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/medreams/wechat/pkg/util"
	"github.com/medreams/wechat/pkg/xhttp"
)

type WxCommonResponse struct {
	ErrCode int64  `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

func DoRequestGet(c context.Context, uri string, ptr interface{}) (err error) {
	httpClient := xhttp.NewClient()

	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Get(uri).EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(GET, %s)：%w", uri, err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("StatusCode(%d) != 200", res.StatusCode)
	}

	if err := CheckRequestError(bs); err != nil {
		return err
	}

	fmt.Println("返回结果：", string(bs))

	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}

func DoRequestPost(c context.Context, uri string, body map[string]interface{}, ptr interface{}) (err error) {

	bs, err := DoRequestPostGetByte(c, uri, body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}
	return
}

func DoRequestPostGetByte(c context.Context, uri string, body map[string]interface{}) (bs []byte, err error) {
	httpClient := xhttp.NewClient()

	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Post(uri).SendBodyMap(body).EndBytes(c)
	if err != nil {
		return nil, fmt.Errorf("http.request(POST, %s)：%w", uri, err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode(%d) != 200", res.StatusCode)
	}

	if err := CheckRequestError(bs); err != nil {
		return nil, err
	}

	return bs, nil
}

func DoUploadFile(c context.Context, uri string, body map[string]interface{}, ptr interface{}) error {
	httpClient := xhttp.NewClient()
	httpClient.Header.Add(xhttp.HeaderRequestID, fmt.Sprintf("%s-%d", util.RandomString(21), time.Now().Unix()))
	res, bs, err := httpClient.Type(xhttp.TypeMultipartFormData).
		Post(uri).
		SendMultipartBodyMap(body).
		EndBytes(c)
	if err != nil {
		return fmt.Errorf("http.request(POST, %s)：%w", uri, err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("StatusCode(%d) != 200", res.StatusCode)
	}

	if err := CheckRequestError(bs); err != nil {
		return err
	}

	if err = json.Unmarshal(bs, ptr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), ptr, err)
	}

	return nil
}

func CheckRequestError(bs []byte) error {

	msg := &WxCommonResponse{}
	json.Unmarshal(bs, msg)
	if msg.ErrCode != 0 {
		return fmt.Errorf("ErrCode(%d),ErrMsg(%s)", msg.ErrCode, msg.ErrMsg)
	}

	return nil
}
