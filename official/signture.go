package official

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
)

// 验证消息的确来自微信服务器()
// 微信加密签名，signature结合了开发者填写的 token 参数和请求中的 timestamp 参数、nonce参数。
func CheckSignature(token, signature, timestamp, nonce string) bool {
	params := []string{
		token,
		timestamp,
		nonce,
	}
	sort.Strings(params)

	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}

	tempStr := fmt.Sprintf("%x", h.Sum(nil))
	return tempStr == signature
}
