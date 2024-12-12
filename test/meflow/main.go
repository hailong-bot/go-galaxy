package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/url"
	"time"
)

func main() {
	time.Now().Format("")
}

func main1() {
	var (
		StampDetailURL = "%s/#/approval-order?applyId=%s"
		host           = "http://121.40.97.11:8008"
		applyID        = "f5d74098-2665-4937-a3f3-6004392ab06e"
		publicKey      = "MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJik6X5GaNihtBfL1NIh95/KnA3bnWuNnmUB8WMd2Mtq9CwK1pZoFRXevqzZZvpp0fTPJ23rqPCZUXXFXjuVyzsCAwEAAQ=="
		redirctURL     = "%s/sso/login?username=%s&callbackUrl=%s"
	)
	stampUrl := fmt.Sprintf(StampDetailURL, host, applyID)
	escape := url.QueryEscape(stampUrl)
	fmt.Println(escape)
	encrypt, err := RsaEncrypt("1000003791", publicKey)
	if err != nil {
		log.Fatal(err)
	}
	escapeEncryptUserID := url.QueryEscape(encrypt)
	result := fmt.Sprintf(redirctURL, host, escapeEncryptUserID, escape)
	fmt.Println(result)
}

// RSA 加密
func RsaEncrypt(plainText, publicKey string) (string, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", errors.WithStack(err)
	}

	pubKey, err := x509.ParsePKIXPublicKey(decodedKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse public key: %v", err)
	}
	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("not a valid RSA publicKey")
	}
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPubKey, []byte(plainText))
	if err != nil {
		return "", fmt.Errorf("faied to encrypt data: %v", err)
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

//func RsaEncrypt(plainText, publicKey string) (string, error) {
//	decodedKey, err := base64.StdEncoding.DecodeString(publicKey)
//	if err != nil {
//		return "", errors.WithStack(err)
//	}
//	pubKey, err := x509.ParsePKIXPublicKey(decodedKey)
//	if err != nil {
//		return "", fmt.Errorf("failed to parse public key: %v", err)
//	}
//	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
//	if !ok {
//		return "", errors.New("not a valid RSA publicKey")
//	}
//
//	message, err := EncryptMessage(rsaPubKey, plainText)
//	//cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPubKey, []byte(plainText))
//	//if err != nil {
//	//	return "", fmt.Errorf("faied to encrypt data: %v", err)
//	//}
//	//return base64.StdEncoding.EncodeToString(cipherText), nil
//	if err != nil {
//		return "", errors.WithStack(err)
//	}
//	return message, nil
//}
//
//func EncryptMessage(publicKey *rsa.PublicKey, message string) (string, error) {
//	data := []byte(message)
//	keySize, srcSize := publicKey.Size(), len(data)
//
//	//单次加密的长度需要减掉padding的长度，PKCS1为11
//	offSet, once := 0, keySize-11
//	buffer := bytes.Buffer{}
//	for offSet < srcSize {
//		endIndex := offSet + once
//		if endIndex > srcSize {
//			endIndex = srcSize
//		}
//		// 加密一部分
//		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data[offSet:endIndex])
//		if err != nil {
//			return "", err
//		}
//		buffer.Write(bytesOnce)
//		offSet = endIndex
//	}
//	bytesEncrypt := base64.StdEncoding.EncodeToString(buffer.Bytes())
//	return bytesEncrypt, nil
//	//ciphertext, err := rsa.EncryptOAEP(
//	//	sha256.New(),
//	//	rand.Reader,
//	//	publicKey,
//	//	[]byte(message),
//	//	nil,
//	//)
//	//if err != nil {
//	//	return "", err
//	//}
//	//return base64.StdEncoding.EncodeToString(ciphertext), nil
//}
