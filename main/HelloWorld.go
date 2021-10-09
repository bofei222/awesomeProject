package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
	//"github.com/farmerx/gorsa"
)

//go:embed message.txt
var message string

//var privateKey []byte

func main() {

	name := "172.172.R8AQd8ZICP1TvVGB5NRERBNAFDEXf1NLM+A1zIZC1PSgkCOJt5esDbncErQKk7w5wj6Do5uop2TtC6cJTlXZ/0Rqxr+cVXujFpw01RLeDQsCPtcn2BD7Fy69fSfYc8FGnLyABYtydrJyfRMpC4YoEX9rA9HunB/QcBxJxBd05ow=EiQxgNXw80brPFJ+sK8FoecsIzuWqqDg0iwetAOgMzYcUEnpMYR5R0TQeWpbeQl6m88KgQASQ/gaIWhPqhbDPykaHUGR+1WzWrJr1dag6sJWE3efHvbpX8FzVPHyYKPrPs6pPr3rCg9tThXaHd7AXlzFEfCizgDZxSFTiXSsoos=5ad102c61f9b64ec89c5342f0725c62fef1334551fb96072254dfe11d5e9cee0746fce53098cf65a2eeca993c7d29c000bfee0160781adb96fcf24bafc35b046"
	nameArr := strings.Split(name, ".")

	aesKeyLenth, err := strconv.Atoi(nameArr[0])
	if err != nil {
		return
	}

	shaLength, err := strconv.Atoi(nameArr[1])
	//shaLength : = nameArr[0]

	ms := nameArr[2]
	//
	aesKeyCipher, err := RsaDecrypt2(ms[0:aesKeyLenth])
	if nil != err {
		fmt.Println(err)
	}

	// 将rsa加密的 数据体摘要 解密
	jsonDataHash, err := RsaDecrypt2(ms[aesKeyLenth : aesKeyLenth+shaLength])
	if nil != err {
		fmt.Println(err)
	}
	//
	fmt.Println(jsonDataHash)

	dencryTxt := AesDecryptCBC(ms[aesKeyLenth+shaLength:], aesKeyCipher)

	sum := sha256.Sum256([]byte(dencryTxt))
	//fmt.Printf("%x", sum)
	toString := hex.EncodeToString(sum[:])
	fmt.Println(toString)
	if jsonDataHash != toString {
		fmt.Println("验签错误")
		os.Exit(-3)
	}

	fmt.Println(dencryTxt)

	println(name) // Tom
	timeUnix := time.Now().Unix()
	fmt.Println(timeUnix * 1000)

	resultMap := make(map[string]interface{})
	json.Unmarshal([]byte(dencryTxt), &resultMap)

	cpuid := resultMap["cpuid"]

	if getCpuId() != cpuid {
		fmt.Println("无效授权")
		os.Exit(-3)
	}

	parseInt := int64(resultMap["expirationTime"].(float64))
	//parseInt, err := strconv.ParseInt(a, 10, 64)
	if parseInt < timeUnix*100 {
		fmt.Println("授权过期， 请重新获取授权")
		os.Exit(-3)
	}
	fmt.Println("授权通过")
}

func getCpuId() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	//	fmt.Println(string(out))
	str := string(out)
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	return str[11:]
}

func AesDecryptCBC(encryptHexString string, aesKeyString string) (decryptedString string) {
	key, _ := hex.DecodeString(aesKeyString)
	block, _ := aes.NewCipher(key) // 分组秘钥
	//blockSize := block.BlockSize()
	iv := []byte("abcdefghijk1mnop")               // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, iv) // 加密模式
	encrypted, _ := hex.DecodeString(encryptHexString)
	decrypted := make([]byte, len(encrypted))   // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted) // 解密
	decrypted = pkcs5UnPadding(decrypted)       // 去除补全码
	return string(decrypted)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 解密
func RsaDecrypt2(encryptedString string) (string, error) {
	base64DecodeBytes, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}
	var prvKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCtSKTqLHEk2LvA2iey0jNqhvlIXA1eafpQDnqflsGdm41Dkp/i\nto/1d7nl7InHVJzd7M1ygsTXVsmcJK3GfMKhAiqbV0SXpHsHxGaMb/OchovGz9vo\nMfTeDEYLfQEdJCDDfpNw/OkJ8haouitA5y1r9Wz0v3ZJ/PTTezZ9NFnjlwIDAQAB\nAoGAWiDJnf8ljjuoVCNjuI+6LHMtn2Q5k9zdU9xkDTOVWst4SygtPvcjo1H1f9Bq\nzSGGQauUJDY9+Z7rV+p/9BgaBlbmiO2I6b/6hpiPjiN304obko14egtQt0M+b6/3\n6BDCR6TS/m9Q6OPu1MV2NNzJIHcT+APJAhSruTefcZyudvkCQQDDOj+bu1VTfYxf\nND9IEOaNp3xxvKjO19X1BkO+E94R5djJLG9T+5LbDPKSUWsYEYTQo41Pk03H9e/X\nHm696LS1AkEA4zmwbKw9SQzA4QQUEhVxL9XjdIRp37eDECOlyBt00v3g4m2/qlwF\nbpBXbWkW96/IxVvjwz38ukXmIbwMHljSmwJARXuMugOBidaMSDITN7X0KIRssRpB\nRmThHHTfVV5C0kHo1yi+crh9+wJvrw3VPNq3V35uQ90ceMeaVgjZxzN0ZQJADKs8\nwPJhgF8rqeAWVmPHqckdI3P2iziqIA48wgl13AW3sig4VYFH9EAr/7eqRikQ4qAa\n9NnR63jWOe9IJbECHQJAAcFguUJAHTxlcN7/9oz08gu2aVJjk0CG03yS07GvqRk+\nrT1kKr5Y04V9t9SmoRHTClCyaJH1XUR+JO2HF3v/Ww==\n-----END RSA PRIVATE KEY-----"
	privateKeyBlock, _ := pem.Decode([]byte(prvKey))
	var pri *rsa.PrivateKey
	pri, parseErr := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if parseErr != nil {
		return "", parseErr
	}
	decryptedData, decryptErr := rsa.DecryptOAEP(sha1.New(), rand.Reader, pri, base64DecodeBytes, nil)
	if decryptErr != nil {
		return "", decryptErr
	}

	return string(decryptedData), nil
}
