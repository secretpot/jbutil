package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func Base64(plainText []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(plainText))
}
func Debase64(cipherText []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(cipherText))
}

func RSAEncrypt(plainText []byte, publicKeyContent []byte) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode(publicKeyContent)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}
	//返回密文
	return cipherText, nil
}
func RSADecrypt(cipherText []byte, privateKeyContent []byte) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode(privateKeyContent)
	//X509解码
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), cipherText)
	if err != nil {
		return nil, err
	}
	//返回明文
	return plainText, nil
}
