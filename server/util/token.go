/*
 * @Author: F1nley
 * @Date: 2021-10-03 10:40:15
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 11:24:34
 * @Description: token有关的工具类
 */
package util

import (
	"Jinghong-LoveWall/server/global"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strconv"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
var c cipher.Block
var cfb, cfbdec cipher.Stream
var isInited bool = false

func initToken() error {
	var e error = nil
	c, e = aes.NewCipher([]byte(global.TokenPrivateKey))
	if e != nil {
		return e
	}
	cfb = cipher.NewCFBEncrypter(c, commonIV)
	cfbdec = cipher.NewCFBDecrypter(c, commonIV)

	isInited = true
	return e
}

func GetToken(uid int64) (string, error) {
	var err error
	if !isInited {
		err = initToken()
		if err != nil {
			return "", err
		}
	}

	platinText := []byte(strconv.FormatInt(uid, 10))
	cipherText := make([]byte, len(platinText))
	cfb.XORKeyStream(cipherText, platinText)
	return base64.StdEncoding.EncodeToString(cipherText), err
}

func DecodeToken(data string) (int64, error) {
	if !isInited {
		err := initToken()
		if err != nil {
			return 0, err
		}
	}
	cipherText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return 0, err
	}
	platinText := make([]byte, len(cipherText))
	cfbdec.XORKeyStream(platinText, cipherText)
	res, err := strconv.ParseInt(string(platinText), 10, 64)
	return res, err
}
