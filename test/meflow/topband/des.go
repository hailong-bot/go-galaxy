package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

// Padding 填充数据，使其长度为块大小的整数倍（PKCS5Padding）。
func Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// Unpadding 去除填充的额外数据。
func Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// Encrypt DESC 加密 (CFB 模式)
func Encrypt(data, key []byte) (string, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	iv := key[:blockSize] // 使用 key 的前 blockSize 字节作为 IV (初始化向量)

	stream := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(data))
	stream.XORKeyStream(cipherText, data)

	// 加密后用 Base64 编码方便传输
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt DESC 解密 (CFB 模式)
func Decrypt(base64Data string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	iv := key[:blockSize] // 使用 key 的前 blockSize 字节作为 IV

	stream := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(data))
	stream.XORKeyStream(plainText, data)

	return string(plainText), nil
}
