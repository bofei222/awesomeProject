package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
)

var prvKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCtSKTqLHEk2LvA2iey0jNqhvlIXA1eafpQDnqflsGdm41Dkp/i\nto/1d7nl7InHVJzd7M1ygsTXVsmcJK3GfMKhAiqbV0SXpHsHxGaMb/OchovGz9vo\nMfTeDEYLfQEdJCDDfpNw/OkJ8haouitA5y1r9Wz0v3ZJ/PTTezZ9NFnjlwIDAQAB\nAoGAWiDJnf8ljjuoVCNjuI+6LHMtn2Q5k9zdU9xkDTOVWst4SygtPvcjo1H1f9Bq\nzSGGQauUJDY9+Z7rV+p/9BgaBlbmiO2I6b/6hpiPjiN304obko14egtQt0M+b6/3\n6BDCR6TS/m9Q6OPu1MV2NNzJIHcT+APJAhSruTefcZyudvkCQQDDOj+bu1VTfYxf\nND9IEOaNp3xxvKjO19X1BkO+E94R5djJLG9T+5LbDPKSUWsYEYTQo41Pk03H9e/X\nHm696LS1AkEA4zmwbKw9SQzA4QQUEhVxL9XjdIRp37eDECOlyBt00v3g4m2/qlwF\nbpBXbWkW96/IxVvjwz38ukXmIbwMHljSmwJARXuMugOBidaMSDITN7X0KIRssRpB\nRmThHHTfVV5C0kHo1yi+crh9+wJvrw3VPNq3V35uQ90ceMeaVgjZxzN0ZQJADKs8\nwPJhgF8rqeAWVmPHqckdI3P2iziqIA48wgl13AW3sig4VYFH9EAr/7eqRikQ4qAa\n9NnR63jWOe9IJbECHQJAAcFguUJAHTxlcN7/9oz08gu2aVJjk0CG03yS07GvqRk+\nrT1kKr5Y04V9t9SmoRHTClCyaJH1XUR+JO2HF3v/Ww==\n-----END RSA PRIVATE KEY-----"
var pubKey = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCtSKTqLHEk2LvA2iey0jNqhvlI\nXA1eafpQDnqflsGdm41Dkp/ito/1d7nl7InHVJzd7M1ygsTXVsmcJK3GfMKhAiqb\nV0SXpHsHxGaMb/OchovGz9voMfTeDEYLfQEdJCDDfpNw/OkJ8haouitA5y1r9Wz0\nv3ZJ/PTTezZ9NFnjlwIDAQAB\n-----END PUBLIC KEY-----"

func main() {
	//rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	//prvKey, pubKey := GenRsaKey()
	fmt.Println(string(prvKey))
	fmt.Println(string(pubKey))

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	fmt.Println("对消息进行签名操作...")
	signData := RsaSignWithSha256([]byte(data), []byte(prvKey))
	fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("\n对签名信息进行验证...")
	if RsaVerySignWithSha256([]byte(data), signData, []byte(pubKey)) {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext := RsaEncrypt([]byte(data), []byte(pubKey))
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData := RsaDecrypt(ciphertext, []byte(prvKey))
	fmt.Println("私钥解密后的数据：", string(sourceData))
}

//RSA公钥私钥产生
func GenRsaKey() (prvkey, pubkey []byte) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey = pem.EncodeToMemory(block)
	return
}

//签名
func RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature
}

//验证
func RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

// 公钥加密
func RsaEncrypt(data, keyBytes []byte) []byte {
	//解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) []byte {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		panic(err)
	}
	return data
}
