package common

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	xaes "github.com/medreams/wechat/pkg/aes"
	"github.com/medreams/wechat/pkg/util"
)

// DecryptOpenDataToStruct 解密开放数据到结构体
//
//	encryptedData：包括敏感数据在内的完整用户信息的加密数据，小程序获取到
//	iv：加密算法的初始向量，小程序获取到
//	sessionKey：会话密钥，通过  gopay.Code2Session() 方法获取到
//	beanPtr：需要解析到的结构体指针，操作完后，声明的结构体会被赋值
//	文档：https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func DecryptOpenDataToStruct(encryptedData, iv, sessionKey string, beanPtr interface{}) (err error) {
	if encryptedData == util.NULL || iv == util.NULL || sessionKey == util.NULL {
		return errors.New("input params can not null")
	}
	var (
		cipherText, aesKey, ivKey, plainText []byte
		block                                cipher.Block
		blockMode                            cipher.BlockMode
	)
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("传入beanPtr类型必须是以指针形式")
	}
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入interface{}必须是结构体")
	}
	cipherText, _ = base64.StdEncoding.DecodeString(encryptedData)
	aesKey, _ = base64.StdEncoding.DecodeString(sessionKey)
	ivKey, _ = base64.StdEncoding.DecodeString(iv)
	// fmt.Println("ivKey:", string(ivKey))
	if len(cipherText)%len(aesKey) != 0 {
		return errors.New("encryptedData is error")
	}
	if block, err = aes.NewCipher(aesKey); err != nil {
		return fmt.Errorf("aes.NewCipher：%w", err)
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	plainText = make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	if len(plainText) > 0 {
		plainText = xaes.PKCS7UnPadding(plainText)
	}
	if err = json.Unmarshal(plainText, beanPtr); err != nil {
		return fmt.Errorf("json.Marshal(%s)：%w", string(plainText), err)
	}

	return
}

// 解密开放数据
func Decrypt(session_key, iv, encrypted_data string) ([]byte, error) {
	if len := strings.Count(session_key, "") - 1; len != 24 {
		return nil, errors.New("invalid value session_key")
	}
	aesKey, err := base64.StdEncoding.DecodeString(session_key)
	if err != nil {
		return nil, err
	}

	if len := strings.Count(iv, "") - 1; len != 24 {
		return nil, errors.New("invalid value iv")
	}
	ivKey, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	decodeData, err := base64.StdEncoding.DecodeString(encrypted_data)
	if err != nil {
		return nil, err
	}

	dataBytes, err := AesDecrypt(decodeData, aesKey, ivKey)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}

func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)

	// 去除填充
	length := len(origData)
	unp := int(origData[length-1])
	return origData[:(length - unp)], nil
}
