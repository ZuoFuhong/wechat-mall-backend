package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"time"
)

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
}

type CallbackParam struct {
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}

type OSSPolicyToken struct {
	AccessKeyId     string
	AccessKeySecret string
	Host            string
	CallbackUrl     string
	UploadDir       string
	ExpireTime      int64
}

func (oss OSSPolicyToken) GetPolicyToken() PolicyToken {
	expireEnd := time.Now().Unix() + oss.ExpireTime

	//create post policy json
	var config ConfigStruct
	config.Expiration = time.Unix(expireEnd, 0).Format("2006-01-02T15:04:05Z")
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, oss.UploadDir)
	config.Conditions = append(config.Conditions, condition)

	// calucate signature
	result, err := json.Marshal(config)
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(oss.AccessKeySecret))
	_, _ = io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var callbackParam CallbackParam
	callbackParam.CallbackUrl = oss.CallbackUrl
	callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	callbackStr, err := json.Marshal(callbackParam)
	if err != nil {
		fmt.Println("callback json err:", err)
	}
	callbackBase64 := base64.StdEncoding.EncodeToString(callbackStr)

	var policyToken PolicyToken
	policyToken.AccessKeyId = oss.AccessKeyId
	policyToken.Host = oss.Host
	policyToken.Expire = expireEnd
	policyToken.Signature = string(signedStr)
	policyToken.Directory = oss.UploadDir
	policyToken.Policy = string(debyte)
	policyToken.Callback = string(callbackBase64)
	return policyToken
}
