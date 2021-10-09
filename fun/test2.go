package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func main() {
	str, err := RsaDecrypt22("OVg3ctVS2tOOx0uspnFXSwrRe2WCIe1AqHHWW4iKMSBhbPkbkS/bfS9+8DprrygtpihHG2wRP+xbcG6ajrM200L8A+d9yAU5W0q/jc6VHA3rM7FEmlYZjvpdm7QIQiNC1n/bML7hHMOJdPTtckbwPxGsxUXhYCuD5rrzlzL8MZY=")
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(str)

}
func RsaDecrypt22(encryptedString string) (string, error) {
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
