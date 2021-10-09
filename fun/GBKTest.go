package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// UTF-8 转 GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// 编码转换测试
func main2() {

	s := "[debug]: 编码转换测试"

	gbk, err := Utf8ToGbk([]byte(s))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(gbk))
		fmt.Println(111)
	}

	utf8, err := GbkToUtf8(gbk)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(utf8))
	}
}

// 编码转换测试
func main() {

	src := "编码转换内容内容"
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
	//fmt.Println(data) //byte
	fmt.Println(string(data)) //打印为乱码
}
