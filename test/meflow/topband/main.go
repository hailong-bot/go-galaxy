package main

import "fmt"

func main() {
	// 原始数据
	plainText := "Hello, Golang DESC CFB!"

	// 密钥：长度必须为 8 字节
	key := []byte("MEFLOWOA")

	// 加密
	encryptedText, err := Encrypt([]byte(plainText), key)
	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}
	fmt.Println("加密结果:", encryptedText)

	// 解密
	decryptedText, err := Decrypt(encryptedText, key)
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}
	fmt.Println("解密结果:", decryptedText)
}
