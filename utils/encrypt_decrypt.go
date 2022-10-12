package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//func Dncrypt(rawData, key, iv string) (string, error) {
//	data, err := base64.StdEncoding.DecodeString(rawData)
//	keyB, err1 := base64.StdEncoding.DecodeString(key)
//	ivB, _ := base64.StdEncoding.DecodeString(iv)
//	if err != nil {
//		return "", err
//	}
//	if err1 != nil {
//		return "", err1
//	}
//	dnData, err := AesDecrypt(data, keyB, ivB)
//	if err != nil {
//		return "", err
//	}
//	return string(dnData), nil
//}

func AesDecrypt(encryptedData string, sessionKey string, iv string) ([]byte, error) {
	//Base64解码
	keyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	cryptData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	origData := make([]byte, len(cryptData))
	//AES
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	//CBC
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	//解密
	mode.CryptBlocks(origData, cryptData)
	//去除填充位
	origData = PKCS7UnPadding(origData)
	return origData, nil

	//block, err := aes.NewCipher(sessionKey)
	//if err != nil {
	//	global.GvaLog.Error("NewCipher 失败", zap.Error(err))
	//}
	//blockSize := block.BlockSize()
	//if len(encryptedData) < blockSize {
	//	global.GvaLog.Info("ciphertext too short")
	//}
	//if len(encryptedData)%blockSize != 0 {
	//	global.GvaLog.Info("ciphertext is not a multiple of the block size")
	//}
	//mode := cipher.NewCBCDecrypter(block, iv)
	//mode.CryptBlocks(encryptedData, encryptedData)
	//// 解填充
	//encryptedData = PKCS7UnPadding(encryptedData)
	//return encryptedData, nil

}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	if length > 0 {
		unPadding := int(plantText[length-1])
		return plantText[:(length - unPadding)]
	}
	return plantText
}
