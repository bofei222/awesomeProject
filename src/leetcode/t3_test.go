package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestT3(t *testing.T) {
	//len := lengthOfLongestSubstring("abcabcdbb") //
	len := lengthOfLongestSubstring("ohomm") //
	//len := lengthOfLongestSubstring("dvdf") //
	//len := lengthOfLongestSubstring("pwwkew") //
	fmt.Println(len)
}

func lengthOfLongestSubstring(s string) int {
	return 1
}

//输入: s = "abcabcbb"
//输出: 3
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度
func lengthOfLongestSubstring0(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var longStr string
	var preLongString string
	j := 0
	for i := 0; i < len(s); i++ {

		sprintf := fmt.Sprintf("%c", s[i])
		if i == 0 {
			longStr = sprintf
		} else {
			if !strings.Contains(longStr, sprintf) {
				longStr = longStr + sprintf
				continue
			}
			if j == 0 {
				preLongString = longStr
				j++
			}
			index := strings.Index(longStr, sprintf)
			i = i - len(longStr) + index

			if len(longStr) > len(preLongString) {
				preLongString = longStr
			}
			longStr = ""
			continue
		}
	}
	if len(preLongString) > len(longStr) {
		fmt.Println(preLongString)
		return len(preLongString)
	} else {
		fmt.Println(longStr)
		return len(longStr)
	}
}

func lengthOfLongestSubstring1(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var longStr string
	var preLongString string
	fmt.Println(preLongString)
	j := 0
	for i := 0; i < len(s); i++ {

		sprintf := fmt.Sprintf("%c", s[i])
		if i == 0 {
			longStr = sprintf
		} else {
			if !strings.Contains(longStr, sprintf) {
				longStr = longStr + sprintf
				continue
			}
			if j == 0 {
				preLongString = longStr
				j++
			}
			if len(longStr) > len(preLongString) {
				preLongString = longStr
			}
			longStr = sprintf
		}
	}
	if len(preLongString) > len(longStr) {
		return len(preLongString)
	} else {
		return len(longStr)
	}
}
func TestChinaStringForrU(t *testing.T) {
	s := "我是中国人"
	for index, runeValue := range s {
		fmt.Printf("%#U 起始于字位置%d\n", runeValue, index)
	}
}

func TestChinaStringForrC(t *testing.T) {
	s := "我是中国人"
	for index, runeValue := range s {
		fmt.Printf("%c 起始于字位置%d\n", runeValue, index)
	}
}

func TestChinaStringFor(t *testing.T) {
	s := "我是中国人"
	fmt.Printf("% x\n", s)
	for i := 0; i < len(s); i++ {
		//fmt.Println(s[i])
		fmt.Printf("%c 起始于字位置%d\n", s[i], i)
	}
}

func TestEnglishStringFor(t *testing.T) {
	s := "abcabcbb"
	fmt.Printf("% x\n", s)
	for i := 0; i < len(s); i++ {
		//fmt.Println(s[i])
		fmt.Printf("%c 起始于字位置%d\n", s[i], i)
	}
}
