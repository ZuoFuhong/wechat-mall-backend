package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

type WxSensitiveData struct {
	AppId         string
	SessionKey    string
	Iv            string
	EncryptedData string
}

func (r *WxSensitiveData) Decrypt() (map[string]interface{}, error) {
	var userData = map[string]interface{}{}
	sessionKey, err1 := base64.StdEncoding.DecodeString(r.SessionKey)
	iv, err2 := base64.StdEncoding.DecodeString(r.Iv)
	encryptedData, err3 := base64.StdEncoding.DecodeString(r.EncryptedData)
	if err1 != nil || err2 != nil || err3 != nil {
		return nil, errors.New("base64解密出错")
	}
	cipherBlock, err := aes.NewCipher([]byte(sessionKey))
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(cipherBlock, iv)
	mode.CryptBlocks(encryptedData, encryptedData)
	decrypted := r.unPad(encryptedData)
	err = json.Unmarshal(decrypted, &userData)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (r *WxSensitiveData) unPad(s []byte) []byte {
	return s[:(len(s) - int(s[len(s)-1]))]
}
